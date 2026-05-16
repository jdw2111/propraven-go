package propraven

import (
	"context"
	"net/url"
	"strconv"
)

// ParcelsService binds to /v1/parcels.
type ParcelsService struct {
	client *Client
}

// Get returns a single parcel by its composite parcel_id (county_fips:apn).
// Returns a 404-wrapped [*Error] when the parcel id is not found.
func (s *ParcelsService) Get(ctx context.Context, parcelID string) (*Parcel, error) {
	var p Parcel
	err := s.client.do(ctx, "GET", "/v1/parcels/"+url.PathEscape(parcelID), nil, nil, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// ListOptions are the query options accepted by [ParcelsService.List].
// All fields are optional; zero values are skipped.
type ListOptions struct {
	StateFIPS    string
	CountyFIPS   string
	OwnerName    string
	OwnerEntity  string
	PropertyType string
	Limit        int
	Cursor       string
}

func (o ListOptions) query() url.Values {
	v := url.Values{}
	if o.StateFIPS != "" {
		v.Set("state_fips", o.StateFIPS)
	}
	if o.CountyFIPS != "" {
		v.Set("county_fips", o.CountyFIPS)
	}
	if o.OwnerName != "" {
		v.Set("owner_name", o.OwnerName)
	}
	if o.OwnerEntity != "" {
		v.Set("owner_entity_id", o.OwnerEntity)
	}
	if o.PropertyType != "" {
		v.Set("property_type", o.PropertyType)
	}
	if o.Limit > 0 {
		v.Set("limit", strconv.Itoa(o.Limit))
	}
	if o.Cursor != "" {
		v.Set("cursor", o.Cursor)
	}
	return v
}

// ParcelPage is one page of [Parcel] results plus a cursor for the next page.
type ParcelPage struct {
	Data       []Parcel `json:"data"`
	NextCursor string   `json:"next_cursor"`
}

// List fetches a single page of parcels matching opts. Use [ParcelsService.Iter]
// to range over all pages.
func (s *ParcelsService) List(ctx context.Context, opts ListOptions) (*ParcelPage, error) {
	var page ParcelPage
	if err := s.client.do(ctx, "GET", "/v1/parcels", opts.query(), nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

// Iter returns an iterator over all pages of parcels matching opts. The first
// error short-circuits iteration; subsequent calls to the yielded function
// return zero value.
//
//	for parcel, err := range client.Parcels.Iter(ctx, opts) {
//	    if err != nil { return err }
//	    // ...
//	}
func (s *ParcelsService) Iter(ctx context.Context, opts ListOptions) func(yield func(Parcel, error) bool) {
	return func(yield func(Parcel, error) bool) {
		for {
			page, err := s.List(ctx, opts)
			if err != nil {
				yield(Parcel{}, err)
				return
			}
			for _, p := range page.Data {
				if !yield(p, nil) {
					return
				}
			}
			if page.NextCursor == "" {
				return
			}
			opts.Cursor = page.NextCursor
		}
	}
}
