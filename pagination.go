package propraven

import (
	"context"
	"encoding/json"
	"fmt"
)

// IterOptions controls an auto-paginating iterator (the XxxIter methods).
type IterOptions struct {
	// PageSize is the limit sent per page. Zero uses the params' Limit, or
	// 100 when that is unset too. The server may clamp it; the iterator
	// follows the limit the server echoes back.
	PageSize int
	// MaxItems stops the iteration after this many items. Zero means no cap.
	MaxItems int
}

// Iter walks every item across pages, fetching lazily:
//
//	it := client.Deals.AbsenteeIter(ctx, &propraven.DealsAbsenteeParams{CountyFIPS: propraven.String("37119")},
//		propraven.IterOptions{PageSize: 100, MaxItems: 1000})
//	for it.Next() {
//		row := it.Current()
//		_ = row
//	}
//	if err := it.Err(); err != nil { ... }
//
// The context passed to the XxxIter method governs every page request.
type Iter[T any] struct {
	ctx     context.Context
	fetch   func(ctx context.Context) ([]T, bool, error) // items, more, err
	buf     []T
	idx     int
	cur     T
	err     error
	more    bool
	max     int
	yielded int
	pages   int
}

func newIter[T any](ctx context.Context, max int, fetch func(ctx context.Context) ([]T, bool, error)) *Iter[T] {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Iter[T]{ctx: ctx, fetch: fetch, more: true, max: max}
}

// Next advances to the next item, fetching the next page when needed. It
// returns false at the end or on error; check Err afterwards.
func (it *Iter[T]) Next() bool {
	if it.err != nil {
		return false
	}
	if it.max > 0 && it.yielded >= it.max {
		return false
	}
	for it.idx >= len(it.buf) {
		if !it.more {
			return false
		}
		items, more, err := it.fetch(it.ctx)
		if err != nil {
			it.err = err
			return false
		}
		it.pages++
		it.buf, it.idx, it.more = items, 0, more
		if len(items) == 0 {
			it.more = false
			return false
		}
	}
	it.cur = it.buf[it.idx]
	it.idx++
	it.yielded++
	return true
}

// Current returns the item at the iterator's position.
func (it *Iter[T]) Current() T { return it.cur }

// Err returns the first error met while paginating.
func (it *Iter[T]) Err() error { return it.err }

// Pages returns how many pages have been fetched so far.
func (it *Iter[T]) Pages() int { return it.pages }

// page is the generic view of one page body.
type page struct {
	raw map[string]json.RawMessage
}

func parsePage(body json.RawMessage) (page, error) {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(body, &m); err != nil {
		return page{}, fmt.Errorf("propraven: decoding page: %w", err)
	}
	return page{raw: m}, nil
}

func pageItems[T any](p page, key string) ([]T, error) {
	v, ok := p.raw[key]
	if !ok || string(v) == "null" {
		return nil, nil
	}
	var items []T
	if err := json.Unmarshal(v, &items); err != nil {
		return nil, fmt.Errorf("propraven: decoding page items %q: %w", key, err)
	}
	return items, nil
}

func (p page) int(key string) (int64, bool) {
	v, ok := p.raw[key]
	if !ok {
		return 0, false
	}
	var f *float64
	if json.Unmarshal(v, &f) != nil || f == nil {
		return 0, false
	}
	return int64(*f), true
}

func (p page) boolean(keys ...string) (bool, bool) {
	for _, k := range keys {
		if v, ok := p.raw[k]; ok {
			var b *bool
			if json.Unmarshal(v, &b) == nil && b != nil {
				return *b, true
			}
		}
	}
	return false, false
}

func (p page) str(key string) string {
	var s *string
	if v, ok := p.raw[key]; ok && json.Unmarshal(v, &s) == nil && s != nil {
		return *s
	}
	return ""
}

func pageSize(o IterOptions, paramLimit *int64) int64 {
	if o.PageSize > 0 {
		return int64(o.PageSize)
	}
	if paramLimit != nil && *paramLimit > 0 {
		return *paramLimit
	}
	return 100
}

// newOffsetIter pages with limit/offset: offset advances by the effective
// limit; it stops on a short page, when offset >= total, or has_more=false.
func newOffsetIter[T any](ctx context.Context, c *Client, o IterOptions, itemsKey string, paramLimit, paramOffset *int64,
	build func(limit, offset int64) *apiRequest, opts []RequestOption) *Iter[T] {
	limit := pageSize(o, paramLimit)
	var offset int64
	if paramOffset != nil {
		offset = *paramOffset
	}
	return newIter[T](ctx, o.MaxItems, func(ctx context.Context) ([]T, bool, error) {
		var body json.RawMessage
		if err := c.do(ctx, build(limit, offset), opts, decodeJSON(&body)); err != nil {
			return nil, false, err
		}
		p, err := parsePage(body)
		if err != nil {
			return nil, false, err
		}
		items, err := pageItems[T](p, itemsKey)
		if err != nil {
			return nil, false, err
		}
		eff := limit
		if echoed, ok := p.int("limit"); ok && echoed > 0 {
			eff = echoed
		}
		n := int64(len(items))
		more := n > 0 && n >= eff
		offset += eff
		if total, ok := p.int("total"); ok && offset >= total {
			more = false
		}
		if hm, ok := p.boolean("has_more", "hasMore"); ok && !hm {
			more = false
		}
		limit = eff
		return items, more, nil
	})
}

// newCursorIter pages with an opaque cursor: it passes the previous page's
// next cursor until it is null/absent or hasMore is false.
func newCursorIter[T any](ctx context.Context, c *Client, o IterOptions, itemsKey, nextKey string, paramLimit *int64,
	build func(limit *int64, cursor *string) *apiRequest, opts []RequestOption) *Iter[T] {
	var limit *int64
	if o.PageSize > 0 {
		l := int64(o.PageSize)
		limit = &l
	} else if paramLimit != nil {
		l := *paramLimit
		limit = &l
	}
	var cursor *string
	return newIter[T](ctx, o.MaxItems, func(ctx context.Context) ([]T, bool, error) {
		var body json.RawMessage
		if err := c.do(ctx, build(limit, cursor), opts, decodeJSON(&body)); err != nil {
			return nil, false, err
		}
		p, err := parsePage(body)
		if err != nil {
			return nil, false, err
		}
		items, err := pageItems[T](p, itemsKey)
		if err != nil {
			return nil, false, err
		}
		next := p.str(nextKey)
		more := next != "" && len(items) > 0
		if hm, ok := p.boolean("hasMore", "has_more"); ok && !hm {
			more = false
		}
		if more {
			cursor = &next
		}
		return items, more, nil
	})
}
