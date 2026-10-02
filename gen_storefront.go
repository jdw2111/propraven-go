// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"context"
	"encoding/json"
)

// StorefrontService groups the storefront operations. Use it as client.Storefront.
type StorefrontService struct {
	client *Client
}

// Catalog: Machine Storefront — sealed field catalog
//
// The data catalog IS the storefront: every column of the serving parcel relation with its
// measured national (and optional per-state) coverage, grain, tier, honesty flags, and pipeline
// freshness warts, plus the dossier pricing model and base quote (the per-parcel value-tiered
// price comes from /storefront/availability?parcel_id=). FREE -- authenticated callers are
// unmetered and uncapped (rate-limited at the scale window); anonymous callers are served and
// IP-throttled at the free tier. Answered from a committed, sealed artifact -- no database access.
//
// HTTP: GET /api/v1/storefront/catalog
func (s *StorefrontService) Catalog(ctx context.Context, params *StorefrontCatalogParams, opts ...RequestOption) (*StorefrontCatalogResponse, error) {
	var out StorefrontCatalogResponse
	if err := s.client.do(ctx, buildStorefrontCatalogRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildStorefrontCatalogRequest(params *StorefrontCatalogParams) *apiRequest {
	req := newRequest("GET", "/api/v1/storefront/catalog")
	if params != nil {
		addQuery(req.query, "state", params.State)
		addQuery(req.query, "tier", params.Tier)
		addQuery(req.query, "section", params.Section)
		addQuery(req.query, "grain", params.Grain)
		addQuery(req.query, "min_coverage", params.MinCoverage)
		addQuery(req.query, "red_cells", params.RedCells)
		addQuery(req.query, "q", params.Q)
		addQuery(req.query, "include", params.Include)
		addQuery(req.query, "fields", params.Fields)
	}
	return req
}

// StorefrontCatalogParams holds the query, header and JSON-body parameters of
// [StorefrontService.Catalog]. Pass nil when you need none.
type StorefrontCatalogParams struct {
	// USPS code ("NC") or 2-digit FIPS ("37"). Adds per-state coverage to every field and a state gaps
	// block.
	State *string `query:"state" json:"-"`

	// prime|strong|good|partial|sparse|trace — filter by coverage tier.
	Tier *string `query:"tier" json:"-"`

	// Filter to one catalog section (identity, valuation, hazard, …).
	Section *string `query:"section" json:"-"`

	// parcel|county|tract|block_group|zip|unknown.
	Grain *string `query:"grain" json:"-"`

	// 0..1 — only fields at/above this national coverage.
	MinCoverage *string `query:"min_coverage" json:"-"`

	// 1 → only fields carrying a red-cell honesty flag.
	RedCells *string `query:"red_cells" json:"-"`

	// Substring on name, label or description.
	Q *string `query:"q" json:"-"`

	// plumbing → include the internal provenance columns.
	Include *string `query:"include" json:"-"`

	// none → metadata header only, without the ~615 field entries.
	Fields *string `query:"fields" json:"-"`
}

// StorefrontCatalogResponse: The sealed field catalog. Additional top-level blocks (headline,
// sections, tiers, freshness, advisories, state_index, and an optional per-state `state` block)
// are present; the load-bearing ones are documented here.
type StorefrontCatalogResponse struct {
	ContractVersion float64 `json:"contract_version"`

	// The serving epoch this catalog was built for.
	CatalogVersion string `json:"catalog_version"`
	GeneratedAt    string `json:"generated_at"`

	// Content hash: two callers under the same seal get byte-identical numbers.
	Seal        StorefrontCatalogResponseSeal        `json:"seal"`
	Disclosures StorefrontCatalogResponseDisclosures `json:"disclosures"`

	// National counts (distinct parcels, rows, mapped locations, geocoded parcels) each with the
	// definition it was measured under.
	MeasuredFrom StorefrontCatalogResponseMeasuredFrom `json:"measured_from"`
	Counts       StorefrontCatalogResponseCounts       `json:"counts"`
	Tiers        StorefrontCatalogResponseTiers        `json:"tiers"`
	Headline     StorefrontCatalogResponseHeadline     `json:"headline"`
	Sections     StorefrontCatalogResponseSections     `json:"sections"`

	// The dossier pricing model (base price, floor/cap, add-ons).
	Pricing StorefrontCatalogResponsePricing `json:"pricing"`

	// The base dossier quote derived from the pricing model. The per-parcel, value-tiered price is
	// quoted by /storefront/availability?parcel_id=.
	Quote StorefrontCatalogResponseQuote `json:"quote"`

	// Every non-ok freshness probe, shipped rather than hidden (the honesty layer).
	Warts      []StorefrontCatalogResponseWarts      `json:"warts"`
	Freshness  []StorefrontCatalogResponseFreshness  `json:"freshness"`
	Advisories []StorefrontCatalogResponseAdvisories `json:"advisories"`
	StateIndex []json.RawMessage                     `json:"state_index"`
	State      StorefrontCatalogResponseState        `json:"state"`

	// Number of field entries returned after filters.
	FieldCount int64 `json:"field_count"`

	// The field dictionary: one entry per serving column with name, label, section, grain, tier,
	// national (and optional per-state) coverage, and honesty flags.
	Fields []StorefrontCatalogResponseFields `json:"fields"`
}

// UnmarshalJSON decodes StorefrontCatalogResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponse) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponse
	aux := struct {
		*plain
		ContractVersion lenientNumber[float64] `json:"contract_version"`
		FieldCount      lenientNumber[int64]   `json:"field_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ContractVersion.assign(&r.ContractVersion)
	aux.FieldCount.assign(&r.FieldCount)
	return softTypeError(err)
}

// StorefrontCatalogResponseSeal: Content hash: two callers under the same seal get byte-identical
// numbers.
type StorefrontCatalogResponseSeal struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
	Covers    string `json:"covers"`
}

// StorefrontCatalogResponseDisclosures is generated from the OpenAPI spec.
type StorefrontCatalogResponseDisclosures struct {
	AnnotationVersion       string                                                `json:"annotation_version"`
	SealScope               string                                                `json:"seal_scope"`
	WithheldFields          []StorefrontCatalogResponseDisclosuresWithheldFields  `json:"withheld_fields"`
	AnnotationBasis         string                                                `json:"annotation_basis"`
	ReferenceCatalog        StorefrontCatalogResponseDisclosuresReferenceCatalog  `json:"reference_catalog"`
	ReferenceCatalogMatches bool                                                  `json:"reference_catalog_matches"`
	CatalogPublishedAt      string                                                `json:"catalog_published_at"`
	CatalogGeneratedAt      string                                                `json:"catalog_generated_at"`
	ServingEpochMatch       string                                                `json:"serving_epoch_match"`
	CurrentSourceVintage    string                                                `json:"current_source_vintage"`
	FreshnessScope          string                                                `json:"freshness_scope"`
	CoverageScope           string                                                `json:"coverage_scope"`
	PricingScope            string                                                `json:"pricing_scope"`
	Applicability           string                                                `json:"applicability"`
	FieldAdvisories         []StorefrontCatalogResponseDisclosuresFieldAdvisories `json:"field_advisories"`
}

// StorefrontCatalogResponseDisclosuresWithheldFields is generated from the OpenAPI spec.
type StorefrontCatalogResponseDisclosuresWithheldFields struct {
	Field      *string `json:"field,omitempty"`
	Reason     *string `json:"reason,omitempty"`
	WithheldOn *string `json:"withheld_on,omitempty"`
	Note       *string `json:"note,omitempty"`
}

// StorefrontCatalogResponseDisclosuresReferenceCatalog is generated from the OpenAPI spec.
type StorefrontCatalogResponseDisclosuresReferenceCatalog struct {
	Version string `json:"version"`
	Seal    string `json:"seal"`
}

// StorefrontCatalogResponseDisclosuresFieldAdvisories is generated from the OpenAPI spec.
type StorefrontCatalogResponseDisclosuresFieldAdvisories struct {
	Field                   *string `json:"field,omitempty"`
	DocumentedAt            *string `json:"documented_at,omitempty"`
	EvidenceSource          *string `json:"evidence_source,omitempty"`
	EvidenceScope           *string `json:"evidence_scope,omitempty"`
	ReferenceCatalogVersion *string `json:"reference_catalog_version,omitempty"`
	ReferenceCatalogSeal    *string `json:"reference_catalog_seal,omitempty"`
	ReferenceCatalogMatches *bool   `json:"reference_catalog_matches,omitempty"`
	CurrentValueVerified    *bool   `json:"current_value_verified,omitempty"`
	Code                    *string `json:"code,omitempty"`
	Note                    *string `json:"note,omitempty"`
}

// StorefrontCatalogResponseMeasuredFrom: National counts (distinct parcels, rows, mapped
// locations, geocoded parcels) each with the definition it was measured under.
type StorefrontCatalogResponseMeasuredFrom struct {
	Relation                   string                                                      `json:"relation"`
	Schema                     string                                                      `json:"schema"`
	Parent                     string                                                      `json:"parent"`
	Parcels                    float64                                                     `json:"parcels"`
	States                     float64                                                     `json:"states"`
	Method                     string                                                      `json:"method"`
	WeightSource               string                                                      `json:"weight_source"`
	WeightsReconcileToSnapshot bool                                                        `json:"weights_reconcile_to_snapshot"`
	Manifest                   string                                                      `json:"manifest"`
	SnapshotWatermark          string                                                      `json:"snapshot_watermark"`
	SnapshotSwappedAt          string                                                      `json:"snapshot_swapped_at"`
	ParityStatus               string                                                      `json:"parity_status"`
	PartitionAnalyzeMin        string                                                      `json:"partition_analyze_min"`
	PartitionAnalyzeMax        string                                                      `json:"partition_analyze_max"`
	GeocodeCoverage            float64                                                     `json:"geocode_coverage"`
	Sampling                   string                                                      `json:"sampling"`
	Baseline                   StorefrontCatalogResponseMeasuredFromBaseline               `json:"baseline"`
	ParcelCountsEpoch          string                                                      `json:"parcel_counts_epoch"`
	MappedLocations            float64                                                     `json:"mapped_locations"`
	GeocodedParcels            float64                                                     `json:"geocoded_parcels"`
	Rows                       float64                                                     `json:"rows"`
	ParcelCountDefinitions     StorefrontCatalogResponseMeasuredFromParcelCountDefinitions `json:"parcel_count_definitions"`
	GeocodeCoverageBasis       string                                                      `json:"geocode_coverage_basis"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseMeasuredFrom, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseMeasuredFrom) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseMeasuredFrom
	aux := struct {
		*plain
		Parcels         lenientNumber[float64] `json:"parcels"`
		States          lenientNumber[float64] `json:"states"`
		GeocodeCoverage lenientNumber[float64] `json:"geocode_coverage"`
		MappedLocations lenientNumber[float64] `json:"mapped_locations"`
		GeocodedParcels lenientNumber[float64] `json:"geocoded_parcels"`
		Rows            lenientNumber[float64] `json:"rows"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Parcels.assign(&r.Parcels)
	aux.States.assign(&r.States)
	aux.GeocodeCoverage.assign(&r.GeocodeCoverage)
	aux.MappedLocations.assign(&r.MappedLocations)
	aux.GeocodedParcels.assign(&r.GeocodedParcels)
	aux.Rows.assign(&r.Rows)
	return softTypeError(err)
}

// StorefrontCatalogResponseMeasuredFromBaseline is generated from the OpenAPI spec.
type StorefrontCatalogResponseMeasuredFromBaseline struct {
	Source      string `json:"source"`
	LastAnalyze string `json:"last_analyze"`
	Note        string `json:"note"`
}

// StorefrontCatalogResponseMeasuredFromParcelCountDefinitions is generated from the OpenAPI spec.
type StorefrontCatalogResponseMeasuredFromParcelCountDefinitions struct {
	Parcels         string `json:"parcels"`
	MappedLocations string `json:"mapped_locations"`
	GeocodedParcels string `json:"geocoded_parcels"`
	Rows            string `json:"rows"`
}

// StorefrontCatalogResponseCounts is generated from the OpenAPI spec.
type StorefrontCatalogResponseCounts struct {
	Columns                 float64                                     `json:"columns"`
	Sellable                float64                                     `json:"sellable"`
	Plumbing                float64                                     `json:"plumbing"`
	HeadlineParcelGrainAt80 float64                                     `json:"headline_parcel_grain_at_80"`
	RedCells                float64                                     `json:"red_cells"`
	EpochRegressions        float64                                     `json:"epoch_regressions"`
	CurationReviewRequired  float64                                     `json:"curation_review_required"`
	ByGrain                 StorefrontCatalogResponseCountsByGrain      `json:"by_grain"`
	CumulativeAt            StorefrontCatalogResponseCountsCumulativeAt `json:"cumulative_at"`
	Withheld                float64                                     `json:"withheld"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseCounts, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseCounts) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseCounts
	aux := struct {
		*plain
		Columns                 lenientNumber[float64] `json:"columns"`
		Sellable                lenientNumber[float64] `json:"sellable"`
		Plumbing                lenientNumber[float64] `json:"plumbing"`
		HeadlineParcelGrainAt80 lenientNumber[float64] `json:"headline_parcel_grain_at_80"`
		RedCells                lenientNumber[float64] `json:"red_cells"`
		EpochRegressions        lenientNumber[float64] `json:"epoch_regressions"`
		CurationReviewRequired  lenientNumber[float64] `json:"curation_review_required"`
		Withheld                lenientNumber[float64] `json:"withheld"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Columns.assign(&r.Columns)
	aux.Sellable.assign(&r.Sellable)
	aux.Plumbing.assign(&r.Plumbing)
	aux.HeadlineParcelGrainAt80.assign(&r.HeadlineParcelGrainAt80)
	aux.RedCells.assign(&r.RedCells)
	aux.EpochRegressions.assign(&r.EpochRegressions)
	aux.CurationReviewRequired.assign(&r.CurationReviewRequired)
	aux.Withheld.assign(&r.Withheld)
	return softTypeError(err)
}

// StorefrontCatalogResponseCountsByGrain is generated from the OpenAPI spec.
type StorefrontCatalogResponseCountsByGrain struct {
	County  int64   `json:"county"`
	Parcel  float64 `json:"parcel"`
	Tract   float64 `json:"tract"`
	Unknown float64 `json:"unknown"`
	Zip     float64 `json:"zip"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseCountsByGrain, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseCountsByGrain) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseCountsByGrain
	aux := struct {
		*plain
		County  lenientNumber[int64]   `json:"county"`
		Parcel  lenientNumber[float64] `json:"parcel"`
		Tract   lenientNumber[float64] `json:"tract"`
		Unknown lenientNumber[float64] `json:"unknown"`
		Zip     lenientNumber[float64] `json:"zip"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.County.assign(&r.County)
	aux.Parcel.assign(&r.Parcel)
	aux.Tract.assign(&r.Tract)
	aux.Unknown.assign(&r.Unknown)
	aux.Zip.assign(&r.Zip)
	return softTypeError(err)
}

// StorefrontCatalogResponseCountsCumulativeAt is generated from the OpenAPI spec.
type StorefrontCatalogResponseCountsCumulativeAt struct {
	V095 float64 `json:"0.95"`
	V080 float64 `json:"0.80"`
	V060 float64 `json:"0.60"`
	V030 float64 `json:"0.30"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseCountsCumulativeAt, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseCountsCumulativeAt) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseCountsCumulativeAt
	aux := struct {
		*plain
		V095 lenientNumber[float64] `json:"0.95"`
		V080 lenientNumber[float64] `json:"0.80"`
		V060 lenientNumber[float64] `json:"0.60"`
		V030 lenientNumber[float64] `json:"0.30"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.V095.assign(&r.V095)
	aux.V080.assign(&r.V080)
	aux.V060.assign(&r.V060)
	aux.V030.assign(&r.V030)
	return softTypeError(err)
}

// StorefrontCatalogResponseTiers is generated from the OpenAPI spec.
type StorefrontCatalogResponseTiers struct {
	Prime   float64 `json:"prime"`
	Strong  float64 `json:"strong"`
	Good    float64 `json:"good"`
	Partial float64 `json:"partial"`
	Sparse  float64 `json:"sparse"`
	Trace   float64 `json:"trace"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseTiers, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseTiers) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseTiers
	aux := struct {
		*plain
		Prime   lenientNumber[float64] `json:"prime"`
		Strong  lenientNumber[float64] `json:"strong"`
		Good    lenientNumber[float64] `json:"good"`
		Partial lenientNumber[float64] `json:"partial"`
		Sparse  lenientNumber[float64] `json:"sparse"`
		Trace   lenientNumber[float64] `json:"trace"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Prime.assign(&r.Prime)
	aux.Strong.assign(&r.Strong)
	aux.Good.assign(&r.Good)
	aux.Partial.assign(&r.Partial)
	aux.Sparse.assign(&r.Sparse)
	aux.Trace.assign(&r.Trace)
	return softTypeError(err)
}

// StorefrontCatalogResponseHeadline is generated from the OpenAPI spec.
type StorefrontCatalogResponseHeadline struct {
	Floor             float64 `json:"floor"`
	ParcelGrainFields float64 `json:"parcel_grain_fields"`
	AllSellableFields float64 `json:"all_sellable_fields"`
	Claim             string  `json:"claim"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseHeadline, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseHeadline) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseHeadline
	aux := struct {
		*plain
		Floor             lenientNumber[float64] `json:"floor"`
		ParcelGrainFields lenientNumber[float64] `json:"parcel_grain_fields"`
		AllSellableFields lenientNumber[float64] `json:"all_sellable_fields"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Floor.assign(&r.Floor)
	aux.ParcelGrainFields.assign(&r.ParcelGrainFields)
	aux.AllSellableFields.assign(&r.AllSellableFields)
	return softTypeError(err)
}

// StorefrontCatalogResponseSections is generated from the OpenAPI spec.
type StorefrontCatalogResponseSections struct {
	Parcel   StorefrontCatalogResponseSectionsParcel   `json:"parcel"`
	Permits  StorefrontCatalogResponseSectionsPermits  `json:"permits"`
	Deeds    StorefrontCatalogResponseSectionsDeeds    `json:"deeds"`
	Comps    StorefrontCatalogResponseSectionsComps    `json:"comps"`
	Geometry StorefrontCatalogResponseSectionsGeometry `json:"geometry"`
}

// StorefrontCatalogResponseSectionsParcel is generated from the OpenAPI spec.
type StorefrontCatalogResponseSectionsParcel struct {
	Relation string  `json:"relation"`
	Rows     float64 `json:"rows"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseSectionsParcel, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseSectionsParcel) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseSectionsParcel
	aux := struct {
		*plain
		Rows lenientNumber[float64] `json:"rows"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Rows.assign(&r.Rows)
	return softTypeError(err)
}

// StorefrontCatalogResponseSectionsPermits is generated from the OpenAPI spec.
type StorefrontCatalogResponseSectionsPermits struct {
	Relation      string  `json:"relation"`
	Rows          float64 `json:"rows"`
	NationalShare float64 `json:"national_share"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseSectionsPermits, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseSectionsPermits) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseSectionsPermits
	aux := struct {
		*plain
		Rows          lenientNumber[float64] `json:"rows"`
		NationalShare lenientNumber[float64] `json:"national_share"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Rows.assign(&r.Rows)
	aux.NationalShare.assign(&r.NationalShare)
	return softTypeError(err)
}

// StorefrontCatalogResponseSectionsDeeds is generated from the OpenAPI spec.
type StorefrontCatalogResponseSectionsDeeds struct {
	Relation      string  `json:"relation"`
	Rows          float64 `json:"rows"`
	NationalShare float64 `json:"national_share"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseSectionsDeeds, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseSectionsDeeds) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseSectionsDeeds
	aux := struct {
		*plain
		Rows          lenientNumber[float64] `json:"rows"`
		NationalShare lenientNumber[float64] `json:"national_share"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Rows.assign(&r.Rows)
	aux.NationalShare.assign(&r.NationalShare)
	return softTypeError(err)
}

// StorefrontCatalogResponseSectionsComps is generated from the OpenAPI spec.
type StorefrontCatalogResponseSectionsComps struct {
	Relation      string  `json:"relation"`
	Rows          float64 `json:"rows"`
	NationalShare float64 `json:"national_share"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseSectionsComps, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseSectionsComps) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseSectionsComps
	aux := struct {
		*plain
		Rows          lenientNumber[float64] `json:"rows"`
		NationalShare lenientNumber[float64] `json:"national_share"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Rows.assign(&r.Rows)
	aux.NationalShare.assign(&r.NationalShare)
	return softTypeError(err)
}

// StorefrontCatalogResponseSectionsGeometry is generated from the OpenAPI spec.
type StorefrontCatalogResponseSectionsGeometry struct {
	Relation      string  `json:"relation"`
	Rows          float64 `json:"rows"`
	NationalShare float64 `json:"national_share"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseSectionsGeometry, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseSectionsGeometry) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseSectionsGeometry
	aux := struct {
		*plain
		Rows          lenientNumber[float64] `json:"rows"`
		NationalShare lenientNumber[float64] `json:"national_share"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Rows.assign(&r.Rows)
	aux.NationalShare.assign(&r.NationalShare)
	return softTypeError(err)
}

// StorefrontCatalogResponsePricing: The dossier pricing model (base price, floor/cap, add-ons).
type StorefrontCatalogResponsePricing struct {
	Dossier      StorefrontCatalogResponsePricingDossier  `json:"dossier"`
	AddOns       []StorefrontCatalogResponsePricingAddOns `json:"add_ons"`
	Catalog      string                                   `json:"catalog"`
	Availability string                                   `json:"availability"`
	PurchaseFlow string                                   `json:"purchase_flow"`
}

// StorefrontCatalogResponsePricingDossier is generated from the OpenAPI spec.
type StorefrontCatalogResponsePricingDossier struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// StorefrontCatalogResponsePricingAddOns is generated from the OpenAPI spec.
type StorefrontCatalogResponsePricingAddOns struct {
	Code        *string                                     `json:"code,omitempty"`
	Label       *string                                     `json:"label,omitempty"`
	Price       StorefrontCatalogResponsePricingAddOnsPrice `json:"price"`
	Purchasable *bool                                       `json:"purchasable,omitempty"`
	Note        *string                                     `json:"note,omitempty"`
}

// StorefrontCatalogResponsePricingAddOnsPrice is generated from the OpenAPI spec.
type StorefrontCatalogResponsePricingAddOnsPrice struct {
	Amount   *string `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// StorefrontCatalogResponseQuote: The base dossier quote derived from the pricing model. The
// per-parcel, value-tiered price is quoted by /storefront/availability?parcel_id=.
type StorefrontCatalogResponseQuote struct {
	Price        StorefrontCatalogResponseQuotePrice    `json:"price"`
	Currency     string                                 `json:"currency"`
	AddOns       []StorefrontCatalogResponseQuoteAddOns `json:"add_ons"`
	Pay          []json.RawMessage                      `json:"pay"`
	PurchaseFlow string                                 `json:"purchase_flow"`
}

// StorefrontCatalogResponseQuotePrice is generated from the OpenAPI spec.
type StorefrontCatalogResponseQuotePrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// StorefrontCatalogResponseQuoteAddOns is generated from the OpenAPI spec.
type StorefrontCatalogResponseQuoteAddOns struct {
	Code        *string                                   `json:"code,omitempty"`
	Label       *string                                   `json:"label,omitempty"`
	Price       StorefrontCatalogResponseQuoteAddOnsPrice `json:"price"`
	Available   *bool                                     `json:"available,omitempty"`
	Purchasable *bool                                     `json:"purchasable,omitempty"`
	Note        *string                                   `json:"note,omitempty"`
}

// StorefrontCatalogResponseQuoteAddOnsPrice is generated from the OpenAPI spec.
type StorefrontCatalogResponseQuoteAddOnsPrice struct {
	Amount   *string `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// StorefrontCatalogResponseWarts is generated from the OpenAPI spec.
type StorefrontCatalogResponseWarts struct {
	Dataset         *string  `json:"dataset,omitempty"`
	Surface         *string  `json:"surface,omitempty"`
	Status          *string  `json:"status,omitempty"`
	ObservedMaxDate *string  `json:"observed_max_date,omitempty"`
	AgeDays         *float64 `json:"age_days,omitempty"`
	MaxAgeDays      *int64   `json:"max_age_days,omitempty"`
	ProbedAt        *string  `json:"probed_at,omitempty"`
	Reason          *string  `json:"reason,omitempty"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseWarts, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseWarts) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseWarts
	aux := struct {
		*plain
		AgeDays    lenientNumber[float64] `json:"age_days"`
		MaxAgeDays lenientNumber[int64]   `json:"max_age_days"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.AgeDays.assignPtr(&r.AgeDays)
	aux.MaxAgeDays.assignPtr(&r.MaxAgeDays)
	return softTypeError(err)
}

// StorefrontCatalogResponseFreshness is generated from the OpenAPI spec.
type StorefrontCatalogResponseFreshness struct {
	Dataset         *string  `json:"dataset,omitempty"`
	Surface         *string  `json:"surface,omitempty"`
	Status          *string  `json:"status,omitempty"`
	ObservedMaxDate *string  `json:"observed_max_date,omitempty"`
	AgeDays         *float64 `json:"age_days,omitempty"`
	MaxAgeDays      *int64   `json:"max_age_days,omitempty"`
	ProbedAt        *string  `json:"probed_at,omitempty"`
	Reason          *string  `json:"reason,omitempty"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseFreshness, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseFreshness) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseFreshness
	aux := struct {
		*plain
		AgeDays    lenientNumber[float64] `json:"age_days"`
		MaxAgeDays lenientNumber[int64]   `json:"max_age_days"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.AgeDays.assignPtr(&r.AgeDays)
	aux.MaxAgeDays.assignPtr(&r.MaxAgeDays)
	return softTypeError(err)
}

// StorefrontCatalogResponseAdvisories is generated from the OpenAPI spec.
type StorefrontCatalogResponseAdvisories struct {
	Code          *string                                            `json:"code,omitempty"`
	Severity      *string                                            `json:"severity,omitempty"`
	Headline      *string                                            `json:"headline,omitempty"`
	Cause         *string                                            `json:"cause,omitempty"`
	Effect        *string                                            `json:"effect,omitempty"`
	Remedy        *string                                            `json:"remedy,omitempty"`
	WorstAffected []StorefrontCatalogResponseAdvisoriesWorstAffected `json:"worst_affected,omitempty"`
}

// StorefrontCatalogResponseAdvisoriesWorstAffected is generated from the OpenAPI spec.
type StorefrontCatalogResponseAdvisoriesWorstAffected struct {
	Name             *string  `json:"name,omitempty"`
	Coverage         *float64 `json:"coverage,omitempty"`
	BaselineCoverage *float64 `json:"baseline_coverage,omitempty"`
	Delta            *float64 `json:"delta,omitempty"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseAdvisoriesWorstAffected, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseAdvisoriesWorstAffected) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseAdvisoriesWorstAffected
	aux := struct {
		*plain
		Coverage         lenientNumber[float64] `json:"coverage"`
		BaselineCoverage lenientNumber[float64] `json:"baseline_coverage"`
		Delta            lenientNumber[float64] `json:"delta"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Coverage.assignPtr(&r.Coverage)
	aux.BaselineCoverage.assignPtr(&r.BaselineCoverage)
	aux.Delta.assignPtr(&r.Delta)
	return softTypeError(err)
}

// StorefrontCatalogResponseState is generated from the OpenAPI spec.
type StorefrontCatalogResponseState struct {
	StateFIPS       string                                    `json:"state_fips"`
	State           string                                    `json:"state"`
	Counties        int64                                     `json:"counties"`
	Parcels         float64                                   `json:"parcels"`
	RollupParcels   float64                                   `json:"rollup_parcels"`
	ShareOfNational float64                                   `json:"share_of_national"`
	Coverage        StorefrontCatalogResponseStateCoverage    `json:"coverage"`
	LastAnalyze     string                                    `json:"last_analyze"`
	WorstGaps       []StorefrontCatalogResponseStateWorstGaps `json:"worst_gaps"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseState, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseState) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseState
	aux := struct {
		*plain
		Counties        lenientNumber[int64]   `json:"counties"`
		Parcels         lenientNumber[float64] `json:"parcels"`
		RollupParcels   lenientNumber[float64] `json:"rollup_parcels"`
		ShareOfNational lenientNumber[float64] `json:"share_of_national"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Counties.assign(&r.Counties)
	aux.Parcels.assign(&r.Parcels)
	aux.RollupParcels.assign(&r.RollupParcels)
	aux.ShareOfNational.assign(&r.ShareOfNational)
	return softTypeError(err)
}

// StorefrontCatalogResponseStateCoverage is generated from the OpenAPI spec.
type StorefrontCatalogResponseStateCoverage struct {
	Address  float64 `json:"address"`
	Geocode  float64 `json:"geocode"`
	Owner    float64 `json:"owner"`
	Value    float64 `json:"value"`
	Geometry float64 `json:"geometry"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseStateCoverage, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseStateCoverage) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseStateCoverage
	aux := struct {
		*plain
		Address  lenientNumber[float64] `json:"address"`
		Geocode  lenientNumber[float64] `json:"geocode"`
		Owner    lenientNumber[float64] `json:"owner"`
		Value    lenientNumber[float64] `json:"value"`
		Geometry lenientNumber[float64] `json:"geometry"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Address.assign(&r.Address)
	aux.Geocode.assign(&r.Geocode)
	aux.Owner.assign(&r.Owner)
	aux.Value.assign(&r.Value)
	aux.Geometry.assign(&r.Geometry)
	return softTypeError(err)
}

// StorefrontCatalogResponseStateWorstGaps is generated from the OpenAPI spec.
type StorefrontCatalogResponseStateWorstGaps struct {
	Name     *string  `json:"name,omitempty"`
	Label    *string  `json:"label,omitempty"`
	National *float64 `json:"national,omitempty"`
	State    *float64 `json:"state,omitempty"`
	Section  *string  `json:"section,omitempty"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseStateWorstGaps, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseStateWorstGaps) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseStateWorstGaps
	aux := struct {
		*plain
		National lenientNumber[float64] `json:"national"`
		State    lenientNumber[float64] `json:"state"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.National.assignPtr(&r.National)
	aux.State.assignPtr(&r.State)
	return softTypeError(err)
}

// StorefrontCatalogResponseFields is generated from the OpenAPI spec.
type StorefrontCatalogResponseFields struct {
	Name                   *string           `json:"name,omitempty"`
	Type                   *string           `json:"type,omitempty"`
	Section                *string           `json:"section,omitempty"`
	Grain                  *string           `json:"grain,omitempty"`
	Category               *string           `json:"category,omitempty"`
	Sellable               *bool             `json:"sellable,omitempty"`
	Label                  *string           `json:"label,omitempty"`
	Description            *string           `json:"description,omitempty"`
	Source                 *string           `json:"source,omitempty"`
	Curated                *bool             `json:"curated,omitempty"`
	Coverage               *float64          `json:"coverage,omitempty"`
	Tier                   *string           `json:"tier,omitempty"`
	BaselineCoverage       *float64          `json:"baseline_coverage,omitempty"`
	CoverageDelta          *float64          `json:"coverage_delta,omitempty"`
	StatesMeasured         *float64          `json:"states_measured,omitempty"`
	StatesWithAnyCoverage  *float64          `json:"states_with_any_coverage,omitempty"`
	NDistinct              *float64          `json:"n_distinct,omitempty"`
	Flags                  []*string         `json:"flags"`
	RedCell                *bool             `json:"red_cell,omitempty"`
	HeadlineEligible       *bool             `json:"headline_eligible,omitempty"`
	CurationReviewRequired *bool             `json:"curation_review_required,omitempty"`
	SourceAsOf             *string           `json:"source_as_of,omitempty"`
	SourceAsOfBasis        string            `json:"source_as_of_basis"`
	CatalogPublishedAt     *string           `json:"catalog_published_at,omitempty"`
	CatalogGeneratedAt     *string           `json:"catalog_generated_at,omitempty"`
	SourceLabelBasis       string            `json:"source_label_basis"`
	GrainBasis             string            `json:"grain_basis"`
	CatalogDescription     *string           `json:"catalog_description,omitempty"`
	HistoricalAdvisories   []json.RawMessage `json:"historical_advisories"`
	StateCoverage          *float64          `json:"state_coverage,omitempty"`
}

// UnmarshalJSON decodes StorefrontCatalogResponseFields, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontCatalogResponseFields) UnmarshalJSON(data []byte) error {
	type plain StorefrontCatalogResponseFields
	aux := struct {
		*plain
		Coverage              lenientNumber[float64] `json:"coverage"`
		BaselineCoverage      lenientNumber[float64] `json:"baseline_coverage"`
		CoverageDelta         lenientNumber[float64] `json:"coverage_delta"`
		StatesMeasured        lenientNumber[float64] `json:"states_measured"`
		StatesWithAnyCoverage lenientNumber[float64] `json:"states_with_any_coverage"`
		NDistinct             lenientNumber[float64] `json:"n_distinct"`
		StateCoverage         lenientNumber[float64] `json:"state_coverage"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Coverage.assignPtr(&r.Coverage)
	aux.BaselineCoverage.assignPtr(&r.BaselineCoverage)
	aux.CoverageDelta.assignPtr(&r.CoverageDelta)
	aux.StatesMeasured.assignPtr(&r.StatesMeasured)
	aux.StatesWithAnyCoverage.assignPtr(&r.StatesWithAnyCoverage)
	aux.NDistinct.assignPtr(&r.NDistinct)
	aux.StateCoverage.assignPtr(&r.StateCoverage)
	return softTypeError(err)
}

// Availability: Machine Storefront -- try-before-buy (jurisdiction coverage or per-parcel quote)
//
// Try-before-buy at two scopes, both FREE and answered from the sealed catalog:
//
// JURISDICTION mode (pass ?state=, optional &county=): how complete each field is in that
// state/county side-by-side with the national average, plus the worst local gaps and the base
// dossier quote. No parcel scan.
//
// PARCEL mode (pass ?parcel_id=): the exact value-tiered dossier price for ONE parcel BEFORE
// paying, with the V/R/F multiplier breakdown, band, and the cheap signals it was derived from.
// This uses the SAME quote math as the x402 402 on /parcels/{id}/report, so the previewed price
// equals the amount the payer is charged. Does one light, indexed single-parcel read.
//
// Provide EITHER parcel_id OR state. Same access model as the catalog: authenticated callers
// unmetered/uncapped, anonymous callers served and IP-throttled.
//
// HTTP: GET /api/v1/storefront/availability
func (s *StorefrontService) Availability(ctx context.Context, params *StorefrontAvailabilityParams, opts ...RequestOption) (*StorefrontAvailabilityResponse, error) {
	var out StorefrontAvailabilityResponse
	if err := s.client.do(ctx, buildStorefrontAvailabilityRequest(params), opts, decodeJSON(&out)); err != nil {
		return nil, err
	}
	return &out, nil
}

func buildStorefrontAvailabilityRequest(params *StorefrontAvailabilityParams) *apiRequest {
	req := newRequest("GET", "/api/v1/storefront/availability")
	if params != nil {
		addQuery(req.query, "parcel_id", params.ParcelID)
		addQuery(req.query, "state", params.State)
		addQuery(req.query, "county", params.County)
		addQuery(req.query, "limit", params.Limit)
	}
	return req
}

// StorefrontAvailabilityParams holds the query, header and JSON-body parameters of
// [StorefrontService.Availability]. Pass nil when you need none.
type StorefrontAvailabilityParams struct {
	// PARCEL mode. A canonical state_fips:county_fips:parcel_id, a legacy 5-digit-county:parcel_id, or
	// a parcel UUID. When present, returns the value-tiered dossier quote for this parcel
	// (state/county are ignored).
	ParcelID *string `query:"parcel_id" json:"-"`

	// JURISDICTION mode. USPS code ("NC") or 2-digit FIPS ("37"). Required unless parcel_id is given.
	State *string `query:"state" json:"-"`

	// 3-digit within-state code ("183") or 5-digit state+county ("37183"). Optional; narrows
	// JURISDICTION mode to one county.
	County *string `query:"county" json:"-"`

	// JURISDICTION mode only: cap on the returned county list (default 25, max 400).
	Limit *int64 `query:"limit" json:"-"`
}

// StorefrontAvailabilityResponse: Shape depends on mode. Both carry contract_version,
// catalog_version, generated_at and seal. PARCEL mode adds mode='parcel', a `parcel` identity
// block, and `dossier_quote`; JURISDICTION mode adds `jurisdiction`, `coverage`, `fields`,
// `worst_gaps`, `quote`, `warts` and `counties`.
type StorefrontAvailabilityResponse struct {
	ContractVersion float64                                   `json:"contract_version"`
	CatalogVersion  string                                    `json:"catalog_version"`
	GeneratedAt     string                                    `json:"generated_at"`
	Seal            StorefrontAvailabilityResponseSeal        `json:"seal"`
	Disclosures     StorefrontAvailabilityResponseDisclosures `json:"disclosures"`

	// JURISDICTION mode: the state/county being described.
	Jurisdiction *StorefrontAvailabilityResponseJurisdiction `json:"jurisdiction,omitempty"`

	// JURISDICTION mode: headline coverage facts and parcel counts.
	Coverage *StorefrontAvailabilityResponseCoverage `json:"coverage,omitempty"`

	// JURISDICTION mode: national vs in-jurisdiction field-coverage summaries.
	Fields    *StorefrontAvailabilityResponseFields     `json:"fields,omitempty"`
	WorstGaps []StorefrontAvailabilityResponseWorstGaps `json:"worst_gaps,omitempty"`

	// JURISDICTION mode: the base dossier quote.
	Quote *StorefrontAvailabilityResponseQuote  `json:"quote,omitempty"`
	Warts []StorefrontAvailabilityResponseWarts `json:"warts,omitempty"`

	// JURISDICTION mode: the state's counties by parcel count.
	Counties *StorefrontAvailabilityResponseCounties `json:"counties,omitempty"`

	// 'parcel' in PARCEL mode; absent in JURISDICTION mode.
	Mode *StorefrontAvailabilityResponseMode `json:"mode,omitempty"`

	// PARCEL mode: the resolved parcel identity.
	Parcel *StorefrontAvailabilityResponseParcel `json:"parcel,omitempty"`

	// PARCEL mode: the value-tiered quote for this parcel.
	DossierQuote *StorefrontAvailabilityResponseDossierQuote `json:"dossier_quote,omitempty"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponse, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponse) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponse
	aux := struct {
		*plain
		ContractVersion lenientNumber[float64] `json:"contract_version"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ContractVersion.assign(&r.ContractVersion)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseSeal is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseSeal struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
	Covers    string `json:"covers"`
}

// StorefrontAvailabilityResponseDisclosures is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseDisclosures struct {
	AnnotationVersion       string                                                     `json:"annotation_version"`
	SealScope               string                                                     `json:"seal_scope"`
	WithheldFields          []StorefrontAvailabilityResponseDisclosuresWithheldFields  `json:"withheld_fields"`
	AnnotationBasis         string                                                     `json:"annotation_basis"`
	ReferenceCatalog        StorefrontAvailabilityResponseDisclosuresReferenceCatalog  `json:"reference_catalog"`
	ReferenceCatalogMatches bool                                                       `json:"reference_catalog_matches"`
	CatalogPublishedAt      string                                                     `json:"catalog_published_at"`
	CatalogGeneratedAt      string                                                     `json:"catalog_generated_at"`
	ServingEpochMatch       string                                                     `json:"serving_epoch_match"`
	CurrentSourceVintage    string                                                     `json:"current_source_vintage"`
	FreshnessScope          string                                                     `json:"freshness_scope"`
	CoverageScope           string                                                     `json:"coverage_scope"`
	PricingScope            string                                                     `json:"pricing_scope"`
	Applicability           string                                                     `json:"applicability"`
	FieldAdvisories         []StorefrontAvailabilityResponseDisclosuresFieldAdvisories `json:"field_advisories"`
}

// StorefrontAvailabilityResponseDisclosuresWithheldFields is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseDisclosuresWithheldFields struct {
	Field      *string `json:"field,omitempty"`
	Reason     *string `json:"reason,omitempty"`
	WithheldOn *string `json:"withheld_on,omitempty"`
	Note       *string `json:"note,omitempty"`
}

// StorefrontAvailabilityResponseDisclosuresReferenceCatalog is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseDisclosuresReferenceCatalog struct {
	Version string `json:"version"`
	Seal    string `json:"seal"`
}

// StorefrontAvailabilityResponseDisclosuresFieldAdvisories is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseDisclosuresFieldAdvisories struct {
	Field                   *string `json:"field,omitempty"`
	DocumentedAt            *string `json:"documented_at,omitempty"`
	EvidenceSource          *string `json:"evidence_source,omitempty"`
	EvidenceScope           *string `json:"evidence_scope,omitempty"`
	ReferenceCatalogVersion *string `json:"reference_catalog_version,omitempty"`
	ReferenceCatalogSeal    *string `json:"reference_catalog_seal,omitempty"`
	ReferenceCatalogMatches *bool   `json:"reference_catalog_matches,omitempty"`
	CurrentValueVerified    *bool   `json:"current_value_verified,omitempty"`
	Code                    *string `json:"code,omitempty"`
	Note                    *string `json:"note,omitempty"`
}

// StorefrontAvailabilityResponseJurisdiction: JURISDICTION mode: the state/county being described.
type StorefrontAvailabilityResponseJurisdiction struct {
	Level           string  `json:"level"`
	State           string  `json:"state"`
	StateFIPS       string  `json:"state_fips"`
	CountyFIPS      string  `json:"county_fips"`
	CountyName      string  `json:"county_name"`
	Parcels         float64 `json:"parcels"`
	ShareOfNational float64 `json:"share_of_national"`
	Counties        int64   `json:"counties"`
	LastAnalyze     string  `json:"last_analyze"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseJurisdiction, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseJurisdiction) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseJurisdiction
	aux := struct {
		*plain
		Parcels         lenientNumber[float64] `json:"parcels"`
		ShareOfNational lenientNumber[float64] `json:"share_of_national"`
		Counties        lenientNumber[int64]   `json:"counties"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Parcels.assign(&r.Parcels)
	aux.ShareOfNational.assign(&r.ShareOfNational)
	aux.Counties.assign(&r.Counties)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseCoverage: JURISDICTION mode: headline coverage facts and parcel
// counts.
type StorefrontAvailabilityResponseCoverage struct {
	HeadlineFacts         StorefrontAvailabilityResponseCoverageHeadlineFacts `json:"headline_facts"`
	HeadlineFactsScope    string                                              `json:"headline_facts_scope"`
	StateRollupParcels    float64                                             `json:"state_rollup_parcels"`
	StateRatifiedParcels  float64                                             `json:"state_ratified_parcels"`
	StateParcelCountsUnit string                                              `json:"state_parcel_counts_unit"`
	National              StorefrontAvailabilityResponseCoverageNational      `json:"national"`
	NationalParcels       float64                                             `json:"national_parcels"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseCoverage, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseCoverage) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseCoverage
	aux := struct {
		*plain
		StateRollupParcels   lenientNumber[float64] `json:"state_rollup_parcels"`
		StateRatifiedParcels lenientNumber[float64] `json:"state_ratified_parcels"`
		NationalParcels      lenientNumber[float64] `json:"national_parcels"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.StateRollupParcels.assign(&r.StateRollupParcels)
	aux.StateRatifiedParcels.assign(&r.StateRatifiedParcels)
	aux.NationalParcels.assign(&r.NationalParcels)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseCoverageHeadlineFacts is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseCoverageHeadlineFacts struct {
	Address  float64 `json:"address"`
	Geocode  float64 `json:"geocode"`
	Owner    float64 `json:"owner"`
	Value    float64 `json:"value"`
	Geometry float64 `json:"geometry"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseCoverageHeadlineFacts, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseCoverageHeadlineFacts) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseCoverageHeadlineFacts
	aux := struct {
		*plain
		Address  lenientNumber[float64] `json:"address"`
		Geocode  lenientNumber[float64] `json:"geocode"`
		Owner    lenientNumber[float64] `json:"owner"`
		Value    lenientNumber[float64] `json:"value"`
		Geometry lenientNumber[float64] `json:"geometry"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Address.assign(&r.Address)
	aux.Geocode.assign(&r.Geocode)
	aux.Owner.assign(&r.Owner)
	aux.Value.assign(&r.Value)
	aux.Geometry.assign(&r.Geometry)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseCoverageNational is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseCoverageNational struct {
	ParcelCountsEpoch      string                                                               `json:"parcel_counts_epoch"`
	Parcels                float64                                                              `json:"parcels"`
	MappedLocations        float64                                                              `json:"mapped_locations"`
	GeocodedParcels        float64                                                              `json:"geocoded_parcels"`
	Rows                   float64                                                              `json:"rows"`
	ParcelCountDefinitions StorefrontAvailabilityResponseCoverageNationalParcelCountDefinitions `json:"parcel_count_definitions"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseCoverageNational, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseCoverageNational) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseCoverageNational
	aux := struct {
		*plain
		Parcels         lenientNumber[float64] `json:"parcels"`
		MappedLocations lenientNumber[float64] `json:"mapped_locations"`
		GeocodedParcels lenientNumber[float64] `json:"geocoded_parcels"`
		Rows            lenientNumber[float64] `json:"rows"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Parcels.assign(&r.Parcels)
	aux.MappedLocations.assign(&r.MappedLocations)
	aux.GeocodedParcels.assign(&r.GeocodedParcels)
	aux.Rows.assign(&r.Rows)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseCoverageNationalParcelCountDefinitions is generated from the
// OpenAPI spec.
type StorefrontAvailabilityResponseCoverageNationalParcelCountDefinitions struct {
	Parcels         string `json:"parcels"`
	MappedLocations string `json:"mapped_locations"`
	GeocodedParcels string `json:"geocoded_parcels"`
	Rows            string `json:"rows"`
}

// StorefrontAvailabilityResponseFields: JURISDICTION mode: national vs in-jurisdiction
// field-coverage summaries.
type StorefrontAvailabilityResponseFields struct {
	Sellable float64                                      `json:"sellable"`
	National StorefrontAvailabilityResponseFieldsNational `json:"national"`

	// (Only null in observed responses.)
	InJurisdiction        json.RawMessage `json:"in_jurisdiction"`
	CountyFieldResolution string          `json:"county_field_resolution"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseFields, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseFields) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseFields
	aux := struct {
		*plain
		Sellable lenientNumber[float64] `json:"sellable"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Sellable.assign(&r.Sellable)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseFieldsNational is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseFieldsNational struct {
	MeasuredFields   float64                                               `json:"measured_fields"`
	UnmeasuredFields float64                                               `json:"unmeasured_fields"`
	AtOrAbove        StorefrontAvailabilityResponseFieldsNationalAtOrAbove `json:"at_or_above"`
	Tiers            StorefrontAvailabilityResponseFieldsNationalTiers     `json:"tiers"`
	RedCells         float64                                               `json:"red_cells"`
	RedCellRule      string                                                `json:"red_cell_rule"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseFieldsNational, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseFieldsNational) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseFieldsNational
	aux := struct {
		*plain
		MeasuredFields   lenientNumber[float64] `json:"measured_fields"`
		UnmeasuredFields lenientNumber[float64] `json:"unmeasured_fields"`
		RedCells         lenientNumber[float64] `json:"red_cells"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.MeasuredFields.assign(&r.MeasuredFields)
	aux.UnmeasuredFields.assign(&r.UnmeasuredFields)
	aux.RedCells.assign(&r.RedCells)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseFieldsNationalAtOrAbove is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseFieldsNationalAtOrAbove struct {
	V095 float64 `json:"0.95"`
	V080 float64 `json:"0.80"`
	V060 float64 `json:"0.60"`
	V030 float64 `json:"0.30"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseFieldsNationalAtOrAbove, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseFieldsNationalAtOrAbove) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseFieldsNationalAtOrAbove
	aux := struct {
		*plain
		V095 lenientNumber[float64] `json:"0.95"`
		V080 lenientNumber[float64] `json:"0.80"`
		V060 lenientNumber[float64] `json:"0.60"`
		V030 lenientNumber[float64] `json:"0.30"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.V095.assign(&r.V095)
	aux.V080.assign(&r.V080)
	aux.V060.assign(&r.V060)
	aux.V030.assign(&r.V030)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseFieldsNationalTiers is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseFieldsNationalTiers struct {
	Prime   float64 `json:"prime"`
	Strong  float64 `json:"strong"`
	Good    float64 `json:"good"`
	Partial float64 `json:"partial"`
	Sparse  float64 `json:"sparse"`
	Trace   float64 `json:"trace"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseFieldsNationalTiers, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseFieldsNationalTiers) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseFieldsNationalTiers
	aux := struct {
		*plain
		Prime   lenientNumber[float64] `json:"prime"`
		Strong  lenientNumber[float64] `json:"strong"`
		Good    lenientNumber[float64] `json:"good"`
		Partial lenientNumber[float64] `json:"partial"`
		Sparse  lenientNumber[float64] `json:"sparse"`
		Trace   lenientNumber[float64] `json:"trace"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Prime.assign(&r.Prime)
	aux.Strong.assign(&r.Strong)
	aux.Good.assign(&r.Good)
	aux.Partial.assign(&r.Partial)
	aux.Sparse.assign(&r.Sparse)
	aux.Trace.assign(&r.Trace)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseWorstGaps is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseWorstGaps struct {
	Name     *string  `json:"name,omitempty"`
	Label    *string  `json:"label,omitempty"`
	National *float64 `json:"national,omitempty"`
	State    *float64 `json:"state,omitempty"`
	Section  *string  `json:"section,omitempty"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseWorstGaps, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseWorstGaps) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseWorstGaps
	aux := struct {
		*plain
		National lenientNumber[float64] `json:"national"`
		State    lenientNumber[float64] `json:"state"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.National.assignPtr(&r.National)
	aux.State.assignPtr(&r.State)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseQuote: JURISDICTION mode: the base dossier quote.
type StorefrontAvailabilityResponseQuote struct {
	Price        StorefrontAvailabilityResponseQuotePrice    `json:"price"`
	Currency     string                                      `json:"currency"`
	AddOns       []StorefrontAvailabilityResponseQuoteAddOns `json:"add_ons"`
	Pay          []json.RawMessage                           `json:"pay"`
	PurchaseFlow string                                      `json:"purchase_flow"`
}

// StorefrontAvailabilityResponseQuotePrice is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseQuotePrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// StorefrontAvailabilityResponseQuoteAddOns is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseQuoteAddOns struct {
	Code        *string                                        `json:"code,omitempty"`
	Label       *string                                        `json:"label,omitempty"`
	Price       StorefrontAvailabilityResponseQuoteAddOnsPrice `json:"price"`
	Available   *bool                                          `json:"available,omitempty"`
	Purchasable *bool                                          `json:"purchasable,omitempty"`
	Note        *string                                        `json:"note,omitempty"`
}

// StorefrontAvailabilityResponseQuoteAddOnsPrice is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseQuoteAddOnsPrice struct {
	Amount   *string `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// StorefrontAvailabilityResponseWarts is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseWarts struct {
	Dataset         *string  `json:"dataset,omitempty"`
	Surface         *string  `json:"surface,omitempty"`
	Status          *string  `json:"status,omitempty"`
	ObservedMaxDate *string  `json:"observed_max_date,omitempty"`
	AgeDays         *float64 `json:"age_days,omitempty"`
	MaxAgeDays      *int64   `json:"max_age_days,omitempty"`
	ProbedAt        *string  `json:"probed_at,omitempty"`
	Reason          *string  `json:"reason,omitempty"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseWarts, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseWarts) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseWarts
	aux := struct {
		*plain
		AgeDays    lenientNumber[float64] `json:"age_days"`
		MaxAgeDays lenientNumber[int64]   `json:"max_age_days"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.AgeDays.assignPtr(&r.AgeDays)
	aux.MaxAgeDays.assignPtr(&r.MaxAgeDays)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseCounties: JURISDICTION mode: the state's counties by parcel count.
type StorefrontAvailabilityResponseCounties struct {
	Total    int64                                        `json:"total"`
	Returned float64                                      `json:"returned"`
	Rows     []StorefrontAvailabilityResponseCountiesRows `json:"rows"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseCounties, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseCounties) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseCounties
	aux := struct {
		*plain
		Total    lenientNumber[int64]   `json:"total"`
		Returned lenientNumber[float64] `json:"returned"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Total.assign(&r.Total)
	aux.Returned.assign(&r.Returned)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseCountiesRows is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseCountiesRows struct {
	StateFIPS  string                                             `json:"state_fips"`
	CountyFIPS string                                             `json:"county_fips"`
	State      *string                                            `json:"state,omitempty"`
	Name       *string                                            `json:"name,omitempty"`
	Parcels    *float64                                           `json:"parcels,omitempty"`
	Coverage   StorefrontAvailabilityResponseCountiesRowsCoverage `json:"coverage"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseCountiesRows, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseCountiesRows) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseCountiesRows
	aux := struct {
		*plain
		Parcels lenientNumber[float64] `json:"parcels"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Parcels.assignPtr(&r.Parcels)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseCountiesRowsCoverage is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseCountiesRowsCoverage struct {
	Address  *float64 `json:"address,omitempty"`
	Geocode  *float64 `json:"geocode,omitempty"`
	Owner    *float64 `json:"owner,omitempty"`
	Value    *float64 `json:"value,omitempty"`
	Geometry *float64 `json:"geometry,omitempty"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseCountiesRowsCoverage, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseCountiesRowsCoverage) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseCountiesRowsCoverage
	aux := struct {
		*plain
		Address  lenientNumber[float64] `json:"address"`
		Geocode  lenientNumber[float64] `json:"geocode"`
		Owner    lenientNumber[float64] `json:"owner"`
		Value    lenientNumber[float64] `json:"value"`
		Geometry lenientNumber[float64] `json:"geometry"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Address.assignPtr(&r.Address)
	aux.Geocode.assignPtr(&r.Geocode)
	aux.Owner.assignPtr(&r.Owner)
	aux.Value.assignPtr(&r.Value)
	aux.Geometry.assignPtr(&r.Geometry)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseMode: 'parcel' in PARCEL mode; absent in JURISDICTION mode.
//
// It is a string; the StorefrontAvailabilityResponseMode* constants list the documented values.
type StorefrontAvailabilityResponseMode = string

// Documented values of StorefrontAvailabilityResponseMode.
const (
	StorefrontAvailabilityResponseModeParcel StorefrontAvailabilityResponseMode = "parcel"
)

// StorefrontAvailabilityResponseParcel: PARCEL mode: the resolved parcel identity.
type StorefrontAvailabilityResponseParcel struct {
	CanonicalID string `json:"canonical_id"`
	StateFIPS   string `json:"state_fips"`
	CountyFIPS  string `json:"county_fips"`
	CountyFIPS5 string `json:"county_fips_5"`
	ParcelID    string `json:"parcel_id"`
}

// StorefrontAvailabilityResponseDossierQuote: PARCEL mode: the value-tiered quote for this parcel.
type StorefrontAvailabilityResponseDossierQuote struct {
	Basis    string  `json:"basis"`
	Price    Money   `json:"price"`
	PriceUsd float64 `json:"price_usd"`

	// What the x402 402 advertises as maxAmountRequired (USDC atomic, 6dp).
	PriceAtomicUsdc string `json:"price_atomic_usdc"`
	Asset           string `json:"asset"`

	// Asset-class band (residential|multifamily|premium) derived from assessed value.
	Band StorefrontAvailabilityResponseDossierQuoteBand `json:"band"`

	// The auditable multipliers: base, V (asset value), R (data richness), F (freshness).
	Breakdown StorefrontAvailabilityResponseDossierQuoteBreakdown `json:"breakdown"`

	// The cheap signals the multipliers were derived from.
	Signals       StorefrontAvailabilityResponseDossierQuoteSignals       `json:"signals"`
	GeometryAddOn StorefrontAvailabilityResponseDossierQuoteGeometryAddOn `json:"geometry_add_on"`
	Pay           []*string                                               `json:"pay"`
	Note          string                                                  `json:"note"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseDossierQuote, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseDossierQuote) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseDossierQuote
	aux := struct {
		*plain
		PriceUsd lenientNumber[float64] `json:"price_usd"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PriceUsd.assign(&r.PriceUsd)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseDossierQuoteBand: Asset-class band
// (residential|multifamily|premium) derived from assessed value.
type StorefrontAvailabilityResponseDossierQuoteBand struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Range string `json:"range"`
}

// StorefrontAvailabilityResponseDossierQuoteBreakdown: The auditable multipliers: base, V (asset
// value), R (data richness), F (freshness).
type StorefrontAvailabilityResponseDossierQuoteBreakdown struct {
	Base float64 `json:"base"`
	V    float64 `json:"V"`
	R    float64 `json:"R"`
	F    float64 `json:"F"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseDossierQuoteBreakdown, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseDossierQuoteBreakdown) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseDossierQuoteBreakdown
	aux := struct {
		*plain
		Base lenientNumber[float64] `json:"base"`
		V    lenientNumber[float64] `json:"V"`
		R    lenientNumber[float64] `json:"R"`
		F    lenientNumber[float64] `json:"F"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Base.assign(&r.Base)
	aux.V.assign(&r.V)
	aux.R.assign(&r.R)
	aux.F.assign(&r.F)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseDossierQuoteSignals: The cheap signals the multipliers were
// derived from.
type StorefrontAvailabilityResponseDossierQuoteSignals struct {
	AssessedValue   float64 `json:"assessed_value"`
	PopulatedFields float64 `json:"populated_fields"`
	HasDeeds        bool    `json:"has_deeds"`
	HasPermits      bool    `json:"has_permits"`
	Freshness       string  `json:"freshness"`
}

// UnmarshalJSON decodes StorefrontAvailabilityResponseDossierQuoteSignals, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *StorefrontAvailabilityResponseDossierQuoteSignals) UnmarshalJSON(data []byte) error {
	type plain StorefrontAvailabilityResponseDossierQuoteSignals
	aux := struct {
		*plain
		AssessedValue   lenientNumber[float64] `json:"assessed_value"`
		PopulatedFields lenientNumber[float64] `json:"populated_fields"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.AssessedValue.assign(&r.AssessedValue)
	aux.PopulatedFields.assign(&r.PopulatedFields)
	return softTypeError(err)
}

// StorefrontAvailabilityResponseDossierQuoteGeometryAddOn is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseDossierQuoteGeometryAddOn struct {
	Code        string                                                       `json:"code"`
	Label       string                                                       `json:"label"`
	Price       StorefrontAvailabilityResponseDossierQuoteGeometryAddOnPrice `json:"price"`
	Purchasable bool                                                         `json:"purchasable"`
}

// StorefrontAvailabilityResponseDossierQuoteGeometryAddOnPrice is generated from the OpenAPI spec.
type StorefrontAvailabilityResponseDossierQuoteGeometryAddOnPrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
