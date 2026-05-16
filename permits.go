package propraven

import (
	"context"
	"net/url"
)

// PermitsService binds to /v1/permits.
type PermitsService struct {
	client *Client
}

// PermitSearchOptions filter [PermitsService.Search].
type PermitSearchOptions struct {
	ParcelID   string
	StateFIPS  string
	CountyFIPS string
	PermitType string
	SinceDate  string // YYYY-MM-DD; only permits filed on or after this date
	Limit      int
	Cursor     string
}

func (o PermitSearchOptions) query() url.Values {
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
	if o.PermitType != "" {
		v.Set("permit_type", o.PermitType)
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

// PermitPage is a single page of [Permit].
type PermitPage struct {
	Data       []Permit `json:"data"`
	NextCursor string   `json:"next_cursor"`
}

// Search returns permits matching opts.
func (s *PermitsService) Search(ctx context.Context, opts PermitSearchOptions) (*PermitPage, error) {
	var page PermitPage
	if err := s.client.do(ctx, "GET", "/v1/permits", opts.query(), nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}
