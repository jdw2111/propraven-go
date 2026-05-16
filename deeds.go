package propraven

import (
	"context"
	"net/url"
)

// DeedsService binds to /v1/deeds.
type DeedsService struct {
	client *Client
}

// DeedSearchOptions filter [DeedsService.Search]. ParcelID OR (StateFIPS+CountyFIPS)
// is required by the API.
type DeedSearchOptions struct {
	ParcelID   string
	StateFIPS  string
	CountyFIPS string
	SinceDate  string // YYYY-MM-DD; only sales recorded on or after this date
	Limit      int
	Cursor     string
}

func (o DeedSearchOptions) query() url.Values {
	v := url.Values{}
	if o.ParcelID != "" {
		v.Set("parcel_id", o.ParcelID)
	}
	if o.StateFIPS != "" {
		v.Set("state_fips", o.StateFIPS)
	}
	if o.CountyFIPS != "" {
		v.Set("county_fips", o.CountyFIPS)
	}
	if o.SinceDate != "" {
		v.Set("since_date", o.SinceDate)
	}
	if o.Limit > 0 {
		v.Set("limit", itoa(o.Limit))
	}
	if o.Cursor != "" {
		v.Set("cursor", o.Cursor)
	}
	return v
}

// DeedPage is a single page of [Deed].
type DeedPage struct {
	Data       []Deed `json:"data"`
	NextCursor string `json:"next_cursor"`
}

// Search returns deeds matching opts.
func (s *DeedsService) Search(ctx context.Context, opts DeedSearchOptions) (*DeedPage, error) {
	var page DeedPage
	if err := s.client.do(ctx, "GET", "/v1/deeds", opts.query(), nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}
