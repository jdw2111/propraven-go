package propraven

import (
	"context"
	"net/url"
)

// OwnersService binds to /v1/owners.
type OwnersService struct {
	client *Client
}

// Get returns a single owner entity by its consolidated owner_entity_id.
func (s *OwnersService) Get(ctx context.Context, ownerEntityID string) (*Owner, error) {
	var o Owner
	err := s.client.do(ctx, "GET", "/v1/owners/"+url.PathEscape(ownerEntityID), nil, nil, &o)
	if err != nil {
		return nil, err
	}
	return &o, nil
}

// OwnerSearchOptions filter [OwnersService.Search].
type OwnerSearchOptions struct {
	Query     string // free-text owner-name match
	Ticker    string // exact match against the public-company ticker registry
	StateFIPS string
	Limit     int
	Cursor    string
}

func (o OwnerSearchOptions) query() url.Values {
	v := url.Values{}
	if o.Query != "" {
		v.Set("q", o.Query)
	}
	if o.Ticker != "" {
		v.Set("ticker", o.Ticker)
	}
	if o.StateFIPS != "" {
		v.Set("state_fips", o.StateFIPS)
	}
	if o.Limit > 0 {
		v.Set("limit", itoa(o.Limit))
	}
	if o.Cursor != "" {
		v.Set("cursor", o.Cursor)
	}
	return v
}

// OwnerPage is a single page of owner results.
type OwnerPage struct {
	Data       []Owner `json:"data"`
	NextCursor string  `json:"next_cursor"`
}

// Search returns one page of owners matching opts.
func (s *OwnersService) Search(ctx context.Context, opts OwnerSearchOptions) (*OwnerPage, error) {
	var page OwnerPage
	if err := s.client.do(ctx, "GET", "/v1/owners", opts.query(), nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	const max = 12
	var buf [max]byte
	neg := n < 0
	if neg {
		n = -n
	}
	i := max
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
