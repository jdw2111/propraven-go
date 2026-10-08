// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.

package propraven

import (
	"encoding/json"
)

// AssessmentHistoryRecord is generated from the OpenAPI spec.
type AssessmentHistoryRecord struct {
	// Source-stated year; null means unknown. Never inferred from a snapshot or capture date.
	AssessmentYear *int64 `json:"assessment_year,omitempty"`

	// Source-stated year; null means unknown. Never inferred from a snapshot or capture date.
	TaxYear *int64 `json:"tax_year,omitempty"`

	// Snapshot vintage year, not an assessment year.
	VintageYear      *int64                            `json:"vintage_year,omitempty"`
	TotalValue       *float64                          `json:"total_value,omitempty"`
	LandValue        *float64                          `json:"land_value,omitempty"`
	ImprovementValue *float64                          `json:"improvement_value,omitempty"`
	TaxAmount        *float64                          `json:"tax_amount,omitempty"`
	TaxPaidAmount    *float64                          `json:"tax_paid_amount,omitempty"`
	Vintage          *string                           `json:"vintage,omitempty"`
	SourceURL        *string                           `json:"source_url,omitempty"`
	SourceAsOf       *string                           `json:"source_as_of,omitempty"`
	ValueBasis       AssessmentHistoryRecordValueBasis `json:"value_basis"`
	Source           string                            `json:"source"`
}

// UnmarshalJSON decodes AssessmentHistoryRecord, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *AssessmentHistoryRecord) UnmarshalJSON(data []byte) error {
	type plain AssessmentHistoryRecord
	aux := struct {
		*plain
		AssessmentYear   lenientNumber[int64]   `json:"assessment_year"`
		TaxYear          lenientNumber[int64]   `json:"tax_year"`
		VintageYear      lenientNumber[int64]   `json:"vintage_year"`
		TotalValue       lenientNumber[float64] `json:"total_value"`
		LandValue        lenientNumber[float64] `json:"land_value"`
		ImprovementValue lenientNumber[float64] `json:"improvement_value"`
		TaxAmount        lenientNumber[float64] `json:"tax_amount"`
		TaxPaidAmount    lenientNumber[float64] `json:"tax_paid_amount"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.AssessmentYear.assignPtr(&r.AssessmentYear)
	aux.TaxYear.assignPtr(&r.TaxYear)
	aux.VintageYear.assignPtr(&r.VintageYear)
	aux.TotalValue.assignPtr(&r.TotalValue)
	aux.LandValue.assignPtr(&r.LandValue)
	aux.ImprovementValue.assignPtr(&r.ImprovementValue)
	aux.TaxAmount.assignPtr(&r.TaxAmount)
	aux.TaxPaidAmount.assignPtr(&r.TaxPaidAmount)
	return softTypeError(err)
}

// AssessmentHistoryRecordValueBasis is generated from the OpenAPI spec. It is a string; the
// AssessmentHistoryRecordValueBasis* constants list the documented values.
type AssessmentHistoryRecordValueBasis = string

// Documented values of AssessmentHistoryRecordValueBasis.
const (
	AssessmentHistoryRecordValueBasisAssessed  AssessmentHistoryRecordValueBasis = "assessed"
	AssessmentHistoryRecordValueBasisAppraised AssessmentHistoryRecordValueBasis = "appraised"
	AssessmentHistoryRecordValueBasisMarket    AssessmentHistoryRecordValueBasis = "market"
	AssessmentHistoryRecordValueBasisTaxable   AssessmentHistoryRecordValueBasis = "taxable"
)

// AssessmentHistory is generated from the OpenAPI spec.
type AssessmentHistory struct {
	CanonicalID string                    `json:"canonical_id"`
	Status      AssessmentHistoryStatus   `json:"status"`
	Records     []AssessmentHistoryRecord `json:"records"`
	Coverage    AssessmentHistoryCoverage `json:"coverage"`
}

// AssessmentHistoryStatus is generated from the OpenAPI spec. It is a string; the
// AssessmentHistoryStatus* constants list the documented values.
type AssessmentHistoryStatus = string

// Documented values of AssessmentHistoryStatus.
const (
	AssessmentHistoryStatusOk    AssessmentHistoryStatus = "ok"
	AssessmentHistoryStatusEmpty AssessmentHistoryStatus = "empty"
)

// AssessmentHistoryCoverage is generated from the OpenAPI spec.
type AssessmentHistoryCoverage struct {
	AssessmentYears []int64 `json:"assessment_years"`
	TaxYears        []int64 `json:"tax_years"`
	RecordCount     int64   `json:"record_count"`
	Truncated       bool    `json:"truncated"`
	Limit           int64   `json:"limit"`
	Note            string  `json:"note"`
	SourceProduct   string  `json:"source_product"`
	SourceVersion   string  `json:"source_version"`
}

// UnmarshalJSON decodes AssessmentHistoryCoverage, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *AssessmentHistoryCoverage) UnmarshalJSON(data []byte) error {
	type plain AssessmentHistoryCoverage
	aux := struct {
		*plain
		AssessmentYears lenientSlice[int64]  `json:"assessment_years"`
		TaxYears        lenientSlice[int64]  `json:"tax_years"`
		RecordCount     lenientNumber[int64] `json:"record_count"`
		Limit           lenientNumber[int64] `json:"limit"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.AssessmentYears.assign(&r.AssessmentYears)
	aux.TaxYears.assign(&r.TaxYears)
	aux.RecordCount.assign(&r.RecordCount)
	aux.Limit.assign(&r.Limit)
	return softTypeError(err)
}

// Parcel is generated from the OpenAPI spec.
type Parcel struct {
	// PropRaven parcel UUID. Accepted by GET /parcels/{id}.
	ID string `json:"id"`

	// 3-digit within-state county FIPS code (the 5-digit form is `state_fips` + `county_fips`).
	CountyFIPS string `json:"county_fips"`

	// 2-digit state FIPS code.
	StateFIPS string `json:"state_fips"`

	// County-assigned parcel identifier (APN as the county publishes it). The canonical id is
	// `state_fips:county_fips:parcel_id`.
	ParcelID          string  `json:"parcel_id"`
	Address           *string `json:"address,omitempty"`
	NormalizedAddress *string `json:"normalized_address,omitempty"`
	City              *string `json:"city,omitempty"`

	// State FIPS as a number (legacy duplicate of `state_fips`).
	State                    *int64   `json:"state,omitempty"`
	Zip                      *string  `json:"zip,omitempty"`
	Zip5                     *string  `json:"zip5,omitempty"`
	ZipPlus4                 *string  `json:"zip_plus4,omitempty"`
	Latitude                 *float64 `json:"latitude,omitempty"`
	Longitude                *float64 `json:"longitude,omitempty"`
	LandUseCode              *string  `json:"land_use_code,omitempty"`
	LandUseDesc              *string  `json:"land_use_desc,omitempty"`
	TotalAssessedValue       *float64 `json:"total_assessed_value,omitempty"`
	LandAssessedValue        *float64 `json:"land_assessed_value,omitempty"`
	ImprovementAssessedValue *float64 `json:"improvement_assessed_value,omitempty"`
	LastSalePrice            *float64 `json:"last_sale_price,omitempty"`
	LastSaleDate             *string  `json:"last_sale_date,omitempty"`
	MarketValue              *float64 `json:"market_value,omitempty"`
	AvmValue                 *float64 `json:"avm_value,omitempty"`
	AvmConfidence            *string  `json:"avm_confidence,omitempty"`
	AvmMethod                *string  `json:"avm_method,omitempty"`
	TaxAmount                *float64 `json:"tax_amount,omitempty"`
	TaxYear                  *int64   `json:"tax_year,omitempty"`

	// Composite deal opportunity score from 0 to 100.
	DealScore                   *float64           `json:"deal_score,omitempty"`
	PricePerSqft                *float64           `json:"price_per_sqft,omitempty"`
	CompSalePrice               *float64           `json:"comp_sale_price,omitempty"`
	CompSaleDate                *string            `json:"comp_sale_date,omitempty"`
	CompSimilarityScore         *float64           `json:"comp_similarity_score,omitempty"`
	EstimatedMonthlyRent        *float64           `json:"estimated_monthly_rent,omitempty"`
	EstimatedAnnualRent         *float64           `json:"estimated_annual_rent,omitempty"`
	RentYieldPct                *float64           `json:"rent_yield_pct,omitempty"`
	RentalConfidence            *string            `json:"rental_confidence,omitempty"`
	GrossRentMultiplier         *float64           `json:"gross_rent_multiplier,omitempty"`
	LandImprovementRatio        *float64           `json:"land_improvement_ratio,omitempty"`
	ImprovementToLandRatio      *float64           `json:"improvement_to_land_ratio,omitempty"`
	IsRedevelopmentCandidate    *bool              `json:"is_redevelopment_candidate,omitempty"`
	BuildingSqft                *float64           `json:"building_sqft,omitempty"`
	YearBuilt                   *int64             `json:"year_built,omitempty"`
	LotSizeAcres                *float64           `json:"lot_size_acres,omitempty"`
	LotSizeSqft                 *float64           `json:"lot_size_sqft,omitempty"`
	Bedrooms                    *int64             `json:"bedrooms,omitempty"`
	Bathrooms                   *float64           `json:"bathrooms,omitempty"`
	Stories                     *float64           `json:"stories,omitempty"`
	Units                       *int64             `json:"units,omitempty"`
	UnitCount                   *int64             `json:"unit_count,omitempty"`
	ConstructionType            *string            `json:"construction_type,omitempty"`
	OwnerName                   *string            `json:"owner_name,omitempty"`
	OwnerAddress                *string            `json:"owner_address,omitempty"`
	OwnerCity                   *string            `json:"owner_city,omitempty"`
	OwnerState                  *string            `json:"owner_state,omitempty"`
	OwnerZip                    *string            `json:"owner_zip,omitempty"`
	OwnershipType               *string            `json:"ownership_type,omitempty"`
	OwnerEntityType             *string            `json:"owner_entity_type,omitempty"`
	EntityType                  *string            `json:"entity_type,omitempty"`
	IsEntityOwned               *bool              `json:"is_entity_owned,omitempty"`
	IsAbsentee                  *bool              `json:"is_absentee,omitempty"`
	IsPeAggregator              *bool              `json:"is_pe_aggregator,omitempty"`
	OwnerOccupiedFlag           *bool              `json:"owner_occupied_flag,omitempty"`
	DataQualityScore            *float64           `json:"data_quality_score,omitempty"`
	FloodZone                   *string            `json:"flood_zone,omitempty"`
	IsSfha                      *bool              `json:"is_sfha,omitempty"`
	IsOpportunityZone           *bool              `json:"is_opportunity_zone,omitempty"`
	IsJustice40                 *bool              `json:"is_justice40,omitempty"`
	IsFlip                      *bool              `json:"is_flip,omitempty"`
	DeedCount                   *int64             `json:"deed_count,omitempty"`
	PermitCount                 *int64             `json:"permit_count,omitempty"`
	PermitCount12mo             *int64             `json:"permit_count_12mo,omitempty"`
	Zoning                      *string            `json:"zoning,omitempty"`
	CountyName                  *string            `json:"county_name,omitempty"`
	PropertyType                *string            `json:"property_type,omitempty"`
	Guards                      []string           `json:"_guards"`
	IsFlipBasis                 string             `json:"is_flip_basis"`
	LandImprovementRatioStored  *float64           `json:"land_improvement_ratio_stored,omitempty"`
	LandImprovementRatioBasis   string             `json:"land_improvement_ratio_basis"`
	ImprovementToLandRatioBasis string             `json:"improvement_to_land_ratio_basis"`
	IsAbsenteeBasis             string             `json:"is_absentee_basis"`
	OwnerStateNorm              *string            `json:"owner_state_norm,omitempty"`
	IsOutOfState                *bool              `json:"is_out_of_state,omitempty"`
	IdentityGate                ParcelIdentityGate `json:"identity_gate"`
	SiteGate                    ParcelSiteGate     `json:"site_gate"`
	DerivedGate                 ParcelDerivedGate  `json:"derived_gate"`
	RentGate                    ParcelRentGate     `json:"rent_gate"`
	CountyNameBasis             string             `json:"county_name_basis"`
	AvmMethodFamily             *string            `json:"avm_method_family,omitempty"`
	AvmMethodBasis              string             `json:"avm_method_basis"`
	IsFlipDominantSharePct      *float64           `json:"is_flip_dominant_share_pct,omitempty"`
	PricePerSqftBasis           *string            `json:"price_per_sqft_basis,omitempty"`
}

// UnmarshalJSON decodes Parcel, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *Parcel) UnmarshalJSON(data []byte) error {
	type plain Parcel
	aux := struct {
		*plain
		State                      lenientNumber[int64]   `json:"state"`
		Latitude                   lenientNumber[float64] `json:"latitude"`
		Longitude                  lenientNumber[float64] `json:"longitude"`
		TotalAssessedValue         lenientNumber[float64] `json:"total_assessed_value"`
		LandAssessedValue          lenientNumber[float64] `json:"land_assessed_value"`
		ImprovementAssessedValue   lenientNumber[float64] `json:"improvement_assessed_value"`
		LastSalePrice              lenientNumber[float64] `json:"last_sale_price"`
		MarketValue                lenientNumber[float64] `json:"market_value"`
		AvmValue                   lenientNumber[float64] `json:"avm_value"`
		TaxAmount                  lenientNumber[float64] `json:"tax_amount"`
		TaxYear                    lenientNumber[int64]   `json:"tax_year"`
		DealScore                  lenientNumber[float64] `json:"deal_score"`
		PricePerSqft               lenientNumber[float64] `json:"price_per_sqft"`
		CompSalePrice              lenientNumber[float64] `json:"comp_sale_price"`
		CompSimilarityScore        lenientNumber[float64] `json:"comp_similarity_score"`
		EstimatedMonthlyRent       lenientNumber[float64] `json:"estimated_monthly_rent"`
		EstimatedAnnualRent        lenientNumber[float64] `json:"estimated_annual_rent"`
		RentYieldPct               lenientNumber[float64] `json:"rent_yield_pct"`
		GrossRentMultiplier        lenientNumber[float64] `json:"gross_rent_multiplier"`
		LandImprovementRatio       lenientNumber[float64] `json:"land_improvement_ratio"`
		ImprovementToLandRatio     lenientNumber[float64] `json:"improvement_to_land_ratio"`
		BuildingSqft               lenientNumber[float64] `json:"building_sqft"`
		YearBuilt                  lenientNumber[int64]   `json:"year_built"`
		LotSizeAcres               lenientNumber[float64] `json:"lot_size_acres"`
		LotSizeSqft                lenientNumber[float64] `json:"lot_size_sqft"`
		Bedrooms                   lenientNumber[int64]   `json:"bedrooms"`
		Bathrooms                  lenientNumber[float64] `json:"bathrooms"`
		Stories                    lenientNumber[float64] `json:"stories"`
		Units                      lenientNumber[int64]   `json:"units"`
		UnitCount                  lenientNumber[int64]   `json:"unit_count"`
		DataQualityScore           lenientNumber[float64] `json:"data_quality_score"`
		DeedCount                  lenientNumber[int64]   `json:"deed_count"`
		PermitCount                lenientNumber[int64]   `json:"permit_count"`
		PermitCount12mo            lenientNumber[int64]   `json:"permit_count_12mo"`
		LandImprovementRatioStored lenientNumber[float64] `json:"land_improvement_ratio_stored"`
		IsFlipDominantSharePct     lenientNumber[float64] `json:"is_flip_dominant_share_pct"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.State.assignPtr(&r.State)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.LandAssessedValue.assignPtr(&r.LandAssessedValue)
	aux.ImprovementAssessedValue.assignPtr(&r.ImprovementAssessedValue)
	aux.LastSalePrice.assignPtr(&r.LastSalePrice)
	aux.MarketValue.assignPtr(&r.MarketValue)
	aux.AvmValue.assignPtr(&r.AvmValue)
	aux.TaxAmount.assignPtr(&r.TaxAmount)
	aux.TaxYear.assignPtr(&r.TaxYear)
	aux.DealScore.assignPtr(&r.DealScore)
	aux.PricePerSqft.assignPtr(&r.PricePerSqft)
	aux.CompSalePrice.assignPtr(&r.CompSalePrice)
	aux.CompSimilarityScore.assignPtr(&r.CompSimilarityScore)
	aux.EstimatedMonthlyRent.assignPtr(&r.EstimatedMonthlyRent)
	aux.EstimatedAnnualRent.assignPtr(&r.EstimatedAnnualRent)
	aux.RentYieldPct.assignPtr(&r.RentYieldPct)
	aux.GrossRentMultiplier.assignPtr(&r.GrossRentMultiplier)
	aux.LandImprovementRatio.assignPtr(&r.LandImprovementRatio)
	aux.ImprovementToLandRatio.assignPtr(&r.ImprovementToLandRatio)
	aux.BuildingSqft.assignPtr(&r.BuildingSqft)
	aux.YearBuilt.assignPtr(&r.YearBuilt)
	aux.LotSizeAcres.assignPtr(&r.LotSizeAcres)
	aux.LotSizeSqft.assignPtr(&r.LotSizeSqft)
	aux.Bedrooms.assignPtr(&r.Bedrooms)
	aux.Bathrooms.assignPtr(&r.Bathrooms)
	aux.Stories.assignPtr(&r.Stories)
	aux.Units.assignPtr(&r.Units)
	aux.UnitCount.assignPtr(&r.UnitCount)
	aux.DataQualityScore.assignPtr(&r.DataQualityScore)
	aux.DeedCount.assignPtr(&r.DeedCount)
	aux.PermitCount.assignPtr(&r.PermitCount)
	aux.PermitCount12mo.assignPtr(&r.PermitCount12mo)
	aux.LandImprovementRatioStored.assignPtr(&r.LandImprovementRatioStored)
	aux.IsFlipDominantSharePct.assignPtr(&r.IsFlipDominantSharePct)
	return softTypeError(err)
}

// ParcelIdentityGate is generated from the OpenAPI spec.
type ParcelIdentityGate struct {
	Applied              *bool     `json:"applied,omitempty"`
	Suppressed           []*string `json:"suppressed"`
	Reason               *string   `json:"reason,omitempty"`
	JoinKeyBasis         string    `json:"join_key_basis"`
	TwinsInOtherCounties *int64    `json:"twins_in_other_counties,omitempty"`
	CoordsBasis          string    `json:"coords_basis"`
	Unmeasured           []*string `json:"unmeasured"`
	ServedUnderDoubt     []*string `json:"served_under_doubt"`
	Note                 *string   `json:"note,omitempty"`
}

// UnmarshalJSON decodes ParcelIdentityGate, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelIdentityGate) UnmarshalJSON(data []byte) error {
	type plain ParcelIdentityGate
	aux := struct {
		*plain
		TwinsInOtherCounties lenientNumber[int64] `json:"twins_in_other_counties"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TwinsInOtherCounties.assignPtr(&r.TwinsInOtherCounties)
	return softTypeError(err)
}

// ParcelSiteGate is generated from the OpenAPI spec.
type ParcelSiteGate struct {
	Applied      *bool     `json:"applied,omitempty"`
	Suppressed   []*string `json:"suppressed"`
	Reason       *string   `json:"reason,omitempty"`
	Labelled     []*string `json:"labelled"`
	MailingBasis string    `json:"mailing_basis"`
	Note         *string   `json:"note,omitempty"`
}

// ParcelDerivedGate is generated from the OpenAPI spec.
type ParcelDerivedGate struct {
	Applied     *bool             `json:"applied,omitempty"`
	Suppressed  []*string         `json:"suppressed"`
	Recomputed  []*string         `json:"recomputed"`
	Unevaluated []json.RawMessage `json:"unevaluated"`
	Reason      *string           `json:"reason,omitempty"`
	Note        *string           `json:"note,omitempty"`
}

// ParcelRentGate is generated from the OpenAPI spec.
type ParcelRentGate struct {
	Applied          *bool     `json:"applied,omitempty"`
	Suppressed       []*string `json:"suppressed"`
	Reason           *string   `json:"reason,omitempty"`
	RentalConfidence *string   `json:"rental_confidence,omitempty"`
	Note             *string   `json:"note,omitempty"`
}

// Owner is generated from the OpenAPI spec.
type Owner struct {
	OwnerNameNormalized *string         `json:"owner_name_normalized,omitempty"`
	PropertyCount       *int64          `json:"property_count,omitempty"`
	StateCount          *int64          `json:"state_count,omitempty"`
	CountyCount         *int64          `json:"county_count,omitempty"`
	TotalAssessedValue  *float64        `json:"total_assessed_value,omitempty"`
	AvgAssessedValue    *float64        `json:"avg_assessed_value,omitempty"`
	TotalAcreage        *float64        `json:"total_acreage,omitempty"`
	StatesList          json.RawMessage `json:"states_list,omitempty"`
	PortfolioRank       *int64          `json:"portfolio_rank,omitempty"`

	// Observed values include: CORP.
	EntityType      *string  `json:"entity_type,omitempty"`
	IsPeAggregator  *bool    `json:"is_pe_aggregator,omitempty"`
	IsEntityOwned   *bool    `json:"is_entity_owned,omitempty"`
	PeParentGroup   *string  `json:"pe_parent_group,omitempty"`
	PeSponsor       *string  `json:"pe_sponsor,omitempty"`
	PeConfidence    *float64 `json:"pe_confidence,omitempty"`
	OwnerEntityType *string  `json:"owner_entity_type,omitempty"`
	IsAbsentee      *bool    `json:"is_absentee,omitempty"`
	ContactCard     *string  `json:"contact_card,omitempty"`
	OwnerName       *string  `json:"owner_name,omitempty"`
	States          []string `json:"states,omitempty"`
}

// UnmarshalJSON decodes Owner, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *Owner) UnmarshalJSON(data []byte) error {
	type plain Owner
	aux := struct {
		*plain
		PropertyCount      lenientNumber[int64]   `json:"property_count"`
		StateCount         lenientNumber[int64]   `json:"state_count"`
		CountyCount        lenientNumber[int64]   `json:"county_count"`
		TotalAssessedValue lenientNumber[float64] `json:"total_assessed_value"`
		AvgAssessedValue   lenientNumber[float64] `json:"avg_assessed_value"`
		TotalAcreage       lenientNumber[float64] `json:"total_acreage"`
		PortfolioRank      lenientNumber[int64]   `json:"portfolio_rank"`
		PeConfidence       lenientNumber[float64] `json:"pe_confidence"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PropertyCount.assignPtr(&r.PropertyCount)
	aux.StateCount.assignPtr(&r.StateCount)
	aux.CountyCount.assignPtr(&r.CountyCount)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.AvgAssessedValue.assignPtr(&r.AvgAssessedValue)
	aux.TotalAcreage.assignPtr(&r.TotalAcreage)
	aux.PortfolioRank.assignPtr(&r.PortfolioRank)
	aux.PeConfidence.assignPtr(&r.PeConfidence)
	return softTypeError(err)
}

// Permit is generated from the OpenAPI spec.
type Permit struct {
	PermitNumber      *string       `json:"permit_number,omitempty"`
	PermitType        *string       `json:"permit_type,omitempty"`
	PermitStatus      *string       `json:"permit_status,omitempty"`
	Description       *string       `json:"description,omitempty"`
	WorkClass         *string       `json:"work_class,omitempty"`
	IssuedDate        *string       `json:"issued_date,omitempty"`
	CompletedDate     *string       `json:"completed_date,omitempty"`
	EstimatedCost     *float64      `json:"estimated_cost,omitempty"`
	FeeAmount         *float64      `json:"fee_amount,omitempty"`
	ContractorName    *string       `json:"contractor_name,omitempty"`
	ContractorLicense *string       `json:"contractor_license,omitempty"`
	SiteAddress       *string       `json:"site_address,omitempty"`
	City              *string       `json:"city,omitempty"`
	Type              *string       `json:"type,omitempty"`
	Status            *PermitStatus `json:"status,omitempty"`
	Contractor        *string       `json:"contractor,omitempty"`
}

// UnmarshalJSON decodes Permit, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *Permit) UnmarshalJSON(data []byte) error {
	type plain Permit
	aux := struct {
		*plain
		EstimatedCost lenientNumber[float64] `json:"estimated_cost"`
		FeeAmount     lenientNumber[float64] `json:"fee_amount"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.EstimatedCost.assignPtr(&r.EstimatedCost)
	aux.FeeAmount.assignPtr(&r.FeeAmount)
	return softTypeError(err)
}

// PermitStatus is generated from the OpenAPI spec. It is a string; the PermitStatus* constants
// list the documented values.
type PermitStatus = string

// Documented values of PermitStatus.
const (
	PermitStatusIssued    PermitStatus = "issued"
	PermitStatusPending   PermitStatus = "pending"
	PermitStatusApproved  PermitStatus = "approved"
	PermitStatusExpired   PermitStatus = "expired"
	PermitStatusCompleted PermitStatus = "completed"
	PermitStatusDenied    PermitStatus = "denied"
)

// Deed is generated from the OpenAPI spec.
type Deed struct {
	DocumentNumber   *string  `json:"document_number,omitempty"`
	RecordingDate    *string  `json:"recording_date,omitempty"`
	SaleDate         *string  `json:"sale_date,omitempty"`
	DocumentType     *string  `json:"document_type,omitempty"`
	SalePrice        *float64 `json:"sale_price,omitempty"`
	Consideration    *float64 `json:"consideration,omitempty"`
	GrantorName      *string  `json:"grantor_name,omitempty"`
	GranteeName      *string  `json:"grantee_name,omitempty"`
	GrantorName2     *string  `json:"grantor_name_2,omitempty"`
	GranteeName2     *string  `json:"grantee_name_2,omitempty"`
	GrantorAddress   *string  `json:"grantor_address,omitempty"`
	GranteeAddress   *string  `json:"grantee_address,omitempty"`
	GrantorType      *string  `json:"grantor_type,omitempty"`
	GranteeType      *string  `json:"grantee_type,omitempty"`
	Book             *string  `json:"book,omitempty"`
	Page             *string  `json:"page,omitempty"`
	InstrumentNumber *string  `json:"instrument_number,omitempty"`
	TransferTax      *float64 `json:"transfer_tax,omitempty"`
	ExciseTax        *float64 `json:"excise_tax,omitempty"`
	SaleType         *string  `json:"sale_type,omitempty"`
	IsArmLength      *bool    `json:"is_arm_length,omitempty"`
	LegalDescription *string  `json:"legal_description,omitempty"`
	Lot              *string  `json:"lot,omitempty"`
	Block            *string  `json:"block,omitempty"`
	Subdivision      *string  `json:"subdivision,omitempty"`
	SourceURL        *string  `json:"source_url,omitempty"`
	PropertyAddress  *string  `json:"property_address,omitempty"`
	PropertyCity     *string  `json:"property_city,omitempty"`
	PropertyState    *string  `json:"property_state,omitempty"`
	DeedType         *string  `json:"deed_type,omitempty"`
}

// UnmarshalJSON decodes Deed, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *Deed) UnmarshalJSON(data []byte) error {
	type plain Deed
	aux := struct {
		*plain
		SalePrice     lenientNumber[float64] `json:"sale_price"`
		Consideration lenientNumber[float64] `json:"consideration"`
		TransferTax   lenientNumber[float64] `json:"transfer_tax"`
		ExciseTax     lenientNumber[float64] `json:"excise_tax"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.SalePrice.assignPtr(&r.SalePrice)
	aux.Consideration.assignPtr(&r.Consideration)
	aux.TransferTax.assignPtr(&r.TransferTax)
	aux.ExciseTax.assignPtr(&r.ExciseTax)
	return softTypeError(err)
}

// RiskAssessment is generated from the OpenAPI spec.
type RiskAssessment struct {
	FloodZone    *string                    `json:"flood_zone,omitempty"`
	IsSfha       *bool                      `json:"is_sfha,omitempty"`
	IsSfhaBasis  string                     `json:"is_sfha_basis"`
	Seismic      RiskAssessmentSeismic      `json:"seismic"`
	Windstorm    RiskAssessmentWindstorm    `json:"windstorm"`
	Wildfire     RiskAssessmentWildfire     `json:"wildfire"`
	AirQuality   RiskAssessmentAirQuality   `json:"air_quality"`
	Crime        RiskAssessmentCrime        `json:"crime"`
	WithholdGate RiskAssessmentWithholdGate `json:"withhold_gate"`
	IdentityGate RiskAssessmentIdentityGate `json:"identity_gate"`
}

// RiskAssessmentSeismic is generated from the OpenAPI spec.
type RiskAssessmentSeismic struct {
	Ss                  *float64 `json:"ss,omitempty"`
	S1                  *float64 `json:"s1,omitempty"`
	Sds                 *float64 `json:"sds,omitempty"`
	Sd1                 *float64 `json:"sd1,omitempty"`
	Sdc                 *string  `json:"sdc,omitempty"`
	PgaG                *float64 `json:"pga_g,omitempty"`
	NriEarthquakeRating *string  `json:"nri_earthquake_rating,omitempty"`
}

// UnmarshalJSON decodes RiskAssessmentSeismic, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *RiskAssessmentSeismic) UnmarshalJSON(data []byte) error {
	type plain RiskAssessmentSeismic
	aux := struct {
		*plain
		Ss   lenientNumber[float64] `json:"ss"`
		S1   lenientNumber[float64] `json:"s1"`
		Sds  lenientNumber[float64] `json:"sds"`
		Sd1  lenientNumber[float64] `json:"sd1"`
		PgaG lenientNumber[float64] `json:"pga_g"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Ss.assignPtr(&r.Ss)
	aux.S1.assignPtr(&r.S1)
	aux.Sds.assignPtr(&r.Sds)
	aux.Sd1.assignPtr(&r.Sd1)
	aux.PgaG.assignPtr(&r.PgaG)
	return softTypeError(err)
}

// RiskAssessmentWindstorm is generated from the OpenAPI spec.
type RiskAssessmentWindstorm struct {
	HurricaneRating        *string  `json:"hurricane_rating,omitempty"`
	TornadoRating          *string  `json:"tornado_rating,omitempty"`
	PowerWindMw            *float64 `json:"power_wind_mw,omitempty"`
	StormScore             *float64 `json:"storm_score,omitempty"`
	StormTopEvent          *string  `json:"storm_top_event,omitempty"`
	StormEventCount30y     *int64   `json:"storm_event_count_30y,omitempty"`
	StormTornadoCount30y   *int64   `json:"storm_tornado_count_30y,omitempty"`
	StormHurricaneCount30y *int64   `json:"storm_hurricane_count_30y,omitempty"`
	StormHailCount30y      *int64   `json:"storm_hail_count_30y,omitempty"`
	StormPropertyDmgUsd30y *float64 `json:"storm_property_dmg_usd_30y,omitempty"`
	StormDeaths30y         *float64 `json:"storm_deaths_30y,omitempty"`
}

// UnmarshalJSON decodes RiskAssessmentWindstorm, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *RiskAssessmentWindstorm) UnmarshalJSON(data []byte) error {
	type plain RiskAssessmentWindstorm
	aux := struct {
		*plain
		PowerWindMw            lenientNumber[float64] `json:"power_wind_mw"`
		StormScore             lenientNumber[float64] `json:"storm_score"`
		StormEventCount30y     lenientNumber[int64]   `json:"storm_event_count_30y"`
		StormTornadoCount30y   lenientNumber[int64]   `json:"storm_tornado_count_30y"`
		StormHurricaneCount30y lenientNumber[int64]   `json:"storm_hurricane_count_30y"`
		StormHailCount30y      lenientNumber[int64]   `json:"storm_hail_count_30y"`
		StormPropertyDmgUsd30y lenientNumber[float64] `json:"storm_property_dmg_usd_30y"`
		StormDeaths30y         lenientNumber[float64] `json:"storm_deaths_30y"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PowerWindMw.assignPtr(&r.PowerWindMw)
	aux.StormScore.assignPtr(&r.StormScore)
	aux.StormEventCount30y.assignPtr(&r.StormEventCount30y)
	aux.StormTornadoCount30y.assignPtr(&r.StormTornadoCount30y)
	aux.StormHurricaneCount30y.assignPtr(&r.StormHurricaneCount30y)
	aux.StormHailCount30y.assignPtr(&r.StormHailCount30y)
	aux.StormPropertyDmgUsd30y.assignPtr(&r.StormPropertyDmgUsd30y)
	aux.StormDeaths30y.assignPtr(&r.StormDeaths30y)
	return softTypeError(err)
}

// RiskAssessmentWildfire is generated from the OpenAPI spec.
type RiskAssessmentWildfire struct {
	CountyName       *string  `json:"county_name,omitempty"`
	RiskNationalRank *float64 `json:"risk_national_rank,omitempty"`
	BpNationalRank   *float64 `json:"bp_national_rank,omitempty"`
	RiskStateRank    *float64 `json:"risk_state_rank,omitempty"`
	BpStateRank      *float64 `json:"bp_state_rank,omitempty"`
	RiskClass        *string  `json:"risk_class,omitempty"`

	// Annual burn probability as a decimal.
	BurnProbability *float64 `json:"burn_probability,omitempty"`
}

// UnmarshalJSON decodes RiskAssessmentWildfire, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *RiskAssessmentWildfire) UnmarshalJSON(data []byte) error {
	type plain RiskAssessmentWildfire
	aux := struct {
		*plain
		RiskNationalRank lenientNumber[float64] `json:"risk_national_rank"`
		BpNationalRank   lenientNumber[float64] `json:"bp_national_rank"`
		RiskStateRank    lenientNumber[float64] `json:"risk_state_rank"`
		BpStateRank      lenientNumber[float64] `json:"bp_state_rank"`
		BurnProbability  lenientNumber[float64] `json:"burn_probability"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.RiskNationalRank.assignPtr(&r.RiskNationalRank)
	aux.BpNationalRank.assignPtr(&r.BpNationalRank)
	aux.RiskStateRank.assignPtr(&r.RiskStateRank)
	aux.BpStateRank.assignPtr(&r.BpStateRank)
	aux.BurnProbability.assignPtr(&r.BurnProbability)
	return softTypeError(err)
}

// RiskAssessmentAirQuality is generated from the OpenAPI spec.
type RiskAssessmentAirQuality struct {
	Year *int64 `json:"year,omitempty"`

	// Median Air Quality Index value.
	MedianAqi     *float64 `json:"median_aqi,omitempty"`
	MaxAqi        *float64 `json:"max_aqi,omitempty"`
	GoodDays      *int64   `json:"good_days,omitempty"`
	ModerateDays  *int64   `json:"moderate_days,omitempty"`
	UnhealthyDays *int64   `json:"unhealthy_days,omitempty"`
	Category      *string  `json:"category,omitempty"`
}

// UnmarshalJSON decodes RiskAssessmentAirQuality, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *RiskAssessmentAirQuality) UnmarshalJSON(data []byte) error {
	type plain RiskAssessmentAirQuality
	aux := struct {
		*plain
		Year          lenientNumber[int64]   `json:"year"`
		MedianAqi     lenientNumber[float64] `json:"median_aqi"`
		MaxAqi        lenientNumber[float64] `json:"max_aqi"`
		GoodDays      lenientNumber[int64]   `json:"good_days"`
		ModerateDays  lenientNumber[int64]   `json:"moderate_days"`
		UnhealthyDays lenientNumber[int64]   `json:"unhealthy_days"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Year.assignPtr(&r.Year)
	aux.MedianAqi.assignPtr(&r.MedianAqi)
	aux.MaxAqi.assignPtr(&r.MaxAqi)
	aux.GoodDays.assignPtr(&r.GoodDays)
	aux.ModerateDays.assignPtr(&r.ModerateDays)
	aux.UnhealthyDays.assignPtr(&r.UnhealthyDays)
	return softTypeError(err)
}

// RiskAssessmentCrime is generated from the OpenAPI spec.
type RiskAssessmentCrime struct {
	// Crime score from 0 (low) to 100 (high).
	Score *float64 `json:"score,omitempty"`

	// Observed values include: 3.
	Tier      *float64 `json:"tier,omitempty"`
	TierLabel *string  `json:"tier_label,omitempty"`

	// Observed values include: improving.
	Trend             *string  `json:"trend,omitempty"`
	ViolentCrimeRate  *float64 `json:"violent_crime_rate,omitempty"`
	PropertyCrimeRate *float64 `json:"property_crime_rate,omitempty"`
}

// UnmarshalJSON decodes RiskAssessmentCrime, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *RiskAssessmentCrime) UnmarshalJSON(data []byte) error {
	type plain RiskAssessmentCrime
	aux := struct {
		*plain
		Score             lenientNumber[float64] `json:"score"`
		Tier              lenientNumber[float64] `json:"tier"`
		ViolentCrimeRate  lenientNumber[float64] `json:"violent_crime_rate"`
		PropertyCrimeRate lenientNumber[float64] `json:"property_crime_rate"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Score.assignPtr(&r.Score)
	aux.Tier.assignPtr(&r.Tier)
	aux.ViolentCrimeRate.assignPtr(&r.ViolentCrimeRate)
	aux.PropertyCrimeRate.assignPtr(&r.PropertyCrimeRate)
	return softTypeError(err)
}

// RiskAssessmentWithholdGate is generated from the OpenAPI spec.
type RiskAssessmentWithholdGate struct {
	Applied    *bool                             `json:"applied,omitempty"`
	Suppressed []*string                         `json:"suppressed"`
	Reason     *string                           `json:"reason,omitempty"`
	Reasons    RiskAssessmentWithholdGateReasons `json:"reasons"`
	WithheldOn *string                           `json:"withheld_on,omitempty"`
	Note       *string                           `json:"note,omitempty"`
}

// RiskAssessmentWithholdGateReasons is generated from the OpenAPI spec.
type RiskAssessmentWithholdGateReasons struct {
	PowerWindMw            *string `json:"power_wind_mw,omitempty"`
	StormHurricaneCount30y *string `json:"storm_hurricane_count_30y,omitempty"`
}

// RiskAssessmentIdentityGate is generated from the OpenAPI spec.
type RiskAssessmentIdentityGate struct {
	Applied              *bool             `json:"applied,omitempty"`
	Suppressed           []json.RawMessage `json:"suppressed"`
	Reason               *string           `json:"reason,omitempty"`
	JoinKeyBasis         string            `json:"join_key_basis"`
	TwinsInOtherCounties *int64            `json:"twins_in_other_counties,omitempty"`
	CoordsBasis          *string           `json:"coords_basis,omitempty"`
	Unmeasured           []json.RawMessage `json:"unmeasured"`
	ServedUnderDoubt     []json.RawMessage `json:"served_under_doubt"`
	Note                 *string           `json:"note,omitempty"`
}

// UnmarshalJSON decodes RiskAssessmentIdentityGate, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *RiskAssessmentIdentityGate) UnmarshalJSON(data []byte) error {
	type plain RiskAssessmentIdentityGate
	aux := struct {
		*plain
		TwinsInOtherCounties lenientNumber[int64] `json:"twins_in_other_counties"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TwinsInOtherCounties.assignPtr(&r.TwinsInOtherCounties)
	return softTypeError(err)
}

// Problem: THE error body of every /api/v1 endpoint (RFC 7807 problem details, served as
// `application/problem+json`). `code` is the stable machine-readable identifier (e.g.
// invalid_parameter, authentication_required, account_required, monthly_cap_reached,
// query_timeout, not_found, method_not_allowed); `detail` is human-readable and never contains
// database or driver text. Validation failures add `errors: [{param, message}]`. `request_id`
// identifies the request for support. Per-code members (e.g. `reason`, `retry_after`, `used` /
// `limit` / `plan` / `upgrade`, `allow`) sit beside the core members. Exception: an x402 `402
// Payment Required` keeps the x402 protocol envelope (`x402Version`, `accepts`).
type Problem struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int64  `json:"status"`
	Detail string `json:"detail"`

	// Stable machine-readable code.
	Code string `json:"code"`

	// Finer reason for `code`, e.g. people_data_requires_account.
	Reason *string `json:"reason,omitempty"`

	// Per-parameter validation failures (400 invalid_parameter).
	Errors []ProblemErrors `json:"errors,omitempty"`

	// Support handle for this request.
	RequestID *string `json:"request_id,omitempty"`
}

// UnmarshalJSON decodes Problem, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *Problem) UnmarshalJSON(data []byte) error {
	type plain Problem
	aux := struct {
		*plain
		Status lenientNumber[int64] `json:"status"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Status.assign(&r.Status)
	return softTypeError(err)
}

// ProblemErrors is generated from the OpenAPI spec.
type ProblemErrors struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

// ErrorSchema: Deprecated name for `Problem` (kept so existing references resolve). Every error is
// a `Problem`.
//
// Deprecated: the API marks this deprecated.
type ErrorSchema = Problem

// Contractor is generated from the OpenAPI spec.
type Contractor struct {
	ContractorNameNormalized *string  `json:"contractor_name_normalized,omitempty"`
	ContractorLicense        *string  `json:"contractor_license,omitempty"`
	PermitCount              *int64   `json:"permit_count,omitempty"`
	JurisdictionCount        *int64   `json:"jurisdiction_count,omitempty"`
	StateCount               *int64   `json:"state_count,omitempty"`
	CountyCount              *int64   `json:"county_count,omitempty"`
	TotalPermitValue         *int64   `json:"total_permit_value,omitempty"`
	AvgPermitValue           *float64 `json:"avg_permit_value,omitempty"`
	FirstPermitDate          *string  `json:"first_permit_date,omitempty"`
	LastPermitDate           *string  `json:"last_permit_date,omitempty"`
	ActiveYears              *int64   `json:"active_years,omitempty"`
	TopPermitTypes           *string  `json:"top_permit_types,omitempty"`

	// Comma-delimited 2-letter state codes.
	StatesList *string `json:"states_list,omitempty"`

	// National rank, 1 = highest activity.
	ContractorRank *int64 `json:"contractor_rank,omitempty"`
}

// UnmarshalJSON decodes Contractor, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *Contractor) UnmarshalJSON(data []byte) error {
	type plain Contractor
	aux := struct {
		*plain
		PermitCount       lenientNumber[int64]   `json:"permit_count"`
		JurisdictionCount lenientNumber[int64]   `json:"jurisdiction_count"`
		StateCount        lenientNumber[int64]   `json:"state_count"`
		CountyCount       lenientNumber[int64]   `json:"county_count"`
		TotalPermitValue  lenientNumber[int64]   `json:"total_permit_value"`
		AvgPermitValue    lenientNumber[float64] `json:"avg_permit_value"`
		ActiveYears       lenientNumber[int64]   `json:"active_years"`
		ContractorRank    lenientNumber[int64]   `json:"contractor_rank"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PermitCount.assignPtr(&r.PermitCount)
	aux.JurisdictionCount.assignPtr(&r.JurisdictionCount)
	aux.StateCount.assignPtr(&r.StateCount)
	aux.CountyCount.assignPtr(&r.CountyCount)
	aux.TotalPermitValue.assignPtr(&r.TotalPermitValue)
	aux.AvgPermitValue.assignPtr(&r.AvgPermitValue)
	aux.ActiveYears.assignPtr(&r.ActiveYears)
	aux.ContractorRank.assignPtr(&r.ContractorRank)
	return softTypeError(err)
}

// EntityOwnedParcel is generated from the OpenAPI spec.
type EntityOwnedParcel struct {
	CountyFIPS         string   `json:"county_fips"`
	StateFIPS          string   `json:"state_fips"`
	ParcelID           string   `json:"parcel_id"`
	OwnerName          *string  `json:"owner_name,omitempty"`
	EntityType         *string  `json:"entity_type,omitempty"`
	Address            *string  `json:"address,omitempty"`
	City               *string  `json:"city,omitempty"`
	State              *string  `json:"state,omitempty"`
	Zip                *string  `json:"zip,omitempty"`
	TotalAssessedValue *int64   `json:"total_assessed_value,omitempty"`
	LotSizeAcres       *float64 `json:"lot_size_acres,omitempty"`
	YearBuilt          *int64   `json:"year_built,omitempty"`
	Zoning             *string  `json:"zoning,omitempty"`
	LandUseDesc        *string  `json:"land_use_desc,omitempty"`
	OwnerAddress       *string  `json:"owner_address,omitempty"`
	OwnerCity          *string  `json:"owner_city,omitempty"`
	OwnerState         *string  `json:"owner_state,omitempty"`
	Latitude           *float64 `json:"latitude,omitempty"`
	Longitude          *float64 `json:"longitude,omitempty"`
}

// UnmarshalJSON decodes EntityOwnedParcel, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *EntityOwnedParcel) UnmarshalJSON(data []byte) error {
	type plain EntityOwnedParcel
	aux := struct {
		*plain
		TotalAssessedValue lenientNumber[int64]   `json:"total_assessed_value"`
		LotSizeAcres       lenientNumber[float64] `json:"lot_size_acres"`
		YearBuilt          lenientNumber[int64]   `json:"year_built"`
		Latitude           lenientNumber[float64] `json:"latitude"`
		Longitude          lenientNumber[float64] `json:"longitude"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.LotSizeAcres.assignPtr(&r.LotSizeAcres)
	aux.YearBuilt.assignPtr(&r.YearBuilt)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	return softTypeError(err)
}

// EntityAggregate is generated from the OpenAPI spec.
type EntityAggregate struct {
	OwnerName   *string  `json:"owner_name,omitempty"`
	EntityType  *string  `json:"entity_type,omitempty"`
	ParcelCount *int64   `json:"parcel_count,omitempty"`
	TotalValue  *float64 `json:"total_value,omitempty"`
	StatesArr   []string `json:"states_arr,omitempty"`
}

// UnmarshalJSON decodes EntityAggregate, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *EntityAggregate) UnmarshalJSON(data []byte) error {
	type plain EntityAggregate
	aux := struct {
		*plain
		ParcelCount lenientNumber[int64]   `json:"parcel_count"`
		TotalValue  lenientNumber[float64] `json:"total_value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ParcelCount.assignPtr(&r.ParcelCount)
	aux.TotalValue.assignPtr(&r.TotalValue)
	return softTypeError(err)
}

// EntitySummary is generated from the OpenAPI spec.
type EntitySummary struct {
	TotalEntities *int64 `json:"total_entities,omitempty"`
	TotalParcels  *int64 `json:"total_parcels,omitempty"`
	LLCCount      *int64 `json:"llc_count,omitempty"`
	CorpCount     *int64 `json:"corp_count,omitempty"`
	TrustCount    *int64 `json:"trust_count,omitempty"`
	LpCount       *int64 `json:"lp_count,omitempty"`
}

// UnmarshalJSON decodes EntitySummary, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *EntitySummary) UnmarshalJSON(data []byte) error {
	type plain EntitySummary
	aux := struct {
		*plain
		TotalEntities lenientNumber[int64] `json:"total_entities"`
		TotalParcels  lenientNumber[int64] `json:"total_parcels"`
		LLCCount      lenientNumber[int64] `json:"llc_count"`
		CorpCount     lenientNumber[int64] `json:"corp_count"`
		TrustCount    lenientNumber[int64] `json:"trust_count"`
		LpCount       lenientNumber[int64] `json:"lp_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalEntities.assignPtr(&r.TotalEntities)
	aux.TotalParcels.assignPtr(&r.TotalParcels)
	aux.LLCCount.assignPtr(&r.LLCCount)
	aux.CorpCount.assignPtr(&r.CorpCount)
	aux.TrustCount.assignPtr(&r.TrustCount)
	aux.LpCount.assignPtr(&r.LpCount)
	return softTypeError(err)
}

// HighLandRatioParcel is generated from the OpenAPI spec.
type HighLandRatioParcel struct {
	CountyFIPS               string   `json:"county_fips"`
	StateFIPS                string   `json:"state_fips"`
	ParcelID                 string   `json:"parcel_id"`
	OwnerName                *string  `json:"owner_name,omitempty"`
	Address                  *string  `json:"address,omitempty"`
	City                     *string  `json:"city,omitempty"`
	State                    *string  `json:"state,omitempty"`
	LandAssessedValue        *float64 `json:"land_assessed_value,omitempty"`
	ImprovementAssessedValue *float64 `json:"improvement_assessed_value,omitempty"`
	TotalAssessedValue       *int64   `json:"total_assessed_value,omitempty"`

	// land / improvement, higher = more redevelopment potential.
	LandImprovementRatio *float64 `json:"land_improvement_ratio,omitempty"`
	LotSizeAcres         *float64 `json:"lot_size_acres,omitempty"`
	YearBuilt            *int64   `json:"year_built,omitempty"`
	Zoning               *string  `json:"zoning,omitempty"`
	LandUseDesc          *string  `json:"land_use_desc,omitempty"`
	Latitude             *float64 `json:"latitude,omitempty"`
	Longitude            *float64 `json:"longitude,omitempty"`
}

// UnmarshalJSON decodes HighLandRatioParcel, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *HighLandRatioParcel) UnmarshalJSON(data []byte) error {
	type plain HighLandRatioParcel
	aux := struct {
		*plain
		LandAssessedValue        lenientNumber[float64] `json:"land_assessed_value"`
		ImprovementAssessedValue lenientNumber[float64] `json:"improvement_assessed_value"`
		TotalAssessedValue       lenientNumber[int64]   `json:"total_assessed_value"`
		LandImprovementRatio     lenientNumber[float64] `json:"land_improvement_ratio"`
		LotSizeAcres             lenientNumber[float64] `json:"lot_size_acres"`
		YearBuilt                lenientNumber[int64]   `json:"year_built"`
		Latitude                 lenientNumber[float64] `json:"latitude"`
		Longitude                lenientNumber[float64] `json:"longitude"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.LandAssessedValue.assignPtr(&r.LandAssessedValue)
	aux.ImprovementAssessedValue.assignPtr(&r.ImprovementAssessedValue)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.LandImprovementRatio.assignPtr(&r.LandImprovementRatio)
	aux.LotSizeAcres.assignPtr(&r.LotSizeAcres)
	aux.YearBuilt.assignPtr(&r.YearBuilt)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	return softTypeError(err)
}

// Lender is generated from the OpenAPI spec.
type Lender struct {
	LenderNameNormalized *string  `json:"lender_name_normalized,omitempty"`
	MortgageCount        *int64   `json:"mortgage_count,omitempty"`
	TotalMortgageVolume  *float64 `json:"total_mortgage_volume,omitempty"`
	AvgMortgageAmount    *float64 `json:"avg_mortgage_amount,omitempty"`
	MedianMortgageAmount *float64 `json:"median_mortgage_amount,omitempty"`
	CountyCount          *int64   `json:"county_count,omitempty"`
	StateCount           *int64   `json:"state_count,omitempty"`
	StatesList           *string  `json:"states_list,omitempty"`
	FirstMortgageDate    *string  `json:"first_mortgage_date,omitempty"`
	LastMortgageDate     *string  `json:"last_mortgage_date,omitempty"`

	// National rank, 1 = highest volume.
	LenderRank *int64 `json:"lender_rank,omitempty"`
}

// UnmarshalJSON decodes Lender, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *Lender) UnmarshalJSON(data []byte) error {
	type plain Lender
	aux := struct {
		*plain
		MortgageCount        lenientNumber[int64]   `json:"mortgage_count"`
		TotalMortgageVolume  lenientNumber[float64] `json:"total_mortgage_volume"`
		AvgMortgageAmount    lenientNumber[float64] `json:"avg_mortgage_amount"`
		MedianMortgageAmount lenientNumber[float64] `json:"median_mortgage_amount"`
		CountyCount          lenientNumber[int64]   `json:"county_count"`
		StateCount           lenientNumber[int64]   `json:"state_count"`
		LenderRank           lenientNumber[int64]   `json:"lender_rank"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.MortgageCount.assignPtr(&r.MortgageCount)
	aux.TotalMortgageVolume.assignPtr(&r.TotalMortgageVolume)
	aux.AvgMortgageAmount.assignPtr(&r.AvgMortgageAmount)
	aux.MedianMortgageAmount.assignPtr(&r.MedianMortgageAmount)
	aux.CountyCount.assignPtr(&r.CountyCount)
	aux.StateCount.assignPtr(&r.StateCount)
	aux.LenderRank.assignPtr(&r.LenderRank)
	return softTypeError(err)
}

// LongHoldParcel is generated from the OpenAPI spec.
type LongHoldParcel struct {
	CountyFIPS         string   `json:"county_fips"`
	StateFIPS          string   `json:"state_fips"`
	ParcelID           string   `json:"parcel_id"`
	OwnerName          *string  `json:"owner_name,omitempty"`
	Address            *string  `json:"address,omitempty"`
	City               *string  `json:"city,omitempty"`
	State              *string  `json:"state,omitempty"`
	Zip                *string  `json:"zip,omitempty"`
	TotalAssessedValue *int64   `json:"total_assessed_value,omitempty"`
	LastSaleDate       *string  `json:"last_sale_date,omitempty"`
	LastSalePrice      *float64 `json:"last_sale_price,omitempty"`
	YearsHeld          *float64 `json:"years_held,omitempty"`
	YearBuilt          *int64   `json:"year_built,omitempty"`
	BuildingAge        *int64   `json:"building_age,omitempty"`
	LotSizeAcres       *float64 `json:"lot_size_acres,omitempty"`
	HoldTier           *string  `json:"hold_tier,omitempty"`
	LandUseDesc        *string  `json:"land_use_desc,omitempty"`
	Latitude           *float64 `json:"latitude,omitempty"`
	Longitude          *float64 `json:"longitude,omitempty"`
}

// UnmarshalJSON decodes LongHoldParcel, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LongHoldParcel) UnmarshalJSON(data []byte) error {
	type plain LongHoldParcel
	aux := struct {
		*plain
		TotalAssessedValue lenientNumber[int64]   `json:"total_assessed_value"`
		LastSalePrice      lenientNumber[float64] `json:"last_sale_price"`
		YearsHeld          lenientNumber[float64] `json:"years_held"`
		YearBuilt          lenientNumber[int64]   `json:"year_built"`
		BuildingAge        lenientNumber[int64]   `json:"building_age"`
		LotSizeAcres       lenientNumber[float64] `json:"lot_size_acres"`
		Latitude           lenientNumber[float64] `json:"latitude"`
		Longitude          lenientNumber[float64] `json:"longitude"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.LastSalePrice.assignPtr(&r.LastSalePrice)
	aux.YearsHeld.assignPtr(&r.YearsHeld)
	aux.YearBuilt.assignPtr(&r.YearBuilt)
	aux.BuildingAge.assignPtr(&r.BuildingAge)
	aux.LotSizeAcres.assignPtr(&r.LotSizeAcres)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	return softTypeError(err)
}

// MarketSummary is generated from the OpenAPI spec.
type MarketSummary struct {
	CountyFIPS          string   `json:"county_fips"`
	Year                *int64   `json:"year,omitempty"`
	Quarter             *int64   `json:"quarter,omitempty"`
	TransactionCount    *int64   `json:"transaction_count,omitempty"`
	MedianPrice         *float64 `json:"median_price,omitempty"`
	AvgPrice            *float64 `json:"avg_price,omitempty"`
	TotalVolume         *int64   `json:"total_volume,omitempty"`
	CashSaleCount       *int64   `json:"cash_sale_count,omitempty"`
	CashSalePct         *float64 `json:"cash_sale_pct,omitempty"`
	MedianConsideration *float64 `json:"median_consideration,omitempty"`
	UniqueBuyers        *float64 `json:"unique_buyers,omitempty"`
	UniqueSellers       *float64 `json:"unique_sellers,omitempty"`
	RefreshedAt         *string  `json:"refreshed_at,omitempty"`
	StateFIPS           string   `json:"state_fips"`
	CountyName          *string  `json:"county_name,omitempty"`
	State               *string  `json:"state,omitempty"`
	SaleCount           *int64   `json:"sale_count,omitempty"`
	MedianSalePrice     *float64 `json:"median_sale_price,omitempty"`
	AvgSalePrice        *float64 `json:"avg_sale_price,omitempty"`
}

// UnmarshalJSON decodes MarketSummary, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketSummary) UnmarshalJSON(data []byte) error {
	type plain MarketSummary
	aux := struct {
		*plain
		Year                lenientNumber[int64]   `json:"year"`
		Quarter             lenientNumber[int64]   `json:"quarter"`
		TransactionCount    lenientNumber[int64]   `json:"transaction_count"`
		MedianPrice         lenientNumber[float64] `json:"median_price"`
		AvgPrice            lenientNumber[float64] `json:"avg_price"`
		TotalVolume         lenientNumber[int64]   `json:"total_volume"`
		CashSaleCount       lenientNumber[int64]   `json:"cash_sale_count"`
		CashSalePct         lenientNumber[float64] `json:"cash_sale_pct"`
		MedianConsideration lenientNumber[float64] `json:"median_consideration"`
		UniqueBuyers        lenientNumber[float64] `json:"unique_buyers"`
		UniqueSellers       lenientNumber[float64] `json:"unique_sellers"`
		SaleCount           lenientNumber[int64]   `json:"sale_count"`
		MedianSalePrice     lenientNumber[float64] `json:"median_sale_price"`
		AvgSalePrice        lenientNumber[float64] `json:"avg_sale_price"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Year.assignPtr(&r.Year)
	aux.Quarter.assignPtr(&r.Quarter)
	aux.TransactionCount.assignPtr(&r.TransactionCount)
	aux.MedianPrice.assignPtr(&r.MedianPrice)
	aux.AvgPrice.assignPtr(&r.AvgPrice)
	aux.TotalVolume.assignPtr(&r.TotalVolume)
	aux.CashSaleCount.assignPtr(&r.CashSaleCount)
	aux.CashSalePct.assignPtr(&r.CashSalePct)
	aux.MedianConsideration.assignPtr(&r.MedianConsideration)
	aux.UniqueBuyers.assignPtr(&r.UniqueBuyers)
	aux.UniqueSellers.assignPtr(&r.UniqueSellers)
	aux.SaleCount.assignPtr(&r.SaleCount)
	aux.MedianSalePrice.assignPtr(&r.MedianSalePrice)
	aux.AvgSalePrice.assignPtr(&r.AvgSalePrice)
	return softTypeError(err)
}

// AffordabilityRow is generated from the OpenAPI spec.
type AffordabilityRow struct {
	CountyFIPS             string   `json:"county_fips"`
	MedianHouseholdIncome  *float64 `json:"median_household_income,omitempty"`
	MedianSalePrice        *float64 `json:"median_sale_price,omitempty"`
	PriceToIncomeRatio     *float64 `json:"price_to_income_ratio,omitempty"`
	AffordabilityRating    *string  `json:"affordability_rating,omitempty"`
	Year                   *int64   `json:"year,omitempty"`
	RefreshedAt            *string  `json:"refreshed_at,omitempty"`
	StateFIPS              *string  `json:"state_fips,omitempty"`
	CountyName             *string  `json:"county_name,omitempty"`
	MonthlyPaymentEstimate *float64 `json:"monthly_payment_estimate,omitempty"`
	PctIncomeForHousing    *float64 `json:"pct_income_for_housing,omitempty"`
}

// UnmarshalJSON decodes AffordabilityRow, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *AffordabilityRow) UnmarshalJSON(data []byte) error {
	type plain AffordabilityRow
	aux := struct {
		*plain
		MedianHouseholdIncome  lenientNumber[float64] `json:"median_household_income"`
		MedianSalePrice        lenientNumber[float64] `json:"median_sale_price"`
		PriceToIncomeRatio     lenientNumber[float64] `json:"price_to_income_ratio"`
		Year                   lenientNumber[int64]   `json:"year"`
		MonthlyPaymentEstimate lenientNumber[float64] `json:"monthly_payment_estimate"`
		PctIncomeForHousing    lenientNumber[float64] `json:"pct_income_for_housing"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.MedianHouseholdIncome.assignPtr(&r.MedianHouseholdIncome)
	aux.MedianSalePrice.assignPtr(&r.MedianSalePrice)
	aux.PriceToIncomeRatio.assignPtr(&r.PriceToIncomeRatio)
	aux.Year.assignPtr(&r.Year)
	aux.MonthlyPaymentEstimate.assignPtr(&r.MonthlyPaymentEstimate)
	aux.PctIncomeForHousing.assignPtr(&r.PctIncomeForHousing)
	return softTypeError(err)
}

// PortfolioOwner is generated from the OpenAPI spec.
type PortfolioOwner struct {
	OwnerNameNormalized *string  `json:"owner_name_normalized,omitempty"`
	OwnerState          *string  `json:"owner_state,omitempty"`
	PropertyCount       *int64   `json:"property_count,omitempty"`
	StateCount          *int64   `json:"state_count,omitempty"`
	CountyCount         *int64   `json:"county_count,omitempty"`
	TotalAssessedValue  *int64   `json:"total_assessed_value,omitempty"`
	AvgAssessedValue    *float64 `json:"avg_assessed_value,omitempty"`
	TotalAcreage        *float64 `json:"total_acreage,omitempty"`
	StatesList          *string  `json:"states_list,omitempty"`

	// National rank, 1 = largest portfolio.
	PortfolioRank *int64 `json:"portfolio_rank,omitempty"`
}

// UnmarshalJSON decodes PortfolioOwner, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *PortfolioOwner) UnmarshalJSON(data []byte) error {
	type plain PortfolioOwner
	aux := struct {
		*plain
		PropertyCount      lenientNumber[int64]   `json:"property_count"`
		StateCount         lenientNumber[int64]   `json:"state_count"`
		CountyCount        lenientNumber[int64]   `json:"county_count"`
		TotalAssessedValue lenientNumber[int64]   `json:"total_assessed_value"`
		AvgAssessedValue   lenientNumber[float64] `json:"avg_assessed_value"`
		TotalAcreage       lenientNumber[float64] `json:"total_acreage"`
		PortfolioRank      lenientNumber[int64]   `json:"portfolio_rank"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PropertyCount.assignPtr(&r.PropertyCount)
	aux.StateCount.assignPtr(&r.StateCount)
	aux.CountyCount.assignPtr(&r.CountyCount)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.AvgAssessedValue.assignPtr(&r.AvgAssessedValue)
	aux.TotalAcreage.assignPtr(&r.TotalAcreage)
	aux.PortfolioRank.assignPtr(&r.PortfolioRank)
	return softTypeError(err)
}

// Webhook is generated from the OpenAPI spec.
type Webhook struct {
	ID *string `json:"id,omitempty"`

	// Customer endpoint. Must be https://.
	URL *string `json:"url,omitempty"`

	// First 14 chars of the secret (whsec_ + 8 hex). Use to identify the webhook in your dashboard;
	// full secret is shown only at create time.
	SecretPrefix        *string             `json:"secret_prefix,omitempty"`
	EventTypes          []WebhookEventTypes `json:"event_types,omitempty"`
	FilterKind          *WebhookFilterKind  `json:"filter_kind,omitempty"`
	FilterValue         WebhookFilter       `json:"filter_value,omitempty"`
	Description         *string             `json:"description,omitempty"`
	IsActive            *bool               `json:"is_active,omitempty"`
	CreatedAt           *string             `json:"created_at,omitempty"`
	DisabledAt          *string             `json:"disabled_at,omitempty"`
	DisabledReason      *string             `json:"disabled_reason,omitempty"`
	DeliveriesAttempted *int64              `json:"deliveries_attempted,omitempty"`
	DeliveriesSucceeded *int64              `json:"deliveries_succeeded,omitempty"`
	LastDeliveryAt      *string             `json:"last_delivery_at,omitempty"`
	LastSuccessAt       *string             `json:"last_success_at,omitempty"`
}

// UnmarshalJSON decodes Webhook, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *Webhook) UnmarshalJSON(data []byte) error {
	type plain Webhook
	aux := struct {
		*plain
		DeliveriesAttempted lenientNumber[int64] `json:"deliveries_attempted"`
		DeliveriesSucceeded lenientNumber[int64] `json:"deliveries_succeeded"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.DeliveriesAttempted.assignPtr(&r.DeliveriesAttempted)
	aux.DeliveriesSucceeded.assignPtr(&r.DeliveriesSucceeded)
	return softTypeError(err)
}

// WebhookEventTypes is generated from the OpenAPI spec. It is a string; the WebhookEventTypes*
// constants list the documented values.
type WebhookEventTypes = string

// Documented values of WebhookEventTypes.
const (
	WebhookEventTypesParcelSold         WebhookEventTypes = "parcel.sold"
	WebhookEventTypesParcelPermitFiled  WebhookEventTypes = "parcel.permit_filed"
	WebhookEventTypesParcelOwnerChanged WebhookEventTypes = "parcel.owner_changed"
)

// WebhookFilterKind is generated from the OpenAPI spec. It is a string; the WebhookFilterKind*
// constants list the documented values.
type WebhookFilterKind = string

// Documented values of WebhookFilterKind.
const (
	WebhookFilterKindParcelIDs  WebhookFilterKind = "parcel_ids"
	WebhookFilterKindStateFIPS  WebhookFilterKind = "state_fips"
	WebhookFilterKindCountyFIPS WebhookFilterKind = "county_fips"
)

// WebhookFilter: Shape varies with filter_kind. parcel_ids: explicit list. state_fips: all parcels
// in a state. county_fips: all parcels in a county within a state.
type WebhookFilter = json.RawMessage

// WebhookCreate is generated from the OpenAPI spec.
type WebhookCreate struct {
	// Customer endpoint. https:// only.
	URL string `json:"url"`

	// Event types to subscribe to. NOTE: only parcel.sold is live in v1.0; others 501.
	EventTypes  []WebhookCreateEventTypes `json:"event_types"`
	FilterKind  WebhookCreateFilterKind   `json:"filter_kind"`
	FilterValue WebhookFilter             `json:"filter_value"`

	// Optional human-readable label for your dashboard.
	Description *string `json:"description,omitempty"`
}

// WebhookCreateEventTypes is generated from the OpenAPI spec. It is a string; the
// WebhookCreateEventTypes* constants list the documented values.
type WebhookCreateEventTypes = string

// Documented values of WebhookCreateEventTypes.
const (
	WebhookCreateEventTypesParcelSold         WebhookCreateEventTypes = "parcel.sold"
	WebhookCreateEventTypesParcelPermitFiled  WebhookCreateEventTypes = "parcel.permit_filed"
	WebhookCreateEventTypesParcelOwnerChanged WebhookCreateEventTypes = "parcel.owner_changed"
)

// WebhookCreateFilterKind is generated from the OpenAPI spec. It is a string; the
// WebhookCreateFilterKind* constants list the documented values.
type WebhookCreateFilterKind = string

// Documented values of WebhookCreateFilterKind.
const (
	WebhookCreateFilterKindParcelIDs  WebhookCreateFilterKind = "parcel_ids"
	WebhookCreateFilterKindStateFIPS  WebhookCreateFilterKind = "state_fips"
	WebhookCreateFilterKindCountyFIPS WebhookCreateFilterKind = "county_fips"
)

// WebhookCreated is generated from the OpenAPI spec.
type WebhookCreated struct {
	ID *string `json:"id,omitempty"`

	// Customer endpoint. Must be https://.
	URL *string `json:"url,omitempty"`

	// First 14 chars of the secret (whsec_ + 8 hex). Use to identify the webhook in your dashboard;
	// full secret is shown only at create time.
	SecretPrefix        *string                    `json:"secret_prefix,omitempty"`
	EventTypes          []WebhookCreatedEventTypes `json:"event_types,omitempty"`
	FilterKind          *WebhookCreatedFilterKind  `json:"filter_kind,omitempty"`
	FilterValue         WebhookFilter              `json:"filter_value,omitempty"`
	Description         *string                    `json:"description,omitempty"`
	IsActive            *bool                      `json:"is_active,omitempty"`
	CreatedAt           *string                    `json:"created_at,omitempty"`
	DisabledAt          *string                    `json:"disabled_at,omitempty"`
	DisabledReason      *string                    `json:"disabled_reason,omitempty"`
	DeliveriesAttempted *int64                     `json:"deliveries_attempted,omitempty"`
	DeliveriesSucceeded *int64                     `json:"deliveries_succeeded,omitempty"`
	LastDeliveryAt      *string                    `json:"last_delivery_at,omitempty"`
	LastSuccessAt       *string                    `json:"last_success_at,omitempty"`

	// **Shown once.** Copy and store server-side immediately. Used to sign every outgoing delivery.
	Secret string `json:"secret"`

	// Signature-verification reminder.
	Hint string `json:"hint"`
}

// UnmarshalJSON decodes WebhookCreated, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *WebhookCreated) UnmarshalJSON(data []byte) error {
	type plain WebhookCreated
	aux := struct {
		*plain
		DeliveriesAttempted lenientNumber[int64] `json:"deliveries_attempted"`
		DeliveriesSucceeded lenientNumber[int64] `json:"deliveries_succeeded"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.DeliveriesAttempted.assignPtr(&r.DeliveriesAttempted)
	aux.DeliveriesSucceeded.assignPtr(&r.DeliveriesSucceeded)
	return softTypeError(err)
}

// WebhookCreatedEventTypes is generated from the OpenAPI spec. It is a string; the
// WebhookCreatedEventTypes* constants list the documented values.
type WebhookCreatedEventTypes = string

// Documented values of WebhookCreatedEventTypes.
const (
	WebhookCreatedEventTypesParcelSold         WebhookCreatedEventTypes = "parcel.sold"
	WebhookCreatedEventTypesParcelPermitFiled  WebhookCreatedEventTypes = "parcel.permit_filed"
	WebhookCreatedEventTypesParcelOwnerChanged WebhookCreatedEventTypes = "parcel.owner_changed"
)

// WebhookCreatedFilterKind is generated from the OpenAPI spec. It is a string; the
// WebhookCreatedFilterKind* constants list the documented values.
type WebhookCreatedFilterKind = string

// Documented values of WebhookCreatedFilterKind.
const (
	WebhookCreatedFilterKindParcelIDs  WebhookCreatedFilterKind = "parcel_ids"
	WebhookCreatedFilterKindStateFIPS  WebhookCreatedFilterKind = "state_fips"
	WebhookCreatedFilterKindCountyFIPS WebhookCreatedFilterKind = "county_fips"
)

// WebhookQuota is generated from the OpenAPI spec.
type WebhookQuota struct {
	// Max simultaneous active webhook endpoints on this tier.
	MaxEndpoints float64 `json:"maxEndpoints"`

	// Max event deliveries per UTC day on this tier.
	MaxEventsPerDay float64 `json:"maxEventsPerDay"`
}

// UnmarshalJSON decodes WebhookQuota, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *WebhookQuota) UnmarshalJSON(data []byte) error {
	type plain WebhookQuota
	aux := struct {
		*plain
		MaxEndpoints    lenientNumber[float64] `json:"maxEndpoints"`
		MaxEventsPerDay lenientNumber[float64] `json:"maxEventsPerDay"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.MaxEndpoints.assign(&r.MaxEndpoints)
	aux.MaxEventsPerDay.assign(&r.MaxEventsPerDay)
	return softTypeError(err)
}

// WebhookDelivery is generated from the OpenAPI spec.
type WebhookDelivery struct {
	ID *string `json:"id,omitempty"`

	// Deterministic event identifier — sha256(source || pk || event_type). Idempotent re-deliveries
	// share this.
	EventID            *string                   `json:"event_id,omitempty"`
	EventType          *WebhookDeliveryEventType `json:"event_type,omitempty"`
	EventOccurredAt    *string                   `json:"event_occurred_at,omitempty"`
	Status             *WebhookDeliveryStatus    `json:"status,omitempty"`
	Attempts           *int64                    `json:"attempts,omitempty"`
	LastAttemptAt      *string                   `json:"last_attempt_at,omitempty"`
	NextAttemptAt      *string                   `json:"next_attempt_at,omitempty"`
	LastResponseStatus *int64                    `json:"last_response_status,omitempty"`

	// Truncated to ~1KB.
	LastResponseBody *string `json:"last_response_body,omitempty"`
	LastError        *string `json:"last_error,omitempty"`
	DeadLetteredAt   *string `json:"dead_lettered_at,omitempty"`
	CreatedAt        *string `json:"created_at,omitempty"`
}

// UnmarshalJSON decodes WebhookDelivery, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *WebhookDelivery) UnmarshalJSON(data []byte) error {
	type plain WebhookDelivery
	aux := struct {
		*plain
		Attempts           lenientNumber[int64] `json:"attempts"`
		LastResponseStatus lenientNumber[int64] `json:"last_response_status"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Attempts.assignPtr(&r.Attempts)
	aux.LastResponseStatus.assignPtr(&r.LastResponseStatus)
	return softTypeError(err)
}

// WebhookDeliveryEventType is generated from the OpenAPI spec. It is a string; the
// WebhookDeliveryEventType* constants list the documented values.
type WebhookDeliveryEventType = string

// Documented values of WebhookDeliveryEventType.
const (
	WebhookDeliveryEventTypeParcelSold         WebhookDeliveryEventType = "parcel.sold"
	WebhookDeliveryEventTypeParcelPermitFiled  WebhookDeliveryEventType = "parcel.permit_filed"
	WebhookDeliveryEventTypeParcelOwnerChanged WebhookDeliveryEventType = "parcel.owner_changed"
)

// WebhookDeliveryStatus is generated from the OpenAPI spec. It is a string; the
// WebhookDeliveryStatus* constants list the documented values.
type WebhookDeliveryStatus = string

// Documented values of WebhookDeliveryStatus.
const (
	WebhookDeliveryStatusPending      WebhookDeliveryStatus = "pending"
	WebhookDeliveryStatusInFlight     WebhookDeliveryStatus = "in_flight"
	WebhookDeliveryStatusSucceeded    WebhookDeliveryStatus = "succeeded"
	WebhookDeliveryStatusFailed       WebhookDeliveryStatus = "failed"
	WebhookDeliveryStatusDeadLettered WebhookDeliveryStatus = "dead_lettered"
)

// ParcelSoldEvent is generated from the OpenAPI spec.
type ParcelSoldEvent struct {
	EventType string `json:"event_type"`

	// Deterministic. Use for idempotency.
	EventID    string `json:"event_id"`
	OccurredAt string `json:"occurred_at"`

	// Starts at 1; increments on retry.
	DeliveryAttempt int64 `json:"delivery_attempt"`

	// Composite county_fips:parcel_id.
	ParcelID   string  `json:"parcel_id"`
	StateFIPS  string  `json:"state_fips"`
	CountyFIPS string  `json:"county_fips"`
	SaleDate   *string `json:"sale_date,omitempty"`

	// May be null in non-disclosure states (KS, MS, TX, UT, WY, etc.).
	SalePriceUsd *float64 `json:"sale_price_usd,omitempty"`
	Grantor      *string  `json:"grantor,omitempty"`
	Grantee      *string  `json:"grantee,omitempty"`
	RecordedDate *string  `json:"recorded_date,omitempty"`

	// PropRaven ingest run that surfaced this event.
	SourceRunID *string `json:"source_run_id,omitempty"`
}

// UnmarshalJSON decodes ParcelSoldEvent, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelSoldEvent) UnmarshalJSON(data []byte) error {
	type plain ParcelSoldEvent
	aux := struct {
		*plain
		DeliveryAttempt lenientNumber[int64]   `json:"delivery_attempt"`
		SalePriceUsd    lenientNumber[float64] `json:"sale_price_usd"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.DeliveryAttempt.assign(&r.DeliveryAttempt)
	aux.SalePriceUsd.assignPtr(&r.SalePriceUsd)
	return softTypeError(err)
}

// AutocompleteResult is generated from the OpenAPI spec.
type AutocompleteResult struct {
	Locations       []AutocompleteLocation             `json:"locations"`
	Parcels         []AutocompleteParcel               `json:"parcels"`
	Addresses       []AutocompleteAddress              `json:"addresses"`
	Degraded        bool                               `json:"degraded"`
	OwnerNameSearch *AutocompleteResultOwnerNameSearch `json:"owner_name_search,omitempty"`
	PeopleFields    *AutocompleteResultPeopleFields    `json:"people_fields,omitempty"`
}

// AutocompleteLocation is generated from the OpenAPI spec.
type AutocompleteLocation struct {
	Name        *string  `json:"name,omitempty"`
	Type        *string  `json:"type,omitempty"`
	State       *string  `json:"state,omitempty"`
	City        *string  `json:"city,omitempty"`
	Lat         *float64 `json:"lat,omitempty"`
	Lng         *float64 `json:"lng,omitempty"`
	ParcelCount *int64   `json:"parcel_count,omitempty"`
}

// UnmarshalJSON decodes AutocompleteLocation, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *AutocompleteLocation) UnmarshalJSON(data []byte) error {
	type plain AutocompleteLocation
	aux := struct {
		*plain
		Lat         lenientNumber[float64] `json:"lat"`
		Lng         lenientNumber[float64] `json:"lng"`
		ParcelCount lenientNumber[int64]   `json:"parcel_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Lat.assignPtr(&r.Lat)
	aux.Lng.assignPtr(&r.Lng)
	aux.ParcelCount.assignPtr(&r.ParcelCount)
	return softTypeError(err)
}

// AutocompleteParcel is generated from the OpenAPI spec.
type AutocompleteParcel struct {
	ParcelID           string   `json:"parcel_id"`
	Address            *string  `json:"address,omitempty"`
	City               *string  `json:"city,omitempty"`
	StateFIPS          string   `json:"state_fips"`
	State              *string  `json:"state,omitempty"`
	CountyFIPS         string   `json:"county_fips"`
	OwnerName          *string  `json:"owner_name,omitempty"`
	Latitude           *float64 `json:"latitude,omitempty"`
	Longitude          *float64 `json:"longitude,omitempty"`
	TotalAssessedValue *int64   `json:"total_assessed_value,omitempty"`
	Type               *string  `json:"type,omitempty"`
}

// UnmarshalJSON decodes AutocompleteParcel, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *AutocompleteParcel) UnmarshalJSON(data []byte) error {
	type plain AutocompleteParcel
	aux := struct {
		*plain
		Latitude           lenientNumber[float64] `json:"latitude"`
		Longitude          lenientNumber[float64] `json:"longitude"`
		TotalAssessedValue lenientNumber[int64]   `json:"total_assessed_value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	return softTypeError(err)
}

// AutocompleteAddress: Mapbox-geocoded address suggestion. Use to disambiguate user input before
// calling /api/v1/lookup or /api/v1/parcels/{id}.
type AutocompleteAddress struct {
	Name  *string  `json:"name,omitempty"`
	Type  *string  `json:"type,omitempty"`
	Lat   *float64 `json:"lat,omitempty"`
	Lng   *float64 `json:"lng,omitempty"`
	South *float64 `json:"south,omitempty"`
	North *float64 `json:"north,omitempty"`
	West  *float64 `json:"west,omitempty"`
	East  *float64 `json:"east,omitempty"`
}

// UnmarshalJSON decodes AutocompleteAddress, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *AutocompleteAddress) UnmarshalJSON(data []byte) error {
	type plain AutocompleteAddress
	aux := struct {
		*plain
		Lat   lenientNumber[float64] `json:"lat"`
		Lng   lenientNumber[float64] `json:"lng"`
		South lenientNumber[float64] `json:"south"`
		North lenientNumber[float64] `json:"north"`
		West  lenientNumber[float64] `json:"west"`
		East  lenientNumber[float64] `json:"east"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Lat.assignPtr(&r.Lat)
	aux.Lng.assignPtr(&r.Lng)
	aux.South.assignPtr(&r.South)
	aux.North.assignPtr(&r.North)
	aux.West.assignPtr(&r.West)
	aux.East.assignPtr(&r.East)
	return softTypeError(err)
}

// AutocompleteResultOwnerNameSearch is generated from the OpenAPI spec.
type AutocompleteResultOwnerNameSearch struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Reason string `json:"reason"`
	Note   string `json:"note"`
}

// AutocompleteResultPeopleFields is generated from the OpenAPI spec.
type AutocompleteResultPeopleFields struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Reason string `json:"reason"`
	Note   string `json:"note"`
}

// FullSearchResult is generated from the OpenAPI spec.
type FullSearchResult struct {
	Results       []FullSearchResultResults  `json:"results"`
	Total         int64                      `json:"total"`
	TotalCapped   bool                       `json:"totalCapped"`
	Page          int64                      `json:"page"`
	Pages         int64                      `json:"pages"`
	HasMore       bool                       `json:"hasMore"`
	NextCursor    *string                    `json:"nextCursor,omitempty"`
	Warnings      []FullSearchResultWarnings `json:"warnings,omitempty"`
	TotalIsCapped *bool                      `json:"total_is_capped,omitempty"`
}

// UnmarshalJSON decodes FullSearchResult, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *FullSearchResult) UnmarshalJSON(data []byte) error {
	type plain FullSearchResult
	aux := struct {
		*plain
		Total lenientNumber[int64] `json:"total"`
		Page  lenientNumber[int64] `json:"page"`
		Pages lenientNumber[int64] `json:"pages"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Total.assign(&r.Total)
	aux.Page.assign(&r.Page)
	aux.Pages.assign(&r.Pages)
	return softTypeError(err)
}

// FullSearchResultResults is generated from the OpenAPI spec.
type FullSearchResultResults struct {
	// PropRaven parcel UUID (not the county APN; see `apn`). Pass it to GET /parcels/{id}.
	ParcelID    string   `json:"parcel_id"`
	APN         *string  `json:"apn,omitempty"`
	CountyFIPS  string   `json:"county_fips"`
	StateFIPS   string   `json:"state_fips"`
	SiteAddress *string  `json:"site_address,omitempty"`
	City        *string  `json:"city,omitempty"`
	Zip5        *string  `json:"zip5,omitempty"`
	OwnerName   *string  `json:"owner_name,omitempty"`
	TotalValue  *int64   `json:"total_value,omitempty"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
}

// UnmarshalJSON decodes FullSearchResultResults, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *FullSearchResultResults) UnmarshalJSON(data []byte) error {
	type plain FullSearchResultResults
	aux := struct {
		*plain
		TotalValue lenientNumber[int64]   `json:"total_value"`
		Latitude   lenientNumber[float64] `json:"latitude"`
		Longitude  lenientNumber[float64] `json:"longitude"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TotalValue.assignPtr(&r.TotalValue)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	return softTypeError(err)
}

// FullSearchResultWarnings is generated from the OpenAPI spec.
type FullSearchResultWarnings struct {
	Code    *string `json:"code,omitempty"`
	Param   *string `json:"param,omitempty"`
	Message *string `json:"message,omitempty"`
}

// ParcelGeoJSON: GeoJSON FeatureCollection of parcel polygons. Each feature's properties carry the
// basic parcel summary for popup rendering.
type ParcelGeoJSON struct {
	Type     string             `json:"type"`
	Features []ParcelGeoFeature `json:"features"`
}

// ParcelGeoFeature is generated from the OpenAPI spec.
type ParcelGeoFeature struct {
	Type       *string                    `json:"type,omitempty"`
	Geometry   ParcelGeoFeatureGeometry   `json:"geometry"`
	Properties ParcelGeoFeatureProperties `json:"properties"`
}

// ParcelGeoFeatureGeometry is generated from the OpenAPI spec.
type ParcelGeoFeatureGeometry struct {
	Type *string `json:"type,omitempty"`

	// GeoJSON polygon coordinate rings: outer ring first, then any inner rings.
	Coordinates [][][]*float64 `json:"coordinates"`
}

// ParcelGeoFeatureProperties is generated from the OpenAPI spec.
type ParcelGeoFeatureProperties struct {
	ID                 string   `json:"id"`
	CountyFIPS         string   `json:"county_fips"`
	Address            *string  `json:"address,omitempty"`
	City               *string  `json:"city,omitempty"`
	State              *string  `json:"state,omitempty"`
	Owner              *string  `json:"owner,omitempty"`
	Value              *float64 `json:"value,omitempty"`
	Type               *string  `json:"type,omitempty"`
	YearBuilt          *int64   `json:"year_built,omitempty"`
	ParcelID           *string  `json:"parcel_id,omitempty"`
	OwnerName          *string  `json:"owner_name,omitempty"`
	TotalAssessedValue *float64 `json:"total_assessed_value,omitempty"`
	PropertyType       *string  `json:"property_type,omitempty"`
	Latitude           *float64 `json:"latitude,omitempty"`
	Longitude          *float64 `json:"longitude,omitempty"`
}

// UnmarshalJSON decodes ParcelGeoFeatureProperties, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelGeoFeatureProperties) UnmarshalJSON(data []byte) error {
	type plain ParcelGeoFeatureProperties
	aux := struct {
		*plain
		Value              lenientNumber[float64] `json:"value"`
		YearBuilt          lenientNumber[int64]   `json:"year_built"`
		TotalAssessedValue lenientNumber[float64] `json:"total_assessed_value"`
		Latitude           lenientNumber[float64] `json:"latitude"`
		Longitude          lenientNumber[float64] `json:"longitude"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	aux.YearBuilt.assignPtr(&r.YearBuilt)
	aux.TotalAssessedValue.assignPtr(&r.TotalAssessedValue)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	return softTypeError(err)
}

// Money: A money amount as a decimal string plus its currency code.
type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// DossierSubSection: One real sub-table (deeds, comps, or permits) in a dossier, with the exact
// serving relation it came from and that relation's own watermark (each is NOT pinned to the
// parcels epoch).
type DossierSubSection struct {
	// The serving relation the rows came from.
	Relation *string `json:"relation,omitempty"`

	// Watermark for this relation.
	AsOf     *string                  `json:"as_of,omitempty"`
	Status   *DossierSubSectionStatus `json:"status,omitempty"`
	RowCount *int64                   `json:"row_count,omitempty"`
	Note     *string                  `json:"note,omitempty"`
	Rows     []map[string]any         `json:"rows,omitempty"`
}

// UnmarshalJSON decodes DossierSubSection, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *DossierSubSection) UnmarshalJSON(data []byte) error {
	type plain DossierSubSection
	aux := struct {
		*plain
		RowCount lenientNumber[int64] `json:"row_count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.RowCount.assignPtr(&r.RowCount)
	return softTypeError(err)
}

// DossierSubSectionStatus is generated from the OpenAPI spec. It is a string; the
// DossierSubSectionStatus* constants list the documented values.
type DossierSubSectionStatus = string

// Documented values of DossierSubSectionStatus.
const (
	DossierSubSectionStatusOk    DossierSubSectionStatus = "ok"
	DossierSubSectionStatusEmpty DossierSubSectionStatus = "empty"
)

// ParcelDossier: The paid, provenance-first parcel dossier returned by GET /parcels/{id}/report.
// Every populated field is delivered as `{name, value}`; the receipts (source, as_of, confidence,
// coverage) come as the opt-in `provenance` map (`include_provenance=true`), keyed by field name.
// The deeds/comps/permits sub-tables carry real rows; the GeoJSON boundary is a separately-priced
// add-on omitted from the base payload. Nulls/empties are dropped (see meta.field_count), not
// returned as null; columns a `sections`/`fields` projection leaves out are named in
// `meta.projection.omitted_fields`.
type ParcelDossier struct {
	// Parcel identity.
	Parcel *ParcelDossierParcel `json:"parcel,omitempty"`

	// Every populated field the delivery kept, as {name, value}. The receipt properties below are
	// present only on the pure assembler's unprojected output; on a delivered dossier they live in the
	// `provenance` map when `include_provenance=true` was requested.
	Fields []ParcelDossierFields `json:"fields,omitempty"`

	// Real sub-tables from the serving relations.
	Sections *ParcelDossierSections `json:"sections,omitempty"`

	// The GeoJSON boundary add-on. In the base dossier included=false and geometry is absent; a
	// phase-2 add-on purchase populates geometry.
	Boundary *ParcelDossierBoundary `json:"boundary,omitempty"`

	// Dossier-level provenance and pricing.
	Meta *ParcelDossierMeta `json:"meta,omitempty"`

	// Present only when the buyer has no account (wallet-only x402 or credit token): the people fields
	// in this dossier were withheld.
	PeopleFields *PeopleFieldsWithheld `json:"people_fields,omitempty"`
}

// ParcelDossierParcel: Parcel identity.
type ParcelDossierParcel struct {
	// state_fips:county_fips:parcel_id.
	CanonicalID *string `json:"canonical_id,omitempty"`

	// PropRaven row UUID when the serving row carried one.
	ID        *string `json:"id,omitempty"`
	StateFIPS *string `json:"state_fips,omitempty"`

	// 3-digit within-state county code (the parcels_serving convention).
	CountyFIPS *string `json:"county_fips,omitempty"`

	// Same as county_fips, named for its width.
	CountyFIPS3 *string `json:"county_fips_3,omitempty"`
	CountyFIPS5 *string `json:"county_fips_5,omitempty"`
	ParcelID    *string `json:"parcel_id,omitempty"`
	State       *string `json:"state,omitempty"`
	Address     *string `json:"address,omitempty"`
}

// ParcelDossierFields is generated from the OpenAPI spec.
type ParcelDossierFields struct {
	Name *string `json:"name,omitempty"`

	// The field value (any JSON type).
	Value json.RawMessage `json:"value,omitempty"`

	// Field-level source label from the sealed catalog.
	Source *string `json:"source,omitempty"`
	AsOf   *string `json:"as_of,omitempty"`

	// 0..1 FIELD-LEVEL coverage (this parcel's state, else national); null when uncatalogued. Not a
	// per-cell probability -- see confidence_basis.
	Confidence       *float64 `json:"confidence,omitempty"`
	ConfidenceBasis  *string  `json:"confidence_basis,omitempty"`
	NationalCoverage *float64 `json:"national_coverage,omitempty"`
	StateCoverage    *float64 `json:"state_coverage,omitempty"`

	// prime|strong|good|partial|sparse|trace, re-derived from the applicable (state, else national)
	// coverage.
	Tier    *string  `json:"tier,omitempty"`
	Grain   *string  `json:"grain,omitempty"`
	Section *string  `json:"section,omitempty"`
	Flags   []string `json:"flags,omitempty"`

	// False when the column is absent from the sealed catalog (source unknown).
	Catalogued *bool `json:"catalogued,omitempty"`
}

// UnmarshalJSON decodes ParcelDossierFields, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelDossierFields) UnmarshalJSON(data []byte) error {
	type plain ParcelDossierFields
	aux := struct {
		*plain
		Confidence       lenientNumber[float64] `json:"confidence"`
		NationalCoverage lenientNumber[float64] `json:"national_coverage"`
		StateCoverage    lenientNumber[float64] `json:"state_coverage"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Confidence.assignPtr(&r.Confidence)
	aux.NationalCoverage.assignPtr(&r.NationalCoverage)
	aux.StateCoverage.assignPtr(&r.StateCoverage)
	return softTypeError(err)
}

// ParcelDossierSections: Real sub-tables from the serving relations.
type ParcelDossierSections struct {
	Deeds   *DossierSubSection `json:"deeds,omitempty"`
	Comps   *DossierSubSection `json:"comps,omitempty"`
	Permits *DossierSubSection `json:"permits,omitempty"`
}

// ParcelDossierBoundary: The GeoJSON boundary add-on. In the base dossier included=false and
// geometry is absent; a phase-2 add-on purchase populates geometry.
type ParcelDossierBoundary struct {
	Code        *string `json:"code,omitempty"`
	Price       *Money  `json:"price,omitempty"`
	Included    *bool   `json:"included,omitempty"`
	Purchasable *bool   `json:"purchasable,omitempty"`
	Available   *bool   `json:"available,omitempty"`
	Relation    *string `json:"relation,omitempty"`
	Note        *string `json:"note,omitempty"`

	// GeoJSON geometry, present only for an entitled add-on purchase.
	Geometry json.RawMessage `json:"geometry,omitempty"`
}

// ParcelDossierMeta: Dossier-level provenance and pricing.
type ParcelDossierMeta struct {
	CatalogVersion  *string  `json:"catalog_version,omitempty"`
	CatalogSeal     *string  `json:"catalog_seal,omitempty"`
	ContractVersion *float64 `json:"contract_version,omitempty"`
	SnapshotAsOf    *string  `json:"snapshot_as_of,omitempty"`
	GeneratedAt     *string  `json:"generated_at,omitempty"`

	// The dossier product's contract price. The amount actually charged is the value-tiered per-parcel
	// quote settled via x402 (or billed on a paid subscription) -- see the 402 accepts and
	// /storefront/availability?parcel_id=.
	Price    *Money  `json:"price,omitempty"`
	Currency *string `json:"currency,omitempty"`

	// Populated fields actually emitted.
	FieldCount            *float64 `json:"field_count,omitempty"`
	CatalogSellableFields *float64 `json:"catalog_sellable_fields,omitempty"`
	CatalogTotalColumns   *float64 `json:"catalog_total_columns,omitempty"`
	SectionsIncluded      []string `json:"sections_included,omitempty"`
	GeometryIncluded      *bool    `json:"geometry_included,omitempty"`
	ConfidenceBasis       *string  `json:"confidence_basis,omitempty"`
	ProvenanceNote        *string  `json:"provenance_note,omitempty"`

	// The data license that governs this dossier's data (the Data license section of the Terms of
	// Use).
	License     *string                       `json:"license,omitempty"`
	Entitlement *ParcelDossierMetaEntitlement `json:"entitlement,omitempty"`
}

// UnmarshalJSON decodes ParcelDossierMeta, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelDossierMeta) UnmarshalJSON(data []byte) error {
	type plain ParcelDossierMeta
	aux := struct {
		*plain
		ContractVersion       lenientNumber[float64] `json:"contract_version"`
		FieldCount            lenientNumber[float64] `json:"field_count"`
		CatalogSellableFields lenientNumber[float64] `json:"catalog_sellable_fields"`
		CatalogTotalColumns   lenientNumber[float64] `json:"catalog_total_columns"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ContractVersion.assignPtr(&r.ContractVersion)
	aux.FieldCount.assignPtr(&r.FieldCount)
	aux.CatalogSellableFields.assignPtr(&r.CatalogSellableFields)
	aux.CatalogTotalColumns.assignPtr(&r.CatalogTotalColumns)
	return softTypeError(err)
}

// ParcelDossierMetaEntitlement is generated from the OpenAPI spec.
type ParcelDossierMetaEntitlement struct {
	Gated  *bool   `json:"gated,omitempty"`
	Method *string `json:"method,omitempty"`
	Note   *string `json:"note,omitempty"`
}

// PeopleFieldsWithheld: Additive marker on a record (or list) whose people fields were withheld
// because the caller has no PropRaven account. People fields are owner names, owner mailing /
// owner-address columns, entity principals, recorded-document party names and addresses
// (grantor/grantee, buyer/seller, prior/new owner), permit applicant names and the resolved
// `owner_contact` block. Withheld keys are kept and set to null (lists of people records become
// []), so the record's shape does not change. Payment alone (an x402 `X-PAYMENT` header or a
// prepaid `X-CREDIT-TOKEN`) is not an account: send an API key (`Authorization: Bearer pz_...`) or
// call from a signed-in session.
type PeopleFieldsWithheld struct {
	Status PeopleFieldsWithheldStatus `json:"status"`
	Code   PeopleFieldsWithheldCode   `json:"code"`
	Reason PeopleFieldsWithheldReason `json:"reason"`
	Note   string                     `json:"note"`
}

// PeopleFieldsWithheldStatus is generated from the OpenAPI spec. It is a string; the
// PeopleFieldsWithheldStatus* constants list the documented values.
type PeopleFieldsWithheldStatus = string

// Documented values of PeopleFieldsWithheldStatus.
const (
	PeopleFieldsWithheldStatusWithheld PeopleFieldsWithheldStatus = "withheld"
)

// PeopleFieldsWithheldCode is generated from the OpenAPI spec. It is a string; the
// PeopleFieldsWithheldCode* constants list the documented values.
type PeopleFieldsWithheldCode = string

// Documented values of PeopleFieldsWithheldCode.
const (
	PeopleFieldsWithheldCodeAccountRequired PeopleFieldsWithheldCode = "account_required"
)

// PeopleFieldsWithheldReason is generated from the OpenAPI spec. It is a string; the
// PeopleFieldsWithheldReason* constants list the documented values.
type PeopleFieldsWithheldReason = string

// Documented values of PeopleFieldsWithheldReason.
const (
	PeopleFieldsWithheldReasonPeopleDataRequiresAccount PeopleFieldsWithheldReason = "people_data_requires_account"
)

// UCCLien is generated from the OpenAPI spec.
type UCCLien struct {
	FilingID        *string  `json:"filing_id,omitempty"`
	DebtorName      *string  `json:"debtor_name,omitempty"`
	SecuredParty    *string  `json:"secured_party,omitempty"`
	FilingDate      *string  `json:"filing_date,omitempty"`
	LapseDate       *string  `json:"lapse_date,omitempty"`
	FilingStatus    *string  `json:"filing_status,omitempty"`
	MatchMethod     *string  `json:"match_method,omitempty"`
	MatchConfidence *float64 `json:"match_confidence,omitempty"`
}

// UnmarshalJSON decodes UCCLien, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *UCCLien) UnmarshalJSON(data []byte) error {
	type plain UCCLien
	aux := struct {
		*plain
		MatchConfidence lenientNumber[float64] `json:"match_confidence"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.MatchConfidence.assignPtr(&r.MatchConfidence)
	return softTypeError(err)
}

// ComparableSale is generated from the OpenAPI spec.
type ComparableSale struct {
	CompParcelID    *string  `json:"comp_parcel_id,omitempty"`
	CompSalePrice   *float64 `json:"comp_sale_price,omitempty"`
	CompSaleDate    *string  `json:"comp_sale_date,omitempty"`
	CompSqft        *int64   `json:"comp_sqft,omitempty"`
	CompBeds        *int64   `json:"comp_beds,omitempty"`
	CompBaths       *float64 `json:"comp_baths,omitempty"`
	SimilarityScore *float64 `json:"similarity_score,omitempty"`
	DistanceMiles   *float64 `json:"distance_miles,omitempty"`
}

// UnmarshalJSON decodes ComparableSale, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ComparableSale) UnmarshalJSON(data []byte) error {
	type plain ComparableSale
	aux := struct {
		*plain
		CompSalePrice   lenientNumber[float64] `json:"comp_sale_price"`
		CompSqft        lenientNumber[int64]   `json:"comp_sqft"`
		CompBeds        lenientNumber[int64]   `json:"comp_beds"`
		CompBaths       lenientNumber[float64] `json:"comp_baths"`
		SimilarityScore lenientNumber[float64] `json:"similarity_score"`
		DistanceMiles   lenientNumber[float64] `json:"distance_miles"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.CompSalePrice.assignPtr(&r.CompSalePrice)
	aux.CompSqft.assignPtr(&r.CompSqft)
	aux.CompBeds.assignPtr(&r.CompBeds)
	aux.CompBaths.assignPtr(&r.CompBaths)
	aux.SimilarityScore.assignPtr(&r.SimilarityScore)
	aux.DistanceMiles.assignPtr(&r.DistanceMiles)
	return softTypeError(err)
}

// TrafficStationHistory is generated from the OpenAPI spec.
type TrafficStationHistory struct {
	DirectVpd          *float64                          `json:"direct_vpd,omitempty"`
	NearbyVpd          *float64                          `json:"nearby_vpd,omitempty"`
	VpdVisibilityScore *float64                          `json:"vpd_visibility_score,omitempty"`
	WithholdGate       TrafficStationHistoryWithholdGate `json:"withhold_gate"`
	Points             []TrafficStationHistoryPoints     `json:"points"`
	Station            TrafficStationHistoryStation      `json:"station"`
	TimeSeries         []TrafficStationHistoryTimeSeries `json:"time_series"`

	// Year-keyed historical counts (newest last).
	History []TrafficStationHistoryHistory `json:"history,omitempty"`
}

// UnmarshalJSON decodes TrafficStationHistory, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *TrafficStationHistory) UnmarshalJSON(data []byte) error {
	type plain TrafficStationHistory
	aux := struct {
		*plain
		DirectVpd          lenientNumber[float64] `json:"direct_vpd"`
		NearbyVpd          lenientNumber[float64] `json:"nearby_vpd"`
		VpdVisibilityScore lenientNumber[float64] `json:"vpd_visibility_score"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.DirectVpd.assignPtr(&r.DirectVpd)
	aux.NearbyVpd.assignPtr(&r.NearbyVpd)
	aux.VpdVisibilityScore.assignPtr(&r.VpdVisibilityScore)
	return softTypeError(err)
}

// TrafficStationHistoryWithholdGate is generated from the OpenAPI spec.
type TrafficStationHistoryWithholdGate struct {
	Applied    *bool                                    `json:"applied,omitempty"`
	Suppressed []*string                                `json:"suppressed"`
	Reason     *string                                  `json:"reason,omitempty"`
	Reasons    TrafficStationHistoryWithholdGateReasons `json:"reasons"`
	WithheldOn *string                                  `json:"withheld_on,omitempty"`
	Note       *string                                  `json:"note,omitempty"`
}

// TrafficStationHistoryWithholdGateReasons is generated from the OpenAPI spec.
type TrafficStationHistoryWithholdGateReasons struct {
	DirectVpd          *string `json:"direct_vpd,omitempty"`
	NearbyVpd          *string `json:"nearby_vpd,omitempty"`
	VpdVisibilityScore *string `json:"vpd_visibility_score,omitempty"`
}

// TrafficStationHistoryPoints is generated from the OpenAPI spec.
type TrafficStationHistoryPoints struct {
	Date *string  `json:"date,omitempty"`
	Vpd  *float64 `json:"vpd,omitempty"`
}

// UnmarshalJSON decodes TrafficStationHistoryPoints, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *TrafficStationHistoryPoints) UnmarshalJSON(data []byte) error {
	type plain TrafficStationHistoryPoints
	aux := struct {
		*plain
		Vpd lenientNumber[float64] `json:"vpd"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Vpd.assignPtr(&r.Vpd)
	return softTypeError(err)
}

// TrafficStationHistoryStation is generated from the OpenAPI spec.
type TrafficStationHistoryStation struct {
	ID              string  `json:"id"`
	StationID       *string `json:"station_id,omitempty"`
	RouteName       *string `json:"route_name,omitempty"`
	FunctionalClass *string `json:"functional_class,omitempty"`

	// Most recent AADT count.
	AadtCurrent *float64 `json:"aadt_current,omitempty"`
	AadtYear    *int64   `json:"aadt_year,omitempty"`
	Cagr3yr     *float64 `json:"cagr_3yr,omitempty"`
	Cagr5yr     *float64 `json:"cagr_5yr,omitempty"`
	Cagr7yr     *float64 `json:"cagr_7yr,omitempty"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
}

// UnmarshalJSON decodes TrafficStationHistoryStation, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *TrafficStationHistoryStation) UnmarshalJSON(data []byte) error {
	type plain TrafficStationHistoryStation
	aux := struct {
		*plain
		AadtCurrent lenientNumber[float64] `json:"aadt_current"`
		AadtYear    lenientNumber[int64]   `json:"aadt_year"`
		Cagr3yr     lenientNumber[float64] `json:"cagr_3yr"`
		Cagr5yr     lenientNumber[float64] `json:"cagr_5yr"`
		Cagr7yr     lenientNumber[float64] `json:"cagr_7yr"`
		Latitude    lenientNumber[float64] `json:"latitude"`
		Longitude   lenientNumber[float64] `json:"longitude"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.AadtCurrent.assignPtr(&r.AadtCurrent)
	aux.AadtYear.assignPtr(&r.AadtYear)
	aux.Cagr3yr.assignPtr(&r.Cagr3yr)
	aux.Cagr5yr.assignPtr(&r.Cagr5yr)
	aux.Cagr7yr.assignPtr(&r.Cagr7yr)
	aux.Latitude.assignPtr(&r.Latitude)
	aux.Longitude.assignPtr(&r.Longitude)
	return softTypeError(err)
}

// TrafficStationHistoryTimeSeries is generated from the OpenAPI spec.
type TrafficStationHistoryTimeSeries struct {
	Year *int64   `json:"year,omitempty"`
	Aadt *float64 `json:"aadt,omitempty"`
}

// UnmarshalJSON decodes TrafficStationHistoryTimeSeries, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *TrafficStationHistoryTimeSeries) UnmarshalJSON(data []byte) error {
	type plain TrafficStationHistoryTimeSeries
	aux := struct {
		*plain
		Year lenientNumber[int64]   `json:"year"`
		Aadt lenientNumber[float64] `json:"aadt"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Year.assignPtr(&r.Year)
	aux.Aadt.assignPtr(&r.Aadt)
	return softTypeError(err)
}

// TrafficStationHistoryHistory is generated from the OpenAPI spec.
type TrafficStationHistoryHistory struct {
	Year *int64 `json:"year,omitempty"`
	Aadt *int64 `json:"aadt,omitempty"`
}

// UnmarshalJSON decodes TrafficStationHistoryHistory, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *TrafficStationHistoryHistory) UnmarshalJSON(data []byte) error {
	type plain TrafficStationHistoryHistory
	aux := struct {
		*plain
		Year lenientNumber[int64] `json:"year"`
		Aadt lenientNumber[int64] `json:"aadt"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Year.assignPtr(&r.Year)
	aux.Aadt.assignPtr(&r.Aadt)
	return softTypeError(err)
}

// CountyDetail is generated from the OpenAPI spec.
type CountyDetail struct {
	CountyFIPS     string                     `json:"county_fips"`
	MarketStats    []CountyDetailMarketStats  `json:"market_stats"`
	Affordability  []AffordabilityRow         `json:"affordability"`
	DataProvenance CountyDetailDataProvenance `json:"data_provenance"`
	ParcelSummary  CountyDetailParcelSummary  `json:"parcel_summary"`
	FlipSummary    CountyDetailFlipSummary    `json:"flip_summary"`
}

// CountyDetailMarketStats is generated from the OpenAPI spec.
type CountyDetailMarketStats struct {
	CountyFIPS      string   `json:"county_fips"`
	StateFIPS       string   `json:"state_fips"`
	Quarter         *string  `json:"quarter,omitempty"`
	SaleCount       *int64   `json:"sale_count,omitempty"`
	MedianSalePrice *float64 `json:"median_sale_price,omitempty"`
	AvgSalePrice    *float64 `json:"avg_sale_price,omitempty"`
	TotalVolume     *int64   `json:"total_volume,omitempty"`
	PriceYoyPct     *float64 `json:"price_yoy_pct,omitempty"`

	// Average days on market.
	AvgDom       *float64  `json:"avg_dom,omitempty"`
	RefreshedAt  *string   `json:"refreshed_at,omitempty"`
	UnderReview  []*string `json:"under_review,omitempty"`
	StaleQuarter *bool     `json:"stale_quarter,omitempty"`
}

// UnmarshalJSON decodes CountyDetailMarketStats, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CountyDetailMarketStats) UnmarshalJSON(data []byte) error {
	type plain CountyDetailMarketStats
	aux := struct {
		*plain
		SaleCount       lenientNumber[int64]   `json:"sale_count"`
		MedianSalePrice lenientNumber[float64] `json:"median_sale_price"`
		AvgSalePrice    lenientNumber[float64] `json:"avg_sale_price"`
		TotalVolume     lenientNumber[int64]   `json:"total_volume"`
		PriceYoyPct     lenientNumber[float64] `json:"price_yoy_pct"`
		AvgDom          lenientNumber[float64] `json:"avg_dom"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.SaleCount.assignPtr(&r.SaleCount)
	aux.MedianSalePrice.assignPtr(&r.MedianSalePrice)
	aux.AvgSalePrice.assignPtr(&r.AvgSalePrice)
	aux.TotalVolume.assignPtr(&r.TotalVolume)
	aux.PriceYoyPct.assignPtr(&r.PriceYoyPct)
	aux.AvgDom.assignPtr(&r.AvgDom)
	return softTypeError(err)
}

// CountyDetailDataProvenance is generated from the OpenAPI spec.
type CountyDetailDataProvenance struct {
	Market        CountyDetailDataProvenanceMarket        `json:"market"`
	Affordability CountyDetailDataProvenanceAffordability `json:"affordability"`
}

// CountyDetailDataProvenanceMarket is generated from the OpenAPI spec.
type CountyDetailDataProvenanceMarket struct {
	Dataset               *string  `json:"dataset,omitempty"`
	PeriodUpperBound      *string  `json:"period_upper_bound,omitempty"`
	ReturnedRecords       *float64 `json:"returned_records,omitempty"`
	OldestRecordRefresh   *string  `json:"oldest_record_refresh,omitempty"`
	NewestRecordRefresh   *string  `json:"newest_record_refresh,omitempty"`
	RecordsWithoutRefresh *float64 `json:"records_without_refresh,omitempty"`
	SourceVintage         *string  `json:"source_vintage,omitempty"`
	Basis                 *string  `json:"basis,omitempty"`
}

// UnmarshalJSON decodes CountyDetailDataProvenanceMarket, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CountyDetailDataProvenanceMarket) UnmarshalJSON(data []byte) error {
	type plain CountyDetailDataProvenanceMarket
	aux := struct {
		*plain
		ReturnedRecords       lenientNumber[float64] `json:"returned_records"`
		RecordsWithoutRefresh lenientNumber[float64] `json:"records_without_refresh"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ReturnedRecords.assignPtr(&r.ReturnedRecords)
	aux.RecordsWithoutRefresh.assignPtr(&r.RecordsWithoutRefresh)
	return softTypeError(err)
}

// CountyDetailDataProvenanceAffordability is generated from the OpenAPI spec.
type CountyDetailDataProvenanceAffordability struct {
	Dataset               *string                                              `json:"dataset,omitempty"`
	PeriodUpperBound      *float64                                             `json:"period_upper_bound,omitempty"`
	ReturnedRecords       *float64                                             `json:"returned_records,omitempty"`
	OldestRecordRefresh   *string                                              `json:"oldest_record_refresh,omitempty"`
	NewestRecordRefresh   *string                                              `json:"newest_record_refresh,omitempty"`
	RecordsWithoutRefresh *float64                                             `json:"records_without_refresh,omitempty"`
	SourceVintage         *string                                              `json:"source_vintage,omitempty"`
	IncomeMeasure         CountyDetailDataProvenanceAffordabilityIncomeMeasure `json:"income_measure"`
	Basis                 *string                                              `json:"basis,omitempty"`
}

// UnmarshalJSON decodes CountyDetailDataProvenanceAffordability, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CountyDetailDataProvenanceAffordability) UnmarshalJSON(data []byte) error {
	type plain CountyDetailDataProvenanceAffordability
	aux := struct {
		*plain
		PeriodUpperBound      lenientNumber[float64] `json:"period_upper_bound"`
		ReturnedRecords       lenientNumber[float64] `json:"returned_records"`
		RecordsWithoutRefresh lenientNumber[float64] `json:"records_without_refresh"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PeriodUpperBound.assignPtr(&r.PeriodUpperBound)
	aux.ReturnedRecords.assignPtr(&r.ReturnedRecords)
	aux.RecordsWithoutRefresh.assignPtr(&r.RecordsWithoutRefresh)
	return softTypeError(err)
}

// CountyDetailDataProvenanceAffordabilityIncomeMeasure is generated from the OpenAPI spec.
type CountyDetailDataProvenanceAffordabilityIncomeMeasure struct {
	StoredField           *string `json:"stored_field,omitempty"`
	Interpretation        *string `json:"interpretation,omitempty"`
	SourceTaxYear         *string `json:"source_tax_year,omitempty"`
	SourceLineageVerified *bool   `json:"source_lineage_verified,omitempty"`
	Limitations           *string `json:"limitations,omitempty"`
}

// CountyDetailParcelSummary is generated from the OpenAPI spec.
type CountyDetailParcelSummary struct {
	ParcelCount      *int64   `json:"parcel_count,omitempty"`
	AvgAssessedValue *float64 `json:"avg_assessed_value,omitempty"`
}

// UnmarshalJSON decodes CountyDetailParcelSummary, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CountyDetailParcelSummary) UnmarshalJSON(data []byte) error {
	type plain CountyDetailParcelSummary
	aux := struct {
		*plain
		ParcelCount      lenientNumber[int64]   `json:"parcel_count"`
		AvgAssessedValue lenientNumber[float64] `json:"avg_assessed_value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ParcelCount.assignPtr(&r.ParcelCount)
	aux.AvgAssessedValue.assignPtr(&r.AvgAssessedValue)
	return softTypeError(err)
}

// CountyDetailFlipSummary is generated from the OpenAPI spec.
type CountyDetailFlipSummary struct {
	FlipCount   *int64   `json:"flip_count,omitempty"`
	AvgRoi      *float64 `json:"avg_roi,omitempty"`
	AvgHoldDays *float64 `json:"avg_hold_days,omitempty"`
	TotalProfit *float64 `json:"total_profit,omitempty"`
}

// UnmarshalJSON decodes CountyDetailFlipSummary, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *CountyDetailFlipSummary) UnmarshalJSON(data []byte) error {
	type plain CountyDetailFlipSummary
	aux := struct {
		*plain
		FlipCount   lenientNumber[int64]   `json:"flip_count"`
		AvgRoi      lenientNumber[float64] `json:"avg_roi"`
		AvgHoldDays lenientNumber[float64] `json:"avg_hold_days"`
		TotalProfit lenientNumber[float64] `json:"total_profit"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.FlipCount.assignPtr(&r.FlipCount)
	aux.AvgRoi.assignPtr(&r.AvgRoi)
	aux.AvgHoldDays.assignPtr(&r.AvgHoldDays)
	aux.TotalProfit.assignPtr(&r.TotalProfit)
	return softTypeError(err)
}

// MarketFlipsRow is generated from the OpenAPI spec.
type MarketFlipsRow struct {
	CountyFIPS string `json:"county_fips"`
	FlipCount  *int64 `json:"flip_count,omitempty"`

	// Average profit percentage (e.g. 0.18 = 18%).
	AvgRoi      *float64 `json:"avg_roi,omitempty"`
	AvgHoldDays *float64 `json:"avg_hold_days,omitempty"`
	TotalProfit *int64   `json:"total_profit,omitempty"`
}

// UnmarshalJSON decodes MarketFlipsRow, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MarketFlipsRow) UnmarshalJSON(data []byte) error {
	type plain MarketFlipsRow
	aux := struct {
		*plain
		FlipCount   lenientNumber[int64]   `json:"flip_count"`
		AvgRoi      lenientNumber[float64] `json:"avg_roi"`
		AvgHoldDays lenientNumber[float64] `json:"avg_hold_days"`
		TotalProfit lenientNumber[int64]   `json:"total_profit"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.FlipCount.assignPtr(&r.FlipCount)
	aux.AvgRoi.assignPtr(&r.AvgRoi)
	aux.AvgHoldDays.assignPtr(&r.AvgHoldDays)
	aux.TotalProfit.assignPtr(&r.TotalProfit)
	return softTypeError(err)
}

// OwnerTransaction is generated from the OpenAPI spec.
type OwnerTransaction struct {
	DocumentNumber *string `json:"document_number,omitempty"`
	RecordingDate  *string `json:"recording_date,omitempty"`
	SaleDate       *string `json:"sale_date,omitempty"`

	// Recorded document type (Warranty Deed, Quit Claim, etc.).
	DocumentType *string `json:"document_type,omitempty"`

	// USD. Null when state is non-disclosure.
	SalePrice       *float64 `json:"sale_price,omitempty"`
	GrantorName     *string  `json:"grantor_name,omitempty"`
	GranteeName     *string  `json:"grantee_name,omitempty"`
	PropertyAddress *string  `json:"property_address,omitempty"`
}

// UnmarshalJSON decodes OwnerTransaction, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnerTransaction) UnmarshalJSON(data []byte) error {
	type plain OwnerTransaction
	aux := struct {
		*plain
		SalePrice lenientNumber[float64] `json:"sale_price"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.SalePrice.assignPtr(&r.SalePrice)
	return softTypeError(err)
}

// AccountUsage is generated from the OpenAPI spec.
type AccountUsage struct {
	Tier        AccountUsageTier `json:"tier"`
	Period      string           `json:"period"`
	PeriodStart string           `json:"period_start"`
	PeriodEnd   string           `json:"period_end"`
	CallsUsed   int64            `json:"calls_used"`

	// Plan allotment for the current period.
	CallsIncluded  int64                 `json:"calls_included"`
	CallsRemaining int64                 `json:"calls_remaining"`
	RateLimit      AccountUsageRateLimit `json:"rate_limit"`
	HardCapEnabled bool                  `json:"hard_cap_enabled"`
	AuthSource     string                `json:"auth_source"`

	// Whether further calls will be hard-rejected vs allowed-and-billed.
	HardCapped *bool `json:"hard_capped,omitempty"`
}

// UnmarshalJSON decodes AccountUsage, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *AccountUsage) UnmarshalJSON(data []byte) error {
	type plain AccountUsage
	aux := struct {
		*plain
		CallsUsed      lenientNumber[int64] `json:"calls_used"`
		CallsIncluded  lenientNumber[int64] `json:"calls_included"`
		CallsRemaining lenientNumber[int64] `json:"calls_remaining"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.CallsUsed.assign(&r.CallsUsed)
	aux.CallsIncluded.assign(&r.CallsIncluded)
	aux.CallsRemaining.assign(&r.CallsRemaining)
	return softTypeError(err)
}

// AccountUsageTier is generated from the OpenAPI spec. It is a string; the AccountUsageTier*
// constants list the documented values.
type AccountUsageTier = string

// Documented values of AccountUsageTier.
const (
	AccountUsageTierFree    AccountUsageTier = "free"
	AccountUsageTierStarter AccountUsageTier = "starter"
	AccountUsageTierPro     AccountUsageTier = "pro"
	AccountUsageTierScale   AccountUsageTier = "scale"
	AccountUsageTierAPI100k AccountUsageTier = "api_100k"
)

// AccountUsageRateLimit is generated from the OpenAPI spec.
type AccountUsageRateLimit struct {
	PerMinute float64 `json:"per_minute"`
	PerDay    float64 `json:"per_day"`
}

// UnmarshalJSON decodes AccountUsageRateLimit, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *AccountUsageRateLimit) UnmarshalJSON(data []byte) error {
	type plain AccountUsageRateLimit
	aux := struct {
		*plain
		PerMinute lenientNumber[float64] `json:"per_minute"`
		PerDay    lenientNumber[float64] `json:"per_day"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PerMinute.assign(&r.PerMinute)
	aux.PerDay.assign(&r.PerDay)
	return softTypeError(err)
}

// ParcelOwnerChangedEvent is generated from the OpenAPI spec.
type ParcelOwnerChangedEvent struct {
	EventType string `json:"event_type"`

	// Deterministic. Use for idempotency.
	EventID    string `json:"event_id"`
	OccurredAt string `json:"occurred_at"`

	// Starts at 1; increments on retry.
	DeliveryAttempt int64 `json:"delivery_attempt"`

	// Composite county_fips:parcel_id.
	ParcelID     string  `json:"parcel_id"`
	StateFIPS    string  `json:"state_fips"`
	CountyFIPS   string  `json:"county_fips"`
	RecordedDate *string `json:"recorded_date,omitempty"`

	// Deed/document classification (e.g., Warranty Deed, Quitclaim Deed, Trust Transfer).
	DocumentType   *string `json:"document_type,omitempty"`
	DocumentNumber *string `json:"document_number,omitempty"`

	// Grantor on the recorded deed.
	PriorOwner *string `json:"prior_owner,omitempty"`

	// Grantee on the recorded deed.
	NewOwner *string `json:"new_owner,omitempty"`

	// True when the transfer is a real sale (price > 0, arm's-length). When true, subscribers to
	// parcel.sold ALSO receive that event.
	IsSale *bool `json:"is_sale,omitempty"`

	// PropRaven ingest run that surfaced this event.
	SourceRunID *string `json:"source_run_id,omitempty"`
}

// UnmarshalJSON decodes ParcelOwnerChangedEvent, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelOwnerChangedEvent) UnmarshalJSON(data []byte) error {
	type plain ParcelOwnerChangedEvent
	aux := struct {
		*plain
		DeliveryAttempt lenientNumber[int64] `json:"delivery_attempt"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.DeliveryAttempt.assign(&r.DeliveryAttempt)
	return softTypeError(err)
}

// ParcelPermitFiledEvent is generated from the OpenAPI spec.
type ParcelPermitFiledEvent struct {
	EventType string `json:"event_type"`

	// Deterministic. Use for idempotency.
	EventID    string `json:"event_id"`
	OccurredAt string `json:"occurred_at"`

	// Starts at 1; increments on retry.
	DeliveryAttempt int64 `json:"delivery_attempt"`

	// Composite county_fips:parcel_id.
	ParcelID   string `json:"parcel_id"`
	StateFIPS  string `json:"state_fips"`
	CountyFIPS string `json:"county_fips"`

	// PropRaven internal permit id (stable across re-ingests).
	PermitID string `json:"permit_id"`

	// Jurisdiction-issued permit number.
	PermitNumber *string `json:"permit_number,omitempty"`

	// Building, electrical, roofing, demolition, etc.
	PermitType *string `json:"permit_type,omitempty"`

	// Filed, issued, in_review, final, expired, withdrawn.
	PermitStatus *string `json:"permit_status,omitempty"`
	FiledDate    *string `json:"filed_date,omitempty"`
	IssuedDate   *string `json:"issued_date,omitempty"`
	Description  *string `json:"description,omitempty"`

	// Declared job cost in USD.
	EstimatedCost     *float64 `json:"estimated_cost,omitempty"`
	ContractorName    *string  `json:"contractor_name,omitempty"`
	ContractorLicense *string  `json:"contractor_license,omitempty"`
	ApplicantName     *string  `json:"applicant_name,omitempty"`

	// PropRaven jurisdiction id; join to /v1/jurisdictions.
	JurisdictionID *string `json:"jurisdiction_id,omitempty"`

	// PropRaven ingest run that surfaced this event.
	SourceRunID *string `json:"source_run_id,omitempty"`
}

// UnmarshalJSON decodes ParcelPermitFiledEvent, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ParcelPermitFiledEvent) UnmarshalJSON(data []byte) error {
	type plain ParcelPermitFiledEvent
	aux := struct {
		*plain
		DeliveryAttempt lenientNumber[int64]   `json:"delivery_attempt"`
		EstimatedCost   lenientNumber[float64] `json:"estimated_cost"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.DeliveryAttempt.assign(&r.DeliveryAttempt)
	aux.EstimatedCost.assignPtr(&r.EstimatedCost)
	return softTypeError(err)
}

// X402PaymentRequired: x402 (HTTP 402) payment-required body (x402 protocol v1). `accepts` lists
// the payment requirements a wallet-bearing agent signs to pay per call. Returned by both paid
// products: the per-parcel dossier (GET /api/v1/parcels/{id}/report) and the per-lead feed (GET
// /api/v1/leads/find). Flow: sign an EIP-3009 transferWithAuthorization for
// accepts[0].maxAmountRequired, base64 the PaymentPayload into the `X-PAYMENT` request header, and
// retry. The response also carries `Link: <https://propraven.com/terms#data-license>;
// rel="license"`, the data license that governs what the payment buys.
type X402PaymentRequired struct {
	// x402 protocol version advertised (1 = the base exact scheme with maxAmountRequired).
	X402Version *int64 `json:"x402Version,omitempty"`

	// Human-readable reason payment is required or was rejected.
	Error *string `json:"error,omitempty"`

	// The payment requirements to satisfy (one entry: the dossier, or the lead pull).
	Accepts []X402PaymentRequiredAccepts `json:"accepts,omitempty"`
}

// UnmarshalJSON decodes X402PaymentRequired, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *X402PaymentRequired) UnmarshalJSON(data []byte) error {
	type plain X402PaymentRequired
	aux := struct {
		*plain
		X402Version lenientNumber[int64] `json:"x402Version"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.X402Version.assignPtr(&r.X402Version)
	return softTypeError(err)
}

// X402PaymentRequiredAccepts is generated from the OpenAPI spec.
type X402PaymentRequiredAccepts struct {
	Scheme *string `json:"scheme,omitempty"`

	// base-sepolia (testnet default) or base (mainnet).
	Network *string `json:"network,omitempty"`

	// The DYNAMIC price in the asset's atomic units (USDC, 6 decimals). Dossier: clamp($5 x V x R x F,
	// $2, $20) per parcel. Lead feed: min(count x clamp($0.25 x S x V, $0.05, $1.00), $20) per pull.
	// 2000000 = $2.00, 6250000 = $6.25, 20000000 = the $20 cap. Sign for exactly this amount.
	MaxAmountRequired *string `json:"maxAmountRequired,omitempty"`

	// The absolute URL being paid for.
	Resource    *string `json:"resource,omitempty"`
	Description *string `json:"description,omitempty"`
	MimeType    *string `json:"mimeType,omitempty"`

	// Receiving wallet address.
	PayTo *string `json:"payTo,omitempty"`

	// How long the quote is valid before the client must re-fetch it.
	MaxTimeoutSeconds *int64 `json:"maxTimeoutSeconds,omitempty"`

	// ERC-20 asset contract address (USDC on the given network).
	Asset *string `json:"asset,omitempty"`

	// The asset's EIP-712 domain { name, version } (USDC = { name: 'USDC', version: '2' }) -- required
	// by the exact EVM scheme so the facilitator can reconstruct the domain separator and verify the
	// transferWithAuthorization signature.
	Extra *X402PaymentRequiredAcceptsExtra `json:"extra,omitempty"`
}

// UnmarshalJSON decodes X402PaymentRequiredAccepts, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *X402PaymentRequiredAccepts) UnmarshalJSON(data []byte) error {
	type plain X402PaymentRequiredAccepts
	aux := struct {
		*plain
		MaxTimeoutSeconds lenientNumber[int64] `json:"maxTimeoutSeconds"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.MaxTimeoutSeconds.assignPtr(&r.MaxTimeoutSeconds)
	return softTypeError(err)
}

// X402PaymentRequiredAcceptsExtra: The asset's EIP-712 domain { name, version } (USDC = { name:
// 'USDC', version: '2' }) -- required by the exact EVM scheme so the facilitator can reconstruct
// the domain separator and verify the transferWithAuthorization signature.
type X402PaymentRequiredAcceptsExtra struct {
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// LeadsQuote: The per-lead price quote. IDENTICAL in the free preview and in the 402/charge -- the
// previewed price is the paid price.
type LeadsQuote struct {
	PerLead Money `json:"per_lead"`

	// Unit price to 4dp: clamp($0.25 x S x V, $0.05, $1.00).
	PerLeadUsd float64 `json:"per_lead_usd"`
	Total      Money   `json:"total"`

	// The amount actually charged: min(count x per_lead, $20).
	TotalUsd float64 `json:"total_usd"`

	// total_usd in USDC atomic units (6dp) -- what the 402 advertises as maxAmountRequired.
	TotalAtomicUsdc string `json:"total_atomic_usdc"`
	Asset           string `json:"asset"`

	// Leads priced = leads delivered.
	Count  int64  `json:"count"`
	Signal string `json:"signal"`

	// Asset-value tier, from the median assessed value of the delivered set.
	Tier      LeadsQuoteTier `json:"tier"`
	TierLabel string         `json:"tier_label"`

	// The dials, so the price is auditable.
	Breakdown LeadsQuoteBreakdown `json:"breakdown"`
	Pay       []*string           `json:"pay"`
	Note      string              `json:"note"`
}

// UnmarshalJSON decodes LeadsQuote, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LeadsQuote) UnmarshalJSON(data []byte) error {
	type plain LeadsQuote
	aux := struct {
		*plain
		PerLeadUsd lenientNumber[float64] `json:"per_lead_usd"`
		TotalUsd   lenientNumber[float64] `json:"total_usd"`
		Count      lenientNumber[int64]   `json:"count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.PerLeadUsd.assign(&r.PerLeadUsd)
	aux.TotalUsd.assign(&r.TotalUsd)
	aux.Count.assign(&r.Count)
	return softTypeError(err)
}

// LeadsQuoteTier: Asset-value tier, from the median assessed value of the delivered set.
//
// It is a string; the LeadsQuoteTier* constants list the documented values.
type LeadsQuoteTier = string

// Documented values of LeadsQuoteTier.
const (
	LeadsQuoteTierLow     LeadsQuoteTier = "low"
	LeadsQuoteTierMid     LeadsQuoteTier = "mid"
	LeadsQuoteTierHigh    LeadsQuoteTier = "high"
	LeadsQuoteTierPremium LeadsQuoteTier = "premium"
)

// LeadsQuoteBreakdown: The dials, so the price is auditable.
type LeadsQuoteBreakdown struct {
	BasePerLead float64 `json:"base_per_lead"`

	// Signal-strength multiplier (absentee 1.0 ... distressed 1.9).
	S float64 `json:"S"`

	// Asset-value tier multiplier (low 0.7, mid 1.0, high 1.5, premium 2.2).
	V float64 `json:"V"`

	// Per-lead price BEFORE the [$0.05, $1.00] clamp.
	PerLeadRaw float64 `json:"per_lead_raw"`

	// count x per_lead BEFORE the $20 per-call cap.
	TotalRaw float64 `json:"total_raw"`

	// True when the $20 cap bound -- volume beyond it was free.
	Capped bool `json:"capped"`
}

// UnmarshalJSON decodes LeadsQuoteBreakdown, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LeadsQuoteBreakdown) UnmarshalJSON(data []byte) error {
	type plain LeadsQuoteBreakdown
	aux := struct {
		*plain
		BasePerLead lenientNumber[float64] `json:"base_per_lead"`
		S           lenientNumber[float64] `json:"S"`
		V           lenientNumber[float64] `json:"V"`
		PerLeadRaw  lenientNumber[float64] `json:"per_lead_raw"`
		TotalRaw    lenientNumber[float64] `json:"total_raw"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.BasePerLead.assign(&r.BasePerLead)
	aux.S.assign(&r.S)
	aux.V.assign(&r.V)
	aux.PerLeadRaw.assign(&r.PerLeadRaw)
	aux.TotalRaw.assign(&r.TotalRaw)
	return softTypeError(err)
}

// Lead: One delivered lead. Signal-specific strength fields are flattened alongside the common
// keys (e.g. years_held + hold_tier for long_hold; profit + profit_pct + flip_tier for flip;
// land_improvement_ratio for high_land_ratio/distressed; property_count + states_list for
// portfolio_owner).
type Lead struct {
	// state_fips:county_fips3:parcel_id (the assessor APN, never a PropRaven UUID). Null on the
	// owner-grain portfolio_owner cohort. MASKED to "37:183:..." in the free preview.
	CanonicalID string `json:"canonical_id"`

	// Situs address. House number stripped in the free preview.
	Address *string `json:"address,omitempty"`
	City    *string `json:"city,omitempty"`
	State   *string `json:"state,omitempty"`
	Zip     *string `json:"zip,omitempty"`

	// Total assessed value (sell price on the flip cohort).
	AssessedValue *float64 `json:"assessed_value,omitempty"`

	// Deterministic 1-100: the signal's base intent plus a bounded bonus from that signal's own
	// strength column. Re-derivable, never random.
	LeadScore     *float64 `json:"lead_score,omitempty"`
	OwnerAddress  *string  `json:"owner_address,omitempty"`
	OwnerCity     *string  `json:"owner_city,omitempty"`
	OwnerState    *string  `json:"owner_state,omitempty"`
	IsOutOfState  *bool    `json:"is_out_of_state,omitempty"`
	LastSaleDate  *string  `json:"last_sale_date,omitempty"`
	LastSalePrice *float64 `json:"last_sale_price,omitempty"`

	// Withheld (null) in the free preview. Delivered leads go to accounts only.
	OwnerName *string `json:"owner_name,omitempty"`

	// Present and true ONLY on free-preview sample leads.
	Masked     *bool          `json:"masked,omitempty"`
	Provenance LeadProvenance `json:"provenance"`

	// Delivered leads only: the owner's best mailing address (ONE column family, with its ZIP),
	// flagged mail_ready. Absent from the free preview.
	OwnerContact *LeadOwnerContact `json:"owner_contact,omitempty"`
}

// UnmarshalJSON decodes Lead, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *Lead) UnmarshalJSON(data []byte) error {
	type plain Lead
	aux := struct {
		*plain
		AssessedValue lenientNumber[float64] `json:"assessed_value"`
		LeadScore     lenientNumber[float64] `json:"lead_score"`
		LastSalePrice lenientNumber[float64] `json:"last_sale_price"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.AssessedValue.assignPtr(&r.AssessedValue)
	aux.LeadScore.assignPtr(&r.LeadScore)
	aux.LastSalePrice.assignPtr(&r.LastSalePrice)
	return softTypeError(err)
}

// LeadProvenance is generated from the OpenAPI spec.
type LeadProvenance struct {
	Source         *string   `json:"source,omitempty"`
	SourceDatasets []*string `json:"source_datasets"`

	// The serving epoch's generated_at.
	AsOf             *string                        `json:"as_of,omitempty"`
	AsOfBasis        *string                        `json:"as_of_basis,omitempty"`
	ServingEpoch     *string                        `json:"serving_epoch,omitempty"`
	FreshnessStatus  *string                        `json:"freshness_status,omitempty"`
	CatalogReference LeadProvenanceCatalogReference `json:"catalog_reference"`
	Note             *string                        `json:"note,omitempty"`
}

// LeadProvenanceCatalogReference is generated from the OpenAPI spec.
type LeadProvenanceCatalogReference struct {
	CatalogVersion     *string                                   `json:"catalog_version,omitempty"`
	CatalogGeneratedAt *string                                   `json:"catalog_generated_at,omitempty"`
	CatalogSeal        LeadProvenanceCatalogReferenceCatalogSeal `json:"catalog_seal"`
}

// LeadProvenanceCatalogReferenceCatalogSeal is generated from the OpenAPI spec.
type LeadProvenanceCatalogReferenceCatalogSeal struct {
	Algorithm *string `json:"algorithm,omitempty"`
	Value     *string `json:"value,omitempty"`
	Covers    *string `json:"covers,omitempty"`
}

// LeadOwnerContact: The owner's best mailing address on a delivered lead (account holders only).
// Lead pulls do not run the per-parcel permit phone lookup (`phone_status` is always not_checked);
// call GET /api/v1/owners/card for the full card.
type LeadOwnerContact struct {
	// unavailable = the contact read failed for this delivery (never a silently missing block).
	Status              *LeadOwnerContactStatus      `json:"status,omitempty"`
	Mailing             *MailingAddress              `json:"mailing,omitempty"`
	MailReady           *bool                        `json:"mail_ready,omitempty"`
	PhoneStatus         *LeadOwnerContactPhoneStatus `json:"phone_status,omitempty"`
	HiddenLowConfidence *int64                       `json:"hidden_low_confidence,omitempty"`
}

// UnmarshalJSON decodes LeadOwnerContact, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LeadOwnerContact) UnmarshalJSON(data []byte) error {
	type plain LeadOwnerContact
	aux := struct {
		*plain
		HiddenLowConfidence lenientNumber[int64] `json:"hidden_low_confidence"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.HiddenLowConfidence.assignPtr(&r.HiddenLowConfidence)
	return softTypeError(err)
}

// LeadOwnerContactStatus: unavailable = the contact read failed for this delivery (never a
// silently missing block).
//
// It is a string; the LeadOwnerContactStatus* constants list the documented values.
type LeadOwnerContactStatus = string

// Documented values of LeadOwnerContactStatus.
const (
	LeadOwnerContactStatusResolved    LeadOwnerContactStatus = "resolved"
	LeadOwnerContactStatusNotFound    LeadOwnerContactStatus = "not_found"
	LeadOwnerContactStatusUnavailable LeadOwnerContactStatus = "unavailable"
)

// MailingAddress: One mailing address, read from ONE address column family of the record
// (owner_mailing_* or owner_*), never a street from one family and a ZIP from the other.
type MailingAddress struct {
	// Where the address came from, e.g. owner_mailing, owner_address, deed_grantee.
	Basis string  `json:"basis"`
	Line1 string  `json:"line1"`
	City  string  `json:"city"`
	State string  `json:"state"`
	Zip5  string  `json:"zip5"`
	Zip4  *string `json:"zip4,omitempty"`

	// Street, city, state and a 5-digit ZIP are all present and consistent.
	MailReady bool `json:"mail_ready"`
	PoBox     bool `json:"po_box"`

	// The mailing address is the property itself (owner-occupied).
	EqualsSitus bool `json:"equals_situs"`
	ZipConflict bool `json:"zip_conflict"`

	// Owner (name) cards only: how many of the owner's parcels carry this address.
	ParcelsCiting      float64 `json:"parcels_citing"`
	ParcelsCitingBasis string  `json:"parcels_citing_basis"`

	// The address as one mailing label.
	Label     string        `json:"label"`
	Source    ContactSource `json:"source"`
	AsOf      string        `json:"as_of"`
	AsOfBasis string        `json:"as_of_basis"`

	// A = the authority's own complete record; D = contradictory (hidden unless
	// include_low_confidence=true).
	Grade MailingAddressGrade `json:"grade"`
}

// UnmarshalJSON decodes MailingAddress, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *MailingAddress) UnmarshalJSON(data []byte) error {
	type plain MailingAddress
	aux := struct {
		*plain
		ParcelsCiting lenientNumber[float64] `json:"parcels_citing"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ParcelsCiting.assign(&r.ParcelsCiting)
	return softTypeError(err)
}

// ContactSource is generated from the OpenAPI spec.
type ContactSource struct {
	// Who published the value, e.g. the county assessor.
	Authority string  `json:"authority"`
	Dataset   string  `json:"dataset"`
	URL       *string `json:"url,omitempty"`
}

// MailingAddressGrade: A = the authority's own complete record; D = contradictory (hidden unless
// include_low_confidence=true).
//
// It is a string; the MailingAddressGrade* constants list the documented values.
type MailingAddressGrade = string

// Documented values of MailingAddressGrade.
const (
	MailingAddressGradeA MailingAddressGrade = "A"
	MailingAddressGradeB MailingAddressGrade = "B"
	MailingAddressGradeC MailingAddressGrade = "C"
	MailingAddressGradeD MailingAddressGrade = "D"
)

// LeadOwnerContactPhoneStatus is generated from the OpenAPI spec. It is a string; the
// LeadOwnerContactPhoneStatus* constants list the documented values.
type LeadOwnerContactPhoneStatus = string

// Documented values of LeadOwnerContactPhoneStatus.
const (
	LeadOwnerContactPhoneStatusNotChecked LeadOwnerContactPhoneStatus = "not_checked"
)

// LeadFeedPreview: The FREE preview (preview=true). `count` and `quote` are exactly what a paid
// call would deliver and charge.
type LeadFeedPreview struct {
	Signal string `json:"signal"`

	// The resolved request geography and filters.
	Geo LeadFeedPreviewGeo `json:"geo"`

	// Leads that will be / were delivered: min(matching rows, limit). This is the priced quantity.
	Count int64 `json:"count"`

	// True when the count stopped at `limit` -- more leads exist beyond this pull.
	CountCappedAtLimit bool       `json:"count_capped_at_limit"`
	Quote              LeadsQuote `json:"quote"`
	Preview            bool       `json:"preview"`

	// Up to three MASKED leads: APN truncated to state:county, house number stripped, owner name
	// withheld. Enough to judge the set, not enough to work it.
	Sample []Lead `json:"sample"`
	Note   string `json:"note"`

	// Present only when a known data gap explains an empty result (e.g. the owner-portfolio rollup's
	// unpopulated state columns). Nothing is charged in that case.
	CoverageNote         *string               `json:"coverage_note,omitempty"`
	TaxDelinquencyFilter *TaxDelinquencyFilter `json:"tax_delinquency_filter,omitempty"`
}

// UnmarshalJSON decodes LeadFeedPreview, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LeadFeedPreview) UnmarshalJSON(data []byte) error {
	type plain LeadFeedPreview
	aux := struct {
		*plain
		Count lenientNumber[int64] `json:"count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Count.assign(&r.Count)
	return softTypeError(err)
}

// LeadFeedPreviewGeo: The resolved request geography and filters.
type LeadFeedPreviewGeo struct {
	State         string  `json:"state"`
	StateFIPS     string  `json:"state_fips"`
	CountyFIPS    *string `json:"county_fips,omitempty"`
	Zip           *string `json:"zip,omitempty"`
	ValueMin      *int64  `json:"value_min,omitempty"`
	ValueMax      *int64  `json:"value_max,omitempty"`
	MailReady     bool    `json:"mail_ready"`
	Limit         int64   `json:"limit"`
	TaxDelinquent *bool   `json:"tax_delinquent,omitempty"`
}

// UnmarshalJSON decodes LeadFeedPreviewGeo, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LeadFeedPreviewGeo) UnmarshalJSON(data []byte) error {
	type plain LeadFeedPreviewGeo
	aux := struct {
		*plain
		ValueMin lenientNumber[int64] `json:"value_min"`
		ValueMax lenientNumber[int64] `json:"value_max"`
		Limit    lenientNumber[int64] `json:"limit"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ValueMin.assignPtr(&r.ValueMin)
	aux.ValueMax.assignPtr(&r.ValueMax)
	aux.Limit.assign(&r.Limit)
	return softTypeError(err)
}

// TaxDelinquencyFilter: Present when `tax_delinquent=true` was applied.
type TaxDelinquencyFilter struct {
	Applied  bool     `json:"applied"`
	Statuses []string `json:"statuses"`
	Note     string   `json:"note"`
}

// LeadFeed: The delivered lead set (paid), or an honest empty result when nothing matched (count
// 0, quote total $0, no payment taken).
type LeadFeed struct {
	Signal *string `json:"signal,omitempty"`

	// The resolved request geography and filters.
	Geo *LeadFeedGeo `json:"geo,omitempty"`

	// Leads that will be / were delivered: min(matching rows, limit). This is the priced quantity.
	Count *int64 `json:"count,omitempty"`

	// True when the count stopped at `limit` -- more leads exist beyond this pull.
	CountCappedAtLimit *bool       `json:"count_capped_at_limit,omitempty"`
	Quote              *LeadsQuote `json:"quote,omitempty"`

	// Present only when a known data gap explains an empty result (e.g. the owner-portfolio rollup's
	// unpopulated state columns). Nothing is charged in that case.
	CoverageNote *string `json:"coverage_note,omitempty"`
	Note         *string `json:"note,omitempty"`

	// Absent on an unpaid empty result (count 0).
	PaidVia *LeadFeedPaidVia `json:"paid_via,omitempty"`

	// The delivered, UNMASKED leads.
	Leads                []Lead                `json:"leads,omitempty"`
	TaxDelinquencyFilter *TaxDelinquencyFilter `json:"tax_delinquency_filter,omitempty"`
}

// UnmarshalJSON decodes LeadFeed, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LeadFeed) UnmarshalJSON(data []byte) error {
	type plain LeadFeed
	aux := struct {
		*plain
		Count lenientNumber[int64] `json:"count"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Count.assignPtr(&r.Count)
	return softTypeError(err)
}

// LeadFeedGeo: The resolved request geography and filters.
type LeadFeedGeo struct {
	State         *string `json:"state,omitempty"`
	StateFIPS     *string `json:"state_fips,omitempty"`
	CountyFIPS    *string `json:"county_fips,omitempty"`
	Zip           *string `json:"zip,omitempty"`
	ValueMin      *int64  `json:"value_min,omitempty"`
	ValueMax      *int64  `json:"value_max,omitempty"`
	Limit         *int64  `json:"limit,omitempty"`
	MailReady     *bool   `json:"mail_ready,omitempty"`
	TaxDelinquent *bool   `json:"tax_delinquent,omitempty"`
}

// UnmarshalJSON decodes LeadFeedGeo, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *LeadFeedGeo) UnmarshalJSON(data []byte) error {
	type plain LeadFeedGeo
	aux := struct {
		*plain
		ValueMin lenientNumber[int64] `json:"value_min"`
		ValueMax lenientNumber[int64] `json:"value_max"`
		Limit    lenientNumber[int64] `json:"limit"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ValueMin.assignPtr(&r.ValueMin)
	aux.ValueMax.assignPtr(&r.ValueMax)
	aux.Limit.assignPtr(&r.Limit)
	return softTypeError(err)
}

// LeadFeedPaidVia: Absent on an unpaid empty result (count 0).
//
// It is a string; the LeadFeedPaidVia* constants list the documented values.
type LeadFeedPaidVia = string

// Documented values of LeadFeedPaidVia.
const (
	LeadFeedPaidViaX402         LeadFeedPaidVia = "x402"
	LeadFeedPaidViaCredits      LeadFeedPaidVia = "credits"
	LeadFeedPaidViaSubscription LeadFeedPaidVia = "subscription"
)

// ServiceUnavailable: A refusal that took nothing: the work was not done and nothing was charged.
// Retry after `Retry-After`.
type ServiceUnavailable struct {
	Error string `json:"error"`
}

// CoOwner: A co-owner of record: another owner the SAME assessor record names beside the owner
// (OWNER2, ownname2, ADD_OWNER ...), as the county published it. Same source and as-of date as the
// owner name. Account required, like every people field.
type CoOwner struct {
	Name      *string        `json:"name,omitempty"`
	Basis     *CoOwnerBasis  `json:"basis,omitempty"`
	Source    *ContactSource `json:"source,omitempty"`
	AsOf      *string        `json:"as_of,omitempty"`
	AsOfBasis *string        `json:"as_of_basis,omitempty"`
	Grade     *CoOwnerGrade  `json:"grade,omitempty"`
}

// CoOwnerBasis is generated from the OpenAPI spec. It is a string; the CoOwnerBasis* constants
// list the documented values.
type CoOwnerBasis = string

// Documented values of CoOwnerBasis.
const (
	CoOwnerBasisAssessorCoOwner CoOwnerBasis = "assessor_co_owner"
)

// CoOwnerGrade is generated from the OpenAPI spec. It is a string; the CoOwnerGrade* constants
// list the documented values.
type CoOwnerGrade = string

// Documented values of CoOwnerGrade.
const (
	CoOwnerGradeA CoOwnerGrade = "A"
	CoOwnerGradeB CoOwnerGrade = "B"
	CoOwnerGradeC CoOwnerGrade = "C"
	CoOwnerGradeD CoOwnerGrade = "D"
)

// OwnerCard: The owner card: the owner of record and how to reach them by mail, from what the
// publishing authorities released.
type OwnerCard struct {
	// { kind: "parcel", canonical_id } or { kind: "owner", parcels_considered, parcels_capped }.
	Subject OwnerCardSubject `json:"subject"`
	Owner   OwnerCardOwner   `json:"owner"`
	Contact OwnerCardContact `json:"contact"`
	Note    string           `json:"note"`
}

// OwnerCardSubject: { kind: "parcel", canonical_id } or { kind: "owner", parcels_considered,
// parcels_capped }.
type OwnerCardSubject struct {
	Kind        string `json:"kind"`
	CanonicalID string `json:"canonical_id"`
}

// OwnerCardOwner is generated from the OpenAPI spec.
type OwnerCardOwner struct {
	Name       string                   `json:"name"`
	NameStatus OwnerCardOwnerNameStatus `json:"name_status"`
	EntityType string                   `json:"entity_type"`
	Roles      []json.RawMessage        `json:"roles"`
}

// OwnerCardOwnerNameStatus is generated from the OpenAPI spec. It is a string; the
// OwnerCardOwnerNameStatus* constants list the documented values.
type OwnerCardOwnerNameStatus = string

// Documented values of OwnerCardOwnerNameStatus.
const (
	OwnerCardOwnerNameStatusPresent      OwnerCardOwnerNameStatus = "present"
	OwnerCardOwnerNameStatusMissing      OwnerCardOwnerNameStatus = "missing"
	OwnerCardOwnerNameStatusPlaceholder  OwnerCardOwnerNameStatus = "placeholder"
	OwnerCardOwnerNameStatusConfidential OwnerCardOwnerNameStatus = "confidential"
)

// OwnerCardContact is generated from the OpenAPI spec.
type OwnerCardContact struct {
	OwnerName           string                              `json:"owner_name"`
	OwnerNameStatus     string                              `json:"owner_name_status"`
	OwnerRoles          []json.RawMessage                   `json:"owner_roles"`
	OwnerNameProvenance OwnerCardContactOwnerNameProvenance `json:"owner_name_provenance"`

	// The other owners the same assessor record names, in the record's order; never the owner again.
	// Name mode: across the side-read parcels, distinct by name.
	CoOwners []CoOwner `json:"co_owners"`

	// none_listed is claimed only for a record whose state's co-owners were loaded from the release
	// that serves it; not_checked when the lookup could not run, the state is not loaded yet, or the
	// loaded row was read beside a different owner.
	CoOwnerStatus OwnerCardContactCoOwnerStatus `json:"co_owner_status"`
	Mailing       MailingAddress                `json:"mailing"`

	// Parcel mode: a latest-deed grantee address naming the same owner, when it differs.
	MailingAlternates []MailingAddress `json:"mailing_alternates"`

	// Secretary of State principals (entity type, status, registered agent, officers) when the owner
	// is an entity.
	Entity OwnerCardContactEntity `json:"entity"`

	// OWNER phones (owner role only) published on a building permit filed in the current owner's era
	// and naming them (grade C): { e164, display, ext, phone_raw, role_basis, permit_ref, source,
	// as_of, as_of_basis, grade }.
	Phones []map[string]any `json:"phones"`

	// none_published is claimed only after the permit lookup ran; not_checked when it could not
	// (timeout, no acquisition date, no owner name).
	PhoneStatus         OwnerCardContactPhoneStatus `json:"phone_status"`
	Emails              []map[string]any            `json:"emails"`
	NonePublished       []*string                   `json:"none_published"`
	HiddenLowConfidence float64                     `json:"hidden_low_confidence"`

	// Applicant and contractor phones on the parcel's permits (name mode: across the side-read
	// parcels), newest first, one per (role, number), at most 10: { role: applicant|contractor,
	// role_basis, name, e164, display, ext, phone_raw, permit_ref, era:
	// current_owner|prior_owner|unknown, source, as_of, as_of_basis, grade }. Never the owner's phone.
	PeopleOnPermits []map[string]any `json:"people_on_permits"`

	// none_published only after the parcel's permits were read in full; not_checked when the read did
	// not run or stopped at its cap.
	PeopleOnPermitsStatus OwnerCardContactPeopleOnPermitsStatus `json:"people_on_permits_status"`
	OtherAddresses        []OwnerCardContactOtherAddresses      `json:"other_addresses"`
	OtherAddressesStatus  string                                `json:"other_addresses_status"`
	OtherAddressesScope   OwnerCardContactOtherAddressesScope   `json:"other_addresses_scope"`

	// Name mode: { parcels_checked, parcels_considered }.
	CoOwnerScope map[string]any `json:"co_owner_scope,omitempty"`

	// Name mode: the owner's distinct mailing addresses, most-cited first.
	MailingAddresses []MailingAddress `json:"mailing_addresses,omitempty"`

	// Name mode: { parcels_checked, parcels_considered }.
	PhoneScope map[string]any `json:"phone_scope,omitempty"`
}

// UnmarshalJSON decodes OwnerCardContact, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnerCardContact) UnmarshalJSON(data []byte) error {
	type plain OwnerCardContact
	aux := struct {
		*plain
		HiddenLowConfidence lenientNumber[float64] `json:"hidden_low_confidence"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.HiddenLowConfidence.assign(&r.HiddenLowConfidence)
	return softTypeError(err)
}

// OwnerCardContactOwnerNameProvenance is generated from the OpenAPI spec.
type OwnerCardContactOwnerNameProvenance struct {
	Source    OwnerCardContactOwnerNameProvenanceSource `json:"source"`
	AsOf      string                                    `json:"as_of"`
	AsOfBasis string                                    `json:"as_of_basis"`
	Grade     string                                    `json:"grade"`
}

// OwnerCardContactOwnerNameProvenanceSource is generated from the OpenAPI spec.
type OwnerCardContactOwnerNameProvenanceSource struct {
	Authority string  `json:"authority"`
	Dataset   string  `json:"dataset"`
	URL       *string `json:"url,omitempty"`
}

// OwnerCardContactCoOwnerStatus: none_listed is claimed only for a record whose state's co-owners
// were loaded from the release that serves it; not_checked when the lookup could not run, the
// state is not loaded yet, or the loaded row was read beside a different owner.
//
// It is a string; the OwnerCardContactCoOwnerStatus* constants list the documented values.
type OwnerCardContactCoOwnerStatus = string

// Documented values of OwnerCardContactCoOwnerStatus.
const (
	OwnerCardContactCoOwnerStatusListed     OwnerCardContactCoOwnerStatus = "listed"
	OwnerCardContactCoOwnerStatusNoneListed OwnerCardContactCoOwnerStatus = "none_listed"
	OwnerCardContactCoOwnerStatusNotChecked OwnerCardContactCoOwnerStatus = "not_checked"
)

// OwnerCardContactEntity: Secretary of State principals (entity type, status, registered agent,
// officers) when the owner is an entity.
type OwnerCardContactEntity struct {
	EntityType       string                       `json:"entity_type"`
	Status           *string                      `json:"status,omitempty"`
	StateOfFormation *string                      `json:"state_of_formation,omitempty"`
	FormationDate    *string                      `json:"formation_date,omitempty"`
	RegisteredAgent  *string                      `json:"registered_agent,omitempty"`
	Officers         []json.RawMessage            `json:"officers"`
	Source           OwnerCardContactEntitySource `json:"source"`
	AsOf             *string                      `json:"as_of,omitempty"`
	AsOfBasis        *string                      `json:"as_of_basis,omitempty"`
	Grade            string                       `json:"grade"`
}

// OwnerCardContactEntitySource is generated from the OpenAPI spec.
type OwnerCardContactEntitySource struct {
	Authority string  `json:"authority"`
	Dataset   string  `json:"dataset"`
	URL       *string `json:"url,omitempty"`
}

// OwnerCardContactPhoneStatus: none_published is claimed only after the permit lookup ran;
// not_checked when it could not (timeout, no acquisition date, no owner name).
//
// It is a string; the OwnerCardContactPhoneStatus* constants list the documented values.
type OwnerCardContactPhoneStatus = string

// Documented values of OwnerCardContactPhoneStatus.
const (
	OwnerCardContactPhoneStatusPublished     OwnerCardContactPhoneStatus = "published"
	OwnerCardContactPhoneStatusNonePublished OwnerCardContactPhoneStatus = "none_published"
	OwnerCardContactPhoneStatusNotChecked    OwnerCardContactPhoneStatus = "not_checked"
)

// OwnerCardContactPeopleOnPermitsStatus: none_published only after the parcel's permits were read
// in full; not_checked when the read did not run or stopped at its cap.
//
// It is a string; the OwnerCardContactPeopleOnPermitsStatus* constants list the documented values.
type OwnerCardContactPeopleOnPermitsStatus = string

// Documented values of OwnerCardContactPeopleOnPermitsStatus.
const (
	OwnerCardContactPeopleOnPermitsStatusListed        OwnerCardContactPeopleOnPermitsStatus = "listed"
	OwnerCardContactPeopleOnPermitsStatusNonePublished OwnerCardContactPeopleOnPermitsStatus = "none_published"
	OwnerCardContactPeopleOnPermitsStatusNotChecked    OwnerCardContactPeopleOnPermitsStatus = "not_checked"
)

// OwnerCardContactOtherAddresses is generated from the OpenAPI spec.
type OwnerCardContactOtherAddresses struct {
	ParcelID          string                                  `json:"parcel_id"`
	State             *string                                 `json:"state,omitempty"`
	County            *string                                 `json:"county,omitempty"`
	Kind              *string                                 `json:"kind,omitempty"`
	Address           OwnerCardContactOtherAddressesAddress   `json:"address"`
	SameAs            *string                                 `json:"same_as,omitempty"`
	Grade             *string                                 `json:"grade,omitempty"`
	Link              *string                                 `json:"link,omitempty"`
	Basis             *string                                 `json:"basis,omitempty"`
	LabelNote         *string                                 `json:"label_note,omitempty"`
	Evidence          []json.RawMessage                       `json:"evidence"`
	MailMerge         *bool                                   `json:"mail_merge,omitempty"`
	OwnerNameOnRecord *string                                 `json:"owner_name_on_record,omitempty"`
	OwnerRoles        []json.RawMessage                       `json:"owner_roles"`
	CitedBy           []OwnerCardContactOtherAddressesCitedBy `json:"cited_by"`
	ParcelsCiting     *float64                                `json:"parcels_citing,omitempty"`
	Source            OwnerCardContactOtherAddressesSource    `json:"source"`
	AsOf              *string                                 `json:"as_of,omitempty"`
	AsOfBasis         string                                  `json:"as_of_basis"`
}

// UnmarshalJSON decodes OwnerCardContactOtherAddresses, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnerCardContactOtherAddresses) UnmarshalJSON(data []byte) error {
	type plain OwnerCardContactOtherAddresses
	aux := struct {
		*plain
		ParcelsCiting lenientNumber[float64] `json:"parcels_citing"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.ParcelsCiting.assignPtr(&r.ParcelsCiting)
	return softTypeError(err)
}

// OwnerCardContactOtherAddressesAddress is generated from the OpenAPI spec.
type OwnerCardContactOtherAddressesAddress struct {
	Line1     *string `json:"line1,omitempty"`
	City      *string `json:"city,omitempty"`
	State     *string `json:"state,omitempty"`
	Zip5      *string `json:"zip5,omitempty"`
	Zip4      *string `json:"zip4,omitempty"`
	Label     *string `json:"label,omitempty"`
	MailReady *bool   `json:"mail_ready,omitempty"`
	PoBox     *bool   `json:"po_box,omitempty"`
}

// OwnerCardContactOtherAddressesCitedBy is generated from the OpenAPI spec.
type OwnerCardContactOtherAddressesCitedBy struct {
	ParcelID string  `json:"parcel_id"`
	State    *string `json:"state,omitempty"`
	County   *string `json:"county,omitempty"`
}

// OwnerCardContactOtherAddressesSource is generated from the OpenAPI spec.
type OwnerCardContactOtherAddressesSource struct {
	Authority *string `json:"authority,omitempty"`
	Dataset   *string `json:"dataset,omitempty"`
	URL       *string `json:"url,omitempty"`
}

// OwnerCardContactOtherAddressesScope is generated from the OpenAPI spec.
type OwnerCardContactOtherAddressesScope struct {
	NameBasis              string                                                 `json:"name_basis"`
	Spellings              float64                                                `json:"spellings"`
	ParcelsRead            float64                                                `json:"parcels_read"`
	Capped                 bool                                                   `json:"capped"`
	CopiesSkipped          float64                                                `json:"copies_skipped"`
	ParcelsLinked          float64                                                `json:"parcels_linked"`
	PossibleFound          float64                                                `json:"possible_found"`
	PossibleListed         float64                                                `json:"possible_listed"`
	PossibleCap            float64                                                `json:"possible_cap"`
	PossibleWithheldReason *string                                                `json:"possible_withheld_reason,omitempty"`
	ListedCapped           bool                                                   `json:"listed_capped"`
	EvidenceProviders      []OwnerCardContactOtherAddressesScopeEvidenceProviders `json:"evidence_providers"`
	CoOwnerLink            string                                                 `json:"co_owner_link"`
}

// UnmarshalJSON decodes OwnerCardContactOtherAddressesScope, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *OwnerCardContactOtherAddressesScope) UnmarshalJSON(data []byte) error {
	type plain OwnerCardContactOtherAddressesScope
	aux := struct {
		*plain
		Spellings      lenientNumber[float64] `json:"spellings"`
		ParcelsRead    lenientNumber[float64] `json:"parcels_read"`
		CopiesSkipped  lenientNumber[float64] `json:"copies_skipped"`
		ParcelsLinked  lenientNumber[float64] `json:"parcels_linked"`
		PossibleFound  lenientNumber[float64] `json:"possible_found"`
		PossibleListed lenientNumber[float64] `json:"possible_listed"`
		PossibleCap    lenientNumber[float64] `json:"possible_cap"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Spellings.assign(&r.Spellings)
	aux.ParcelsRead.assign(&r.ParcelsRead)
	aux.CopiesSkipped.assign(&r.CopiesSkipped)
	aux.ParcelsLinked.assign(&r.ParcelsLinked)
	aux.PossibleFound.assign(&r.PossibleFound)
	aux.PossibleListed.assign(&r.PossibleListed)
	aux.PossibleCap.assign(&r.PossibleCap)
	return softTypeError(err)
}

// OwnerCardContactOtherAddressesScopeEvidenceProviders is generated from the OpenAPI spec.
type OwnerCardContactOtherAddressesScopeEvidenceProviders struct {
	ID     string  `json:"id"`
	Status *string `json:"status,omitempty"`
}

// TaxDelinquencyRecord: One property-tax delinquency record as the publishing treasurer / tax
// collector lists it, placed on this parcel by the publisher's own parcel id. Never served past
// `expires_on`.
type TaxDelinquencyRecord struct {
	RecordUid string `json:"record_uid"`

	// ops.sources id (`tax_<st>_<jurisdiction>_<dataset>`).
	SourceID         string `json:"source_id"`
	JurisdictionName string `json:"jurisdiction_name"`
	Publisher        string `json:"publisher"`

	// The parcel id exactly as the publisher prints it.
	PublisherParcelID string `json:"publisher_parcel_id"`

	// Only in_sale and delinquent count as delinquent.
	Status    TaxDelinquencyRecordStatus `json:"status"`
	StatusRaw *string                    `json:"status_raw,omitempty"`

	// Only where the publisher says so; null = not published.
	PaymentPlan *bool `json:"payment_plan,omitempty"`

	// Only where the publisher says so; null = not published.
	Bankruptcy      *bool   `json:"bankruptcy,omitempty"`
	TaxYears        []int64 `json:"tax_years,omitempty"`
	FirstTaxYear    *int64  `json:"first_tax_year,omitempty"`
	LastTaxYear     *int64  `json:"last_tax_year,omitempty"`
	YearsDelinquent *int64  `json:"years_delinquent,omitempty"`

	// Null when the list publishes no amount; what it is is `amount_basis`.
	AmountDue   *float64 `json:"amount_due,omitempty"`
	AmountBasis *string  `json:"amount_basis,omitempty"`
	SaleKind    *string  `json:"sale_kind,omitempty"`
	SaleDate    *string  `json:"sale_date,omitempty"`

	// The publisher's own date for the list.
	PublisherAsOf string `json:"publisher_as_of"`

	// When PropRaven pulled the list.
	AsOf            string                          `json:"as_of"`
	ExpiresOn       string                          `json:"expires_on"`
	MatchMethod     TaxDelinquencyRecordMatchMethod `json:"match_method"`
	MatchConfidence float64                         `json:"match_confidence"`
	SourceURL       string                          `json:"source_url"`
}

// UnmarshalJSON decodes TaxDelinquencyRecord, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *TaxDelinquencyRecord) UnmarshalJSON(data []byte) error {
	type plain TaxDelinquencyRecord
	aux := struct {
		*plain
		TaxYears        lenientSlice[int64]    `json:"tax_years"`
		FirstTaxYear    lenientNumber[int64]   `json:"first_tax_year"`
		LastTaxYear     lenientNumber[int64]   `json:"last_tax_year"`
		YearsDelinquent lenientNumber[int64]   `json:"years_delinquent"`
		AmountDue       lenientNumber[float64] `json:"amount_due"`
		MatchConfidence lenientNumber[float64] `json:"match_confidence"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.TaxYears.assign(&r.TaxYears)
	aux.FirstTaxYear.assignPtr(&r.FirstTaxYear)
	aux.LastTaxYear.assignPtr(&r.LastTaxYear)
	aux.YearsDelinquent.assignPtr(&r.YearsDelinquent)
	aux.AmountDue.assignPtr(&r.AmountDue)
	aux.MatchConfidence.assign(&r.MatchConfidence)
	return softTypeError(err)
}

// TaxDelinquencyRecordStatus: Only in_sale and delinquent count as delinquent.
//
// It is a string; the TaxDelinquencyRecordStatus* constants list the documented values.
type TaxDelinquencyRecordStatus = string

// Documented values of TaxDelinquencyRecordStatus.
const (
	TaxDelinquencyRecordStatusInSale     TaxDelinquencyRecordStatus = "in_sale"
	TaxDelinquencyRecordStatusDelinquent TaxDelinquencyRecordStatus = "delinquent"
	TaxDelinquencyRecordStatusSold       TaxDelinquencyRecordStatus = "sold"
	TaxDelinquencyRecordStatusRedeemed   TaxDelinquencyRecordStatus = "redeemed"
)

// TaxDelinquencyRecordMatchMethod is generated from the OpenAPI spec. It is a string; the
// TaxDelinquencyRecordMatchMethod* constants list the documented values.
type TaxDelinquencyRecordMatchMethod = string

// Documented values of TaxDelinquencyRecordMatchMethod.
const (
	TaxDelinquencyRecordMatchMethodParcelIDExact      TaxDelinquencyRecordMatchMethod = "parcel_id_exact"
	TaxDelinquencyRecordMatchMethodParcelIDNormalized TaxDelinquencyRecordMatchMethod = "parcel_id_normalized"
	TaxDelinquencyRecordMatchMethodAddressExact       TaxDelinquencyRecordMatchMethod = "address_exact"
)

// TaxDelinquencyCoverage: A delinquency list that covers the parcel's county. `list_scope` says
// what absence from the list means.
type TaxDelinquencyCoverage struct {
	SourceID         string                            `json:"source_id"`
	JurisdictionName string                            `json:"jurisdiction_name"`
	Publisher        string                            `json:"publisher"`
	DatasetKind      TaxDelinquencyCoverageDatasetKind `json:"dataset_kind"`
	ListScope        string                            `json:"list_scope"`
	PublisherAsOf    string                            `json:"publisher_as_of"`
	AsOf             string                            `json:"as_of"`
	ExpiresOn        string                            `json:"expires_on"`
	SourceURL        string                            `json:"source_url"`
}

// TaxDelinquencyCoverageDatasetKind is generated from the OpenAPI spec. It is a string; the
// TaxDelinquencyCoverageDatasetKind* constants list the documented values.
type TaxDelinquencyCoverageDatasetKind = string

// Documented values of TaxDelinquencyCoverageDatasetKind.
const (
	TaxDelinquencyCoverageDatasetKindDelinquency     TaxDelinquencyCoverageDatasetKind = "delinquency"
	TaxDelinquencyCoverageDatasetKindLienSaleList    TaxDelinquencyCoverageDatasetKind = "lien_sale_list"
	TaxDelinquencyCoverageDatasetKindTaxSaleList     TaxDelinquencyCoverageDatasetKind = "tax_sale_list"
	TaxDelinquencyCoverageDatasetKindScavengerList   TaxDelinquencyCoverageDatasetKind = "scavenger_list"
	TaxDelinquencyCoverageDatasetKindForeclosureList TaxDelinquencyCoverageDatasetKind = "foreclosure_list"
)

// ParcelTaxStatus is generated from the OpenAPI spec.
type ParcelTaxStatus struct {
	ParcelID string `json:"parcel_id"`

	// listed = on at least one unexpired list; not_listed = an unexpired list covers the county and
	// this parcel is not on it (read `list_scope`: never proof of payment); not_covered = no list for
	// the county; unavailable = the layer could not be read.
	Status   ParcelTaxStatusStatus    `json:"status"`
	Coverage []TaxDelinquencyCoverage `json:"coverage"`

	// in_sale, delinquent, sold, redeemed, then amount_due descending; at most 50.
	Records   []TaxDelinquencyRecord `json:"records"`
	Truncated bool                   `json:"truncated"`
	Note      string                 `json:"note"`

	// Present when the records were withheld by the lookup meter: `code` lookup_cap_reached or
	// lookup_meter_unavailable (nothing charged).
	PeopleFields map[string]any `json:"people_fields,omitempty"`
}

// ParcelTaxStatusStatus: listed = on at least one unexpired list; not_listed = an unexpired list
// covers the county and this parcel is not on it (read `list_scope`: never proof of payment);
// not_covered = no list for the county; unavailable = the layer could not be read.
//
// It is a string; the ParcelTaxStatusStatus* constants list the documented values.
type ParcelTaxStatusStatus = string

// Documented values of ParcelTaxStatusStatus.
const (
	ParcelTaxStatusStatusListed      ParcelTaxStatusStatus = "listed"
	ParcelTaxStatusStatusNotListed   ParcelTaxStatusStatus = "not_listed"
	ParcelTaxStatusStatusNotCovered  ParcelTaxStatusStatus = "not_covered"
	ParcelTaxStatusStatusUnavailable ParcelTaxStatusStatus = "unavailable"
)

// IntelligenceRun: Immutable descriptive calculation run with exact rational wire values.
// Calculation time is not source freshness. Unknown values/dates remain null.
type IntelligenceRun struct {
	ID                string                `json:"id"`
	CanonicalID       string                `json:"canonical_id"`
	DefinitionVersion string                `json:"definition_version"`
	Query             IntelligenceRunQuery  `json:"query"`
	ComputedAt        string                `json:"computed_at"`
	EvidenceIDs       []string              `json:"evidence_ids"`
	Groups            IntelligenceRunGroups `json:"groups"`
}

// IntelligenceRunQuery is generated from the OpenAPI spec.
type IntelligenceRunQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceRunGroups is generated from the OpenAPI spec.
type IntelligenceRunGroups struct {
	Market        []IntelligenceRunGroupsMarket        `json:"market"`
	Seller        []IntelligenceRunGroupsSeller        `json:"seller"`
	Owner         []IntelligenceRunGroupsOwner         `json:"owner"`
	Redevelopment []IntelligenceRunGroupsRedevelopment `json:"redevelopment"`
}

// IntelligenceRunGroupsMarket is generated from the OpenAPI spec.
type IntelligenceRunGroupsMarket struct {
	Definition  string                                             `json:"definition"`
	Status      IntelligenceRunGroupsMarketStatus                  `json:"status"`
	Reasons     []string                                           `json:"reasons"`
	Metrics     map[string]IntelligenceRunGroupsMarketMetricsValue `json:"metrics"`
	Counts      map[string]float64                                 `json:"counts"`
	Exclusions  map[string]float64                                 `json:"exclusions"`
	EvidenceIDs []string                                           `json:"evidence_ids"`
	Context     IntelligenceRunGroupsMarketContext                 `json:"context"`
}

// UnmarshalJSON decodes IntelligenceRunGroupsMarket, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *IntelligenceRunGroupsMarket) UnmarshalJSON(data []byte) error {
	type plain IntelligenceRunGroupsMarket
	aux := struct {
		*plain
		Counts     lenientMap[float64] `json:"counts"`
		Exclusions lenientMap[float64] `json:"exclusions"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Counts.assign(&r.Counts)
	aux.Exclusions.assign(&r.Exclusions)
	return softTypeError(err)
}

// IntelligenceRunGroupsMarketStatus is generated from the OpenAPI spec. It is a string; the
// IntelligenceRunGroupsMarketStatus* constants list the documented values.
type IntelligenceRunGroupsMarketStatus = string

// Documented values of IntelligenceRunGroupsMarketStatus.
const (
	IntelligenceRunGroupsMarketStatusAvailable        IntelligenceRunGroupsMarketStatus = "available"
	IntelligenceRunGroupsMarketStatusPartial          IntelligenceRunGroupsMarketStatus = "partial"
	IntelligenceRunGroupsMarketStatusInsufficientData IntelligenceRunGroupsMarketStatus = "insufficient_data"
	IntelligenceRunGroupsMarketStatusUnavailable      IntelligenceRunGroupsMarketStatus = "unavailable"
	IntelligenceRunGroupsMarketStatusStale            IntelligenceRunGroupsMarketStatus = "stale"
	IntelligenceRunGroupsMarketStatusError            IntelligenceRunGroupsMarketStatus = "error"
)

// IntelligenceRunGroupsMarketMetricsValue is generated from the OpenAPI spec.
type IntelligenceRunGroupsMarketMetricsValue struct {
	Value *IntelligenceRunGroupsMarketMetricsValueValue `json:"value,omitempty"`
	Unit  string                                        `json:"unit"`
}

// IntelligenceRunGroupsMarketMetricsValueValue is generated from the OpenAPI spec.
type IntelligenceRunGroupsMarketMetricsValueValue struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
}

// IntelligenceRunGroupsMarketContext is generated from the OpenAPI spec.
type IntelligenceRunGroupsMarketContext struct {
	InputKind           IntelligenceRunGroupsMarketContextInputKind       `json:"input_kind"`
	RequestedScope      *IntelligenceRunGroupsMarketContextRequestedScope `json:"requested_scope,omitempty"`
	ActualScope         *IntelligenceRunGroupsMarketContextActualScope    `json:"actual_scope,omitempty"`
	Query               *IntelligenceRunGroupsMarketContextQuery          `json:"query,omitempty"`
	Period              json.RawMessage                                   `json:"period,omitempty"`
	MetricPeriods       map[string]json.RawMessage                        `json:"metric_periods"`
	LatestObservationAt *string                                           `json:"latest_observation_at,omitempty"`
	SourceAsOf          *string                                           `json:"source_as_of,omitempty"`
	SourceVintage       *string                                           `json:"source_vintage,omitempty"`
	HistoryComplete     *bool                                             `json:"history_complete,omitempty"`
}

// IntelligenceRunGroupsMarketContextInputKind is generated from the OpenAPI spec. It is a string;
// the IntelligenceRunGroupsMarketContextInputKind* constants list the documented values.
type IntelligenceRunGroupsMarketContextInputKind = string

// Documented values of IntelligenceRunGroupsMarketContextInputKind.
const (
	IntelligenceRunGroupsMarketContextInputKindObserved        IntelligenceRunGroupsMarketContextInputKind = "observed"
	IntelligenceRunGroupsMarketContextInputKindUserAssumptions IntelligenceRunGroupsMarketContextInputKind = "user_assumptions"
	IntelligenceRunGroupsMarketContextInputKindUnavailable     IntelligenceRunGroupsMarketContextInputKind = "unavailable"
)

// IntelligenceRunGroupsMarketContextRequestedScope is generated from the OpenAPI spec.
type IntelligenceRunGroupsMarketContextRequestedScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceRunGroupsMarketContextActualScope is generated from the OpenAPI spec.
type IntelligenceRunGroupsMarketContextActualScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceRunGroupsMarketContextQuery is generated from the OpenAPI spec.
type IntelligenceRunGroupsMarketContextQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceRunGroupsSeller is generated from the OpenAPI spec.
type IntelligenceRunGroupsSeller struct {
	Definition  string                                             `json:"definition"`
	Status      IntelligenceRunGroupsSellerStatus                  `json:"status"`
	Reasons     []string                                           `json:"reasons"`
	Metrics     map[string]IntelligenceRunGroupsSellerMetricsValue `json:"metrics"`
	Counts      map[string]float64                                 `json:"counts"`
	Exclusions  map[string]float64                                 `json:"exclusions"`
	EvidenceIDs []string                                           `json:"evidence_ids"`
	Context     IntelligenceRunGroupsSellerContext                 `json:"context"`
}

// UnmarshalJSON decodes IntelligenceRunGroupsSeller, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *IntelligenceRunGroupsSeller) UnmarshalJSON(data []byte) error {
	type plain IntelligenceRunGroupsSeller
	aux := struct {
		*plain
		Counts     lenientMap[float64] `json:"counts"`
		Exclusions lenientMap[float64] `json:"exclusions"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Counts.assign(&r.Counts)
	aux.Exclusions.assign(&r.Exclusions)
	return softTypeError(err)
}

// IntelligenceRunGroupsSellerStatus is generated from the OpenAPI spec. It is a string; the
// IntelligenceRunGroupsSellerStatus* constants list the documented values.
type IntelligenceRunGroupsSellerStatus = string

// Documented values of IntelligenceRunGroupsSellerStatus.
const (
	IntelligenceRunGroupsSellerStatusAvailable        IntelligenceRunGroupsSellerStatus = "available"
	IntelligenceRunGroupsSellerStatusPartial          IntelligenceRunGroupsSellerStatus = "partial"
	IntelligenceRunGroupsSellerStatusInsufficientData IntelligenceRunGroupsSellerStatus = "insufficient_data"
	IntelligenceRunGroupsSellerStatusUnavailable      IntelligenceRunGroupsSellerStatus = "unavailable"
	IntelligenceRunGroupsSellerStatusStale            IntelligenceRunGroupsSellerStatus = "stale"
	IntelligenceRunGroupsSellerStatusError            IntelligenceRunGroupsSellerStatus = "error"
)

// IntelligenceRunGroupsSellerMetricsValue is generated from the OpenAPI spec.
type IntelligenceRunGroupsSellerMetricsValue struct {
	Value *IntelligenceRunGroupsSellerMetricsValueValue `json:"value,omitempty"`
	Unit  string                                        `json:"unit"`
}

// IntelligenceRunGroupsSellerMetricsValueValue is generated from the OpenAPI spec.
type IntelligenceRunGroupsSellerMetricsValueValue struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
}

// IntelligenceRunGroupsSellerContext is generated from the OpenAPI spec.
type IntelligenceRunGroupsSellerContext struct {
	InputKind           IntelligenceRunGroupsSellerContextInputKind       `json:"input_kind"`
	RequestedScope      *IntelligenceRunGroupsSellerContextRequestedScope `json:"requested_scope,omitempty"`
	ActualScope         *IntelligenceRunGroupsSellerContextActualScope    `json:"actual_scope,omitempty"`
	Query               *IntelligenceRunGroupsSellerContextQuery          `json:"query,omitempty"`
	Period              json.RawMessage                                   `json:"period,omitempty"`
	MetricPeriods       map[string]json.RawMessage                        `json:"metric_periods"`
	LatestObservationAt *string                                           `json:"latest_observation_at,omitempty"`
	SourceAsOf          *string                                           `json:"source_as_of,omitempty"`
	SourceVintage       *string                                           `json:"source_vintage,omitempty"`
	HistoryComplete     *bool                                             `json:"history_complete,omitempty"`
}

// IntelligenceRunGroupsSellerContextInputKind is generated from the OpenAPI spec. It is a string;
// the IntelligenceRunGroupsSellerContextInputKind* constants list the documented values.
type IntelligenceRunGroupsSellerContextInputKind = string

// Documented values of IntelligenceRunGroupsSellerContextInputKind.
const (
	IntelligenceRunGroupsSellerContextInputKindObserved        IntelligenceRunGroupsSellerContextInputKind = "observed"
	IntelligenceRunGroupsSellerContextInputKindUserAssumptions IntelligenceRunGroupsSellerContextInputKind = "user_assumptions"
	IntelligenceRunGroupsSellerContextInputKindUnavailable     IntelligenceRunGroupsSellerContextInputKind = "unavailable"
)

// IntelligenceRunGroupsSellerContextRequestedScope is generated from the OpenAPI spec.
type IntelligenceRunGroupsSellerContextRequestedScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceRunGroupsSellerContextActualScope is generated from the OpenAPI spec.
type IntelligenceRunGroupsSellerContextActualScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceRunGroupsSellerContextQuery is generated from the OpenAPI spec.
type IntelligenceRunGroupsSellerContextQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceRunGroupsOwner is generated from the OpenAPI spec.
type IntelligenceRunGroupsOwner struct {
	Definition  string                                            `json:"definition"`
	Status      IntelligenceRunGroupsOwnerStatus                  `json:"status"`
	Reasons     []string                                          `json:"reasons"`
	Metrics     map[string]IntelligenceRunGroupsOwnerMetricsValue `json:"metrics"`
	Counts      map[string]float64                                `json:"counts"`
	Exclusions  map[string]float64                                `json:"exclusions"`
	EvidenceIDs []string                                          `json:"evidence_ids"`
	Context     IntelligenceRunGroupsOwnerContext                 `json:"context"`
}

// UnmarshalJSON decodes IntelligenceRunGroupsOwner, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *IntelligenceRunGroupsOwner) UnmarshalJSON(data []byte) error {
	type plain IntelligenceRunGroupsOwner
	aux := struct {
		*plain
		Counts     lenientMap[float64] `json:"counts"`
		Exclusions lenientMap[float64] `json:"exclusions"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Counts.assign(&r.Counts)
	aux.Exclusions.assign(&r.Exclusions)
	return softTypeError(err)
}

// IntelligenceRunGroupsOwnerStatus is generated from the OpenAPI spec. It is a string; the
// IntelligenceRunGroupsOwnerStatus* constants list the documented values.
type IntelligenceRunGroupsOwnerStatus = string

// Documented values of IntelligenceRunGroupsOwnerStatus.
const (
	IntelligenceRunGroupsOwnerStatusAvailable        IntelligenceRunGroupsOwnerStatus = "available"
	IntelligenceRunGroupsOwnerStatusPartial          IntelligenceRunGroupsOwnerStatus = "partial"
	IntelligenceRunGroupsOwnerStatusInsufficientData IntelligenceRunGroupsOwnerStatus = "insufficient_data"
	IntelligenceRunGroupsOwnerStatusUnavailable      IntelligenceRunGroupsOwnerStatus = "unavailable"
	IntelligenceRunGroupsOwnerStatusStale            IntelligenceRunGroupsOwnerStatus = "stale"
	IntelligenceRunGroupsOwnerStatusError            IntelligenceRunGroupsOwnerStatus = "error"
)

// IntelligenceRunGroupsOwnerMetricsValue is generated from the OpenAPI spec.
type IntelligenceRunGroupsOwnerMetricsValue struct {
	Value *IntelligenceRunGroupsOwnerMetricsValueValue `json:"value,omitempty"`
	Unit  string                                       `json:"unit"`
}

// IntelligenceRunGroupsOwnerMetricsValueValue is generated from the OpenAPI spec.
type IntelligenceRunGroupsOwnerMetricsValueValue struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
}

// IntelligenceRunGroupsOwnerContext is generated from the OpenAPI spec.
type IntelligenceRunGroupsOwnerContext struct {
	InputKind           IntelligenceRunGroupsOwnerContextInputKind       `json:"input_kind"`
	RequestedScope      *IntelligenceRunGroupsOwnerContextRequestedScope `json:"requested_scope,omitempty"`
	ActualScope         *IntelligenceRunGroupsOwnerContextActualScope    `json:"actual_scope,omitempty"`
	Query               *IntelligenceRunGroupsOwnerContextQuery          `json:"query,omitempty"`
	Period              json.RawMessage                                  `json:"period,omitempty"`
	MetricPeriods       map[string]json.RawMessage                       `json:"metric_periods"`
	LatestObservationAt *string                                          `json:"latest_observation_at,omitempty"`
	SourceAsOf          *string                                          `json:"source_as_of,omitempty"`
	SourceVintage       *string                                          `json:"source_vintage,omitempty"`
	HistoryComplete     *bool                                            `json:"history_complete,omitempty"`
}

// IntelligenceRunGroupsOwnerContextInputKind is generated from the OpenAPI spec. It is a string;
// the IntelligenceRunGroupsOwnerContextInputKind* constants list the documented values.
type IntelligenceRunGroupsOwnerContextInputKind = string

// Documented values of IntelligenceRunGroupsOwnerContextInputKind.
const (
	IntelligenceRunGroupsOwnerContextInputKindObserved        IntelligenceRunGroupsOwnerContextInputKind = "observed"
	IntelligenceRunGroupsOwnerContextInputKindUserAssumptions IntelligenceRunGroupsOwnerContextInputKind = "user_assumptions"
	IntelligenceRunGroupsOwnerContextInputKindUnavailable     IntelligenceRunGroupsOwnerContextInputKind = "unavailable"
)

// IntelligenceRunGroupsOwnerContextRequestedScope is generated from the OpenAPI spec.
type IntelligenceRunGroupsOwnerContextRequestedScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceRunGroupsOwnerContextActualScope is generated from the OpenAPI spec.
type IntelligenceRunGroupsOwnerContextActualScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceRunGroupsOwnerContextQuery is generated from the OpenAPI spec.
type IntelligenceRunGroupsOwnerContextQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceRunGroupsRedevelopment is generated from the OpenAPI spec.
type IntelligenceRunGroupsRedevelopment struct {
	Definition  string                                                    `json:"definition"`
	Status      IntelligenceRunGroupsRedevelopmentStatus                  `json:"status"`
	Reasons     []string                                                  `json:"reasons"`
	Metrics     map[string]IntelligenceRunGroupsRedevelopmentMetricsValue `json:"metrics"`
	Counts      map[string]float64                                        `json:"counts"`
	Exclusions  map[string]float64                                        `json:"exclusions"`
	EvidenceIDs []string                                                  `json:"evidence_ids"`
	Context     IntelligenceRunGroupsRedevelopmentContext                 `json:"context"`
}

// UnmarshalJSON decodes IntelligenceRunGroupsRedevelopment, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *IntelligenceRunGroupsRedevelopment) UnmarshalJSON(data []byte) error {
	type plain IntelligenceRunGroupsRedevelopment
	aux := struct {
		*plain
		Counts     lenientMap[float64] `json:"counts"`
		Exclusions lenientMap[float64] `json:"exclusions"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Counts.assign(&r.Counts)
	aux.Exclusions.assign(&r.Exclusions)
	return softTypeError(err)
}

// IntelligenceRunGroupsRedevelopmentStatus is generated from the OpenAPI spec. It is a string; the
// IntelligenceRunGroupsRedevelopmentStatus* constants list the documented values.
type IntelligenceRunGroupsRedevelopmentStatus = string

// Documented values of IntelligenceRunGroupsRedevelopmentStatus.
const (
	IntelligenceRunGroupsRedevelopmentStatusAvailable        IntelligenceRunGroupsRedevelopmentStatus = "available"
	IntelligenceRunGroupsRedevelopmentStatusPartial          IntelligenceRunGroupsRedevelopmentStatus = "partial"
	IntelligenceRunGroupsRedevelopmentStatusInsufficientData IntelligenceRunGroupsRedevelopmentStatus = "insufficient_data"
	IntelligenceRunGroupsRedevelopmentStatusUnavailable      IntelligenceRunGroupsRedevelopmentStatus = "unavailable"
	IntelligenceRunGroupsRedevelopmentStatusStale            IntelligenceRunGroupsRedevelopmentStatus = "stale"
	IntelligenceRunGroupsRedevelopmentStatusError            IntelligenceRunGroupsRedevelopmentStatus = "error"
)

// IntelligenceRunGroupsRedevelopmentMetricsValue is generated from the OpenAPI spec.
type IntelligenceRunGroupsRedevelopmentMetricsValue struct {
	Value *IntelligenceRunGroupsRedevelopmentMetricsValueValue `json:"value,omitempty"`
	Unit  string                                               `json:"unit"`
}

// IntelligenceRunGroupsRedevelopmentMetricsValueValue is generated from the OpenAPI spec.
type IntelligenceRunGroupsRedevelopmentMetricsValueValue struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
}

// IntelligenceRunGroupsRedevelopmentContext is generated from the OpenAPI spec.
type IntelligenceRunGroupsRedevelopmentContext struct {
	InputKind           IntelligenceRunGroupsRedevelopmentContextInputKind       `json:"input_kind"`
	RequestedScope      *IntelligenceRunGroupsRedevelopmentContextRequestedScope `json:"requested_scope,omitempty"`
	ActualScope         *IntelligenceRunGroupsRedevelopmentContextActualScope    `json:"actual_scope,omitempty"`
	Query               *IntelligenceRunGroupsRedevelopmentContextQuery          `json:"query,omitempty"`
	Period              json.RawMessage                                          `json:"period,omitempty"`
	MetricPeriods       map[string]json.RawMessage                               `json:"metric_periods"`
	LatestObservationAt *string                                                  `json:"latest_observation_at,omitempty"`
	SourceAsOf          *string                                                  `json:"source_as_of,omitempty"`
	SourceVintage       *string                                                  `json:"source_vintage,omitempty"`
	HistoryComplete     *bool                                                    `json:"history_complete,omitempty"`
}

// IntelligenceRunGroupsRedevelopmentContextInputKind is generated from the OpenAPI spec. It is a
// string; the IntelligenceRunGroupsRedevelopmentContextInputKind* constants list the documented
// values.
type IntelligenceRunGroupsRedevelopmentContextInputKind = string

// Documented values of IntelligenceRunGroupsRedevelopmentContextInputKind.
const (
	IntelligenceRunGroupsRedevelopmentContextInputKindObserved        IntelligenceRunGroupsRedevelopmentContextInputKind = "observed"
	IntelligenceRunGroupsRedevelopmentContextInputKindUserAssumptions IntelligenceRunGroupsRedevelopmentContextInputKind = "user_assumptions"
	IntelligenceRunGroupsRedevelopmentContextInputKindUnavailable     IntelligenceRunGroupsRedevelopmentContextInputKind = "unavailable"
)

// IntelligenceRunGroupsRedevelopmentContextRequestedScope is generated from the OpenAPI spec.
type IntelligenceRunGroupsRedevelopmentContextRequestedScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceRunGroupsRedevelopmentContextActualScope is generated from the OpenAPI spec.
type IntelligenceRunGroupsRedevelopmentContextActualScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceRunGroupsRedevelopmentContextQuery is generated from the OpenAPI spec.
type IntelligenceRunGroupsRedevelopmentContextQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceHandoff: Authorized retained results and optional user scenario;
// structured_only_not_sent. No model call, external delivery or duplicate arithmetic.
type IntelligenceHandoff struct {
	Version         string                              `json:"version"`
	CanonicalID     string                              `json:"canonical_id"`
	RunID           string                              `json:"run_id"`
	Query           IntelligenceHandoffQuery            `json:"query"`
	Observations    []IntelligenceHandoffObservations   `json:"observations"`
	ComputedResults IntelligenceHandoffComputedResults  `json:"computed_results"`
	UserAssumptions *IntelligenceHandoffUserAssumptions `json:"user_assumptions,omitempty"`
	Interpretations []string                            `json:"interpretations"`
	NextDiligence   []string                            `json:"next_diligence"`
	Delivery        string                              `json:"delivery"`
}

// IntelligenceHandoffQuery is generated from the OpenAPI spec.
type IntelligenceHandoffQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceHandoffObservations is generated from the OpenAPI spec.
type IntelligenceHandoffObservations struct {
	EvidenceID         string  `json:"evidence_id"`
	SourceProduct      string  `json:"source_product"`
	SourceVersion      string  `json:"source_version"`
	Vintage            string  `json:"vintage"`
	SourceAsOf         *string `json:"source_as_of,omitempty"`
	ObservedAt         *string `json:"observed_at,omitempty"`
	CapturedAt         string  `json:"captured_at"`
	KnowledgeBasis     string  `json:"knowledge_basis"`
	EffectiveTimeBasis string  `json:"effective_time_basis"`
	SourceURL          *string `json:"source_url,omitempty"`
}

// IntelligenceHandoffComputedResults is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResults struct {
	Market        []IntelligenceHandoffComputedResultsMarket        `json:"market"`
	Seller        []IntelligenceHandoffComputedResultsSeller        `json:"seller"`
	Owner         []IntelligenceHandoffComputedResultsOwner         `json:"owner"`
	Redevelopment []IntelligenceHandoffComputedResultsRedevelopment `json:"redevelopment"`
}

// IntelligenceHandoffComputedResultsMarket is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsMarket struct {
	Definition  string                                                          `json:"definition"`
	Status      IntelligenceHandoffComputedResultsMarketStatus                  `json:"status"`
	Reasons     []string                                                        `json:"reasons"`
	Metrics     map[string]IntelligenceHandoffComputedResultsMarketMetricsValue `json:"metrics"`
	Counts      map[string]float64                                              `json:"counts"`
	Exclusions  map[string]float64                                              `json:"exclusions"`
	EvidenceIDs []string                                                        `json:"evidence_ids"`
	Context     IntelligenceHandoffComputedResultsMarketContext                 `json:"context"`
}

// UnmarshalJSON decodes IntelligenceHandoffComputedResultsMarket, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *IntelligenceHandoffComputedResultsMarket) UnmarshalJSON(data []byte) error {
	type plain IntelligenceHandoffComputedResultsMarket
	aux := struct {
		*plain
		Counts     lenientMap[float64] `json:"counts"`
		Exclusions lenientMap[float64] `json:"exclusions"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Counts.assign(&r.Counts)
	aux.Exclusions.assign(&r.Exclusions)
	return softTypeError(err)
}

// IntelligenceHandoffComputedResultsMarketStatus is generated from the OpenAPI spec. It is a
// string; the IntelligenceHandoffComputedResultsMarketStatus* constants list the documented
// values.
type IntelligenceHandoffComputedResultsMarketStatus = string

// Documented values of IntelligenceHandoffComputedResultsMarketStatus.
const (
	IntelligenceHandoffComputedResultsMarketStatusAvailable        IntelligenceHandoffComputedResultsMarketStatus = "available"
	IntelligenceHandoffComputedResultsMarketStatusPartial          IntelligenceHandoffComputedResultsMarketStatus = "partial"
	IntelligenceHandoffComputedResultsMarketStatusInsufficientData IntelligenceHandoffComputedResultsMarketStatus = "insufficient_data"
	IntelligenceHandoffComputedResultsMarketStatusUnavailable      IntelligenceHandoffComputedResultsMarketStatus = "unavailable"
	IntelligenceHandoffComputedResultsMarketStatusStale            IntelligenceHandoffComputedResultsMarketStatus = "stale"
	IntelligenceHandoffComputedResultsMarketStatusError            IntelligenceHandoffComputedResultsMarketStatus = "error"
)

// IntelligenceHandoffComputedResultsMarketMetricsValue is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsMarketMetricsValue struct {
	Value *IntelligenceHandoffComputedResultsMarketMetricsValueValue `json:"value,omitempty"`
	Unit  string                                                     `json:"unit"`
}

// IntelligenceHandoffComputedResultsMarketMetricsValueValue is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsMarketMetricsValueValue struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
}

// IntelligenceHandoffComputedResultsMarketContext is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsMarketContext struct {
	InputKind           IntelligenceHandoffComputedResultsMarketContextInputKind       `json:"input_kind"`
	RequestedScope      *IntelligenceHandoffComputedResultsMarketContextRequestedScope `json:"requested_scope,omitempty"`
	ActualScope         *IntelligenceHandoffComputedResultsMarketContextActualScope    `json:"actual_scope,omitempty"`
	Query               *IntelligenceHandoffComputedResultsMarketContextQuery          `json:"query,omitempty"`
	Period              json.RawMessage                                                `json:"period,omitempty"`
	MetricPeriods       map[string]json.RawMessage                                     `json:"metric_periods"`
	LatestObservationAt *string                                                        `json:"latest_observation_at,omitempty"`
	SourceAsOf          *string                                                        `json:"source_as_of,omitempty"`
	SourceVintage       *string                                                        `json:"source_vintage,omitempty"`
	HistoryComplete     *bool                                                          `json:"history_complete,omitempty"`
}

// IntelligenceHandoffComputedResultsMarketContextInputKind is generated from the OpenAPI spec. It
// is a string; the IntelligenceHandoffComputedResultsMarketContextInputKind* constants list the
// documented values.
type IntelligenceHandoffComputedResultsMarketContextInputKind = string

// Documented values of IntelligenceHandoffComputedResultsMarketContextInputKind.
const (
	IntelligenceHandoffComputedResultsMarketContextInputKindObserved        IntelligenceHandoffComputedResultsMarketContextInputKind = "observed"
	IntelligenceHandoffComputedResultsMarketContextInputKindUserAssumptions IntelligenceHandoffComputedResultsMarketContextInputKind = "user_assumptions"
	IntelligenceHandoffComputedResultsMarketContextInputKindUnavailable     IntelligenceHandoffComputedResultsMarketContextInputKind = "unavailable"
)

// IntelligenceHandoffComputedResultsMarketContextRequestedScope is generated from the OpenAPI
// spec.
type IntelligenceHandoffComputedResultsMarketContextRequestedScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceHandoffComputedResultsMarketContextActualScope is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsMarketContextActualScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceHandoffComputedResultsMarketContextQuery is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsMarketContextQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceHandoffComputedResultsSeller is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsSeller struct {
	Definition  string                                                          `json:"definition"`
	Status      IntelligenceHandoffComputedResultsSellerStatus                  `json:"status"`
	Reasons     []string                                                        `json:"reasons"`
	Metrics     map[string]IntelligenceHandoffComputedResultsSellerMetricsValue `json:"metrics"`
	Counts      map[string]float64                                              `json:"counts"`
	Exclusions  map[string]float64                                              `json:"exclusions"`
	EvidenceIDs []string                                                        `json:"evidence_ids"`
	Context     IntelligenceHandoffComputedResultsSellerContext                 `json:"context"`
}

// UnmarshalJSON decodes IntelligenceHandoffComputedResultsSeller, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *IntelligenceHandoffComputedResultsSeller) UnmarshalJSON(data []byte) error {
	type plain IntelligenceHandoffComputedResultsSeller
	aux := struct {
		*plain
		Counts     lenientMap[float64] `json:"counts"`
		Exclusions lenientMap[float64] `json:"exclusions"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Counts.assign(&r.Counts)
	aux.Exclusions.assign(&r.Exclusions)
	return softTypeError(err)
}

// IntelligenceHandoffComputedResultsSellerStatus is generated from the OpenAPI spec. It is a
// string; the IntelligenceHandoffComputedResultsSellerStatus* constants list the documented
// values.
type IntelligenceHandoffComputedResultsSellerStatus = string

// Documented values of IntelligenceHandoffComputedResultsSellerStatus.
const (
	IntelligenceHandoffComputedResultsSellerStatusAvailable        IntelligenceHandoffComputedResultsSellerStatus = "available"
	IntelligenceHandoffComputedResultsSellerStatusPartial          IntelligenceHandoffComputedResultsSellerStatus = "partial"
	IntelligenceHandoffComputedResultsSellerStatusInsufficientData IntelligenceHandoffComputedResultsSellerStatus = "insufficient_data"
	IntelligenceHandoffComputedResultsSellerStatusUnavailable      IntelligenceHandoffComputedResultsSellerStatus = "unavailable"
	IntelligenceHandoffComputedResultsSellerStatusStale            IntelligenceHandoffComputedResultsSellerStatus = "stale"
	IntelligenceHandoffComputedResultsSellerStatusError            IntelligenceHandoffComputedResultsSellerStatus = "error"
)

// IntelligenceHandoffComputedResultsSellerMetricsValue is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsSellerMetricsValue struct {
	Value *IntelligenceHandoffComputedResultsSellerMetricsValueValue `json:"value,omitempty"`
	Unit  string                                                     `json:"unit"`
}

// IntelligenceHandoffComputedResultsSellerMetricsValueValue is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsSellerMetricsValueValue struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
}

// IntelligenceHandoffComputedResultsSellerContext is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsSellerContext struct {
	InputKind           IntelligenceHandoffComputedResultsSellerContextInputKind       `json:"input_kind"`
	RequestedScope      *IntelligenceHandoffComputedResultsSellerContextRequestedScope `json:"requested_scope,omitempty"`
	ActualScope         *IntelligenceHandoffComputedResultsSellerContextActualScope    `json:"actual_scope,omitempty"`
	Query               *IntelligenceHandoffComputedResultsSellerContextQuery          `json:"query,omitempty"`
	Period              json.RawMessage                                                `json:"period,omitempty"`
	MetricPeriods       map[string]json.RawMessage                                     `json:"metric_periods"`
	LatestObservationAt *string                                                        `json:"latest_observation_at,omitempty"`
	SourceAsOf          *string                                                        `json:"source_as_of,omitempty"`
	SourceVintage       *string                                                        `json:"source_vintage,omitempty"`
	HistoryComplete     *bool                                                          `json:"history_complete,omitempty"`
}

// IntelligenceHandoffComputedResultsSellerContextInputKind is generated from the OpenAPI spec. It
// is a string; the IntelligenceHandoffComputedResultsSellerContextInputKind* constants list the
// documented values.
type IntelligenceHandoffComputedResultsSellerContextInputKind = string

// Documented values of IntelligenceHandoffComputedResultsSellerContextInputKind.
const (
	IntelligenceHandoffComputedResultsSellerContextInputKindObserved        IntelligenceHandoffComputedResultsSellerContextInputKind = "observed"
	IntelligenceHandoffComputedResultsSellerContextInputKindUserAssumptions IntelligenceHandoffComputedResultsSellerContextInputKind = "user_assumptions"
	IntelligenceHandoffComputedResultsSellerContextInputKindUnavailable     IntelligenceHandoffComputedResultsSellerContextInputKind = "unavailable"
)

// IntelligenceHandoffComputedResultsSellerContextRequestedScope is generated from the OpenAPI
// spec.
type IntelligenceHandoffComputedResultsSellerContextRequestedScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceHandoffComputedResultsSellerContextActualScope is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsSellerContextActualScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceHandoffComputedResultsSellerContextQuery is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsSellerContextQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceHandoffComputedResultsOwner is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsOwner struct {
	Definition  string                                                         `json:"definition"`
	Status      IntelligenceHandoffComputedResultsOwnerStatus                  `json:"status"`
	Reasons     []string                                                       `json:"reasons"`
	Metrics     map[string]IntelligenceHandoffComputedResultsOwnerMetricsValue `json:"metrics"`
	Counts      map[string]float64                                             `json:"counts"`
	Exclusions  map[string]float64                                             `json:"exclusions"`
	EvidenceIDs []string                                                       `json:"evidence_ids"`
	Context     IntelligenceHandoffComputedResultsOwnerContext                 `json:"context"`
}

// UnmarshalJSON decodes IntelligenceHandoffComputedResultsOwner, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *IntelligenceHandoffComputedResultsOwner) UnmarshalJSON(data []byte) error {
	type plain IntelligenceHandoffComputedResultsOwner
	aux := struct {
		*plain
		Counts     lenientMap[float64] `json:"counts"`
		Exclusions lenientMap[float64] `json:"exclusions"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Counts.assign(&r.Counts)
	aux.Exclusions.assign(&r.Exclusions)
	return softTypeError(err)
}

// IntelligenceHandoffComputedResultsOwnerStatus is generated from the OpenAPI spec. It is a
// string; the IntelligenceHandoffComputedResultsOwnerStatus* constants list the documented values.
type IntelligenceHandoffComputedResultsOwnerStatus = string

// Documented values of IntelligenceHandoffComputedResultsOwnerStatus.
const (
	IntelligenceHandoffComputedResultsOwnerStatusAvailable        IntelligenceHandoffComputedResultsOwnerStatus = "available"
	IntelligenceHandoffComputedResultsOwnerStatusPartial          IntelligenceHandoffComputedResultsOwnerStatus = "partial"
	IntelligenceHandoffComputedResultsOwnerStatusInsufficientData IntelligenceHandoffComputedResultsOwnerStatus = "insufficient_data"
	IntelligenceHandoffComputedResultsOwnerStatusUnavailable      IntelligenceHandoffComputedResultsOwnerStatus = "unavailable"
	IntelligenceHandoffComputedResultsOwnerStatusStale            IntelligenceHandoffComputedResultsOwnerStatus = "stale"
	IntelligenceHandoffComputedResultsOwnerStatusError            IntelligenceHandoffComputedResultsOwnerStatus = "error"
)

// IntelligenceHandoffComputedResultsOwnerMetricsValue is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsOwnerMetricsValue struct {
	Value *IntelligenceHandoffComputedResultsOwnerMetricsValueValue `json:"value,omitempty"`
	Unit  string                                                    `json:"unit"`
}

// IntelligenceHandoffComputedResultsOwnerMetricsValueValue is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsOwnerMetricsValueValue struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
}

// IntelligenceHandoffComputedResultsOwnerContext is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsOwnerContext struct {
	InputKind           IntelligenceHandoffComputedResultsOwnerContextInputKind       `json:"input_kind"`
	RequestedScope      *IntelligenceHandoffComputedResultsOwnerContextRequestedScope `json:"requested_scope,omitempty"`
	ActualScope         *IntelligenceHandoffComputedResultsOwnerContextActualScope    `json:"actual_scope,omitempty"`
	Query               *IntelligenceHandoffComputedResultsOwnerContextQuery          `json:"query,omitempty"`
	Period              json.RawMessage                                               `json:"period,omitempty"`
	MetricPeriods       map[string]json.RawMessage                                    `json:"metric_periods"`
	LatestObservationAt *string                                                       `json:"latest_observation_at,omitempty"`
	SourceAsOf          *string                                                       `json:"source_as_of,omitempty"`
	SourceVintage       *string                                                       `json:"source_vintage,omitempty"`
	HistoryComplete     *bool                                                         `json:"history_complete,omitempty"`
}

// IntelligenceHandoffComputedResultsOwnerContextInputKind is generated from the OpenAPI spec. It
// is a string; the IntelligenceHandoffComputedResultsOwnerContextInputKind* constants list the
// documented values.
type IntelligenceHandoffComputedResultsOwnerContextInputKind = string

// Documented values of IntelligenceHandoffComputedResultsOwnerContextInputKind.
const (
	IntelligenceHandoffComputedResultsOwnerContextInputKindObserved        IntelligenceHandoffComputedResultsOwnerContextInputKind = "observed"
	IntelligenceHandoffComputedResultsOwnerContextInputKindUserAssumptions IntelligenceHandoffComputedResultsOwnerContextInputKind = "user_assumptions"
	IntelligenceHandoffComputedResultsOwnerContextInputKindUnavailable     IntelligenceHandoffComputedResultsOwnerContextInputKind = "unavailable"
)

// IntelligenceHandoffComputedResultsOwnerContextRequestedScope is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsOwnerContextRequestedScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceHandoffComputedResultsOwnerContextActualScope is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsOwnerContextActualScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceHandoffComputedResultsOwnerContextQuery is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsOwnerContextQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceHandoffComputedResultsRedevelopment is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsRedevelopment struct {
	Definition  string                                                                 `json:"definition"`
	Status      IntelligenceHandoffComputedResultsRedevelopmentStatus                  `json:"status"`
	Reasons     []string                                                               `json:"reasons"`
	Metrics     map[string]IntelligenceHandoffComputedResultsRedevelopmentMetricsValue `json:"metrics"`
	Counts      map[string]float64                                                     `json:"counts"`
	Exclusions  map[string]float64                                                     `json:"exclusions"`
	EvidenceIDs []string                                                               `json:"evidence_ids"`
	Context     IntelligenceHandoffComputedResultsRedevelopmentContext                 `json:"context"`
}

// UnmarshalJSON decodes IntelligenceHandoffComputedResultsRedevelopment, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *IntelligenceHandoffComputedResultsRedevelopment) UnmarshalJSON(data []byte) error {
	type plain IntelligenceHandoffComputedResultsRedevelopment
	aux := struct {
		*plain
		Counts     lenientMap[float64] `json:"counts"`
		Exclusions lenientMap[float64] `json:"exclusions"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Counts.assign(&r.Counts)
	aux.Exclusions.assign(&r.Exclusions)
	return softTypeError(err)
}

// IntelligenceHandoffComputedResultsRedevelopmentStatus is generated from the OpenAPI spec. It is
// a string; the IntelligenceHandoffComputedResultsRedevelopmentStatus* constants list the
// documented values.
type IntelligenceHandoffComputedResultsRedevelopmentStatus = string

// Documented values of IntelligenceHandoffComputedResultsRedevelopmentStatus.
const (
	IntelligenceHandoffComputedResultsRedevelopmentStatusAvailable        IntelligenceHandoffComputedResultsRedevelopmentStatus = "available"
	IntelligenceHandoffComputedResultsRedevelopmentStatusPartial          IntelligenceHandoffComputedResultsRedevelopmentStatus = "partial"
	IntelligenceHandoffComputedResultsRedevelopmentStatusInsufficientData IntelligenceHandoffComputedResultsRedevelopmentStatus = "insufficient_data"
	IntelligenceHandoffComputedResultsRedevelopmentStatusUnavailable      IntelligenceHandoffComputedResultsRedevelopmentStatus = "unavailable"
	IntelligenceHandoffComputedResultsRedevelopmentStatusStale            IntelligenceHandoffComputedResultsRedevelopmentStatus = "stale"
	IntelligenceHandoffComputedResultsRedevelopmentStatusError            IntelligenceHandoffComputedResultsRedevelopmentStatus = "error"
)

// IntelligenceHandoffComputedResultsRedevelopmentMetricsValue is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsRedevelopmentMetricsValue struct {
	Value *IntelligenceHandoffComputedResultsRedevelopmentMetricsValueValue `json:"value,omitempty"`
	Unit  string                                                            `json:"unit"`
}

// IntelligenceHandoffComputedResultsRedevelopmentMetricsValueValue is generated from the OpenAPI
// spec.
type IntelligenceHandoffComputedResultsRedevelopmentMetricsValueValue struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
}

// IntelligenceHandoffComputedResultsRedevelopmentContext is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsRedevelopmentContext struct {
	InputKind           IntelligenceHandoffComputedResultsRedevelopmentContextInputKind       `json:"input_kind"`
	RequestedScope      *IntelligenceHandoffComputedResultsRedevelopmentContextRequestedScope `json:"requested_scope,omitempty"`
	ActualScope         *IntelligenceHandoffComputedResultsRedevelopmentContextActualScope    `json:"actual_scope,omitempty"`
	Query               *IntelligenceHandoffComputedResultsRedevelopmentContextQuery          `json:"query,omitempty"`
	Period              json.RawMessage                                                       `json:"period,omitempty"`
	MetricPeriods       map[string]json.RawMessage                                            `json:"metric_periods"`
	LatestObservationAt *string                                                               `json:"latest_observation_at,omitempty"`
	SourceAsOf          *string                                                               `json:"source_as_of,omitempty"`
	SourceVintage       *string                                                               `json:"source_vintage,omitempty"`
	HistoryComplete     *bool                                                                 `json:"history_complete,omitempty"`
}

// IntelligenceHandoffComputedResultsRedevelopmentContextInputKind is generated from the OpenAPI
// spec. It is a string; the IntelligenceHandoffComputedResultsRedevelopmentContextInputKind*
// constants list the documented values.
type IntelligenceHandoffComputedResultsRedevelopmentContextInputKind = string

// Documented values of IntelligenceHandoffComputedResultsRedevelopmentContextInputKind.
const (
	IntelligenceHandoffComputedResultsRedevelopmentContextInputKindObserved        IntelligenceHandoffComputedResultsRedevelopmentContextInputKind = "observed"
	IntelligenceHandoffComputedResultsRedevelopmentContextInputKindUserAssumptions IntelligenceHandoffComputedResultsRedevelopmentContextInputKind = "user_assumptions"
	IntelligenceHandoffComputedResultsRedevelopmentContextInputKindUnavailable     IntelligenceHandoffComputedResultsRedevelopmentContextInputKind = "unavailable"
)

// IntelligenceHandoffComputedResultsRedevelopmentContextRequestedScope is generated from the
// OpenAPI spec.
type IntelligenceHandoffComputedResultsRedevelopmentContextRequestedScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceHandoffComputedResultsRedevelopmentContextActualScope is generated from the OpenAPI
// spec.
type IntelligenceHandoffComputedResultsRedevelopmentContextActualScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceHandoffComputedResultsRedevelopmentContextQuery is generated from the OpenAPI spec.
type IntelligenceHandoffComputedResultsRedevelopmentContextQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceHandoffUserAssumptions is generated from the OpenAPI spec.
type IntelligenceHandoffUserAssumptions struct {
	ID               string                                   `json:"id"`
	RunID            string                                   `json:"run_id"`
	ParentRevisionID *string                                  `json:"parent_revision_id,omitempty"`
	Label            IntelligenceHandoffUserAssumptionsLabel  `json:"label"`
	CreatedAt        string                                   `json:"created_at"`
	Assumptions      map[string]any                           `json:"assumptions"`
	Result           IntelligenceHandoffUserAssumptionsResult `json:"result"`
}

// IntelligenceHandoffUserAssumptionsLabel is generated from the OpenAPI spec. It is a string; the
// IntelligenceHandoffUserAssumptionsLabel* constants list the documented values.
type IntelligenceHandoffUserAssumptionsLabel = string

// Documented values of IntelligenceHandoffUserAssumptionsLabel.
const (
	IntelligenceHandoffUserAssumptionsLabelBase     IntelligenceHandoffUserAssumptionsLabel = "base"
	IntelligenceHandoffUserAssumptionsLabelDownside IntelligenceHandoffUserAssumptionsLabel = "downside"
	IntelligenceHandoffUserAssumptionsLabelUpside   IntelligenceHandoffUserAssumptionsLabel = "upside"
)

// IntelligenceHandoffUserAssumptionsResult is generated from the OpenAPI spec.
type IntelligenceHandoffUserAssumptionsResult struct {
	Definition  string                                                          `json:"definition"`
	Status      IntelligenceHandoffUserAssumptionsResultStatus                  `json:"status"`
	Reasons     []string                                                        `json:"reasons"`
	Metrics     map[string]IntelligenceHandoffUserAssumptionsResultMetricsValue `json:"metrics"`
	Counts      map[string]float64                                              `json:"counts"`
	Exclusions  map[string]float64                                              `json:"exclusions"`
	EvidenceIDs []string                                                        `json:"evidence_ids"`
	Context     IntelligenceHandoffUserAssumptionsResultContext                 `json:"context"`
}

// UnmarshalJSON decodes IntelligenceHandoffUserAssumptionsResult, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *IntelligenceHandoffUserAssumptionsResult) UnmarshalJSON(data []byte) error {
	type plain IntelligenceHandoffUserAssumptionsResult
	aux := struct {
		*plain
		Counts     lenientMap[float64] `json:"counts"`
		Exclusions lenientMap[float64] `json:"exclusions"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Counts.assign(&r.Counts)
	aux.Exclusions.assign(&r.Exclusions)
	return softTypeError(err)
}

// IntelligenceHandoffUserAssumptionsResultStatus is generated from the OpenAPI spec. It is a
// string; the IntelligenceHandoffUserAssumptionsResultStatus* constants list the documented
// values.
type IntelligenceHandoffUserAssumptionsResultStatus = string

// Documented values of IntelligenceHandoffUserAssumptionsResultStatus.
const (
	IntelligenceHandoffUserAssumptionsResultStatusAvailable        IntelligenceHandoffUserAssumptionsResultStatus = "available"
	IntelligenceHandoffUserAssumptionsResultStatusPartial          IntelligenceHandoffUserAssumptionsResultStatus = "partial"
	IntelligenceHandoffUserAssumptionsResultStatusInsufficientData IntelligenceHandoffUserAssumptionsResultStatus = "insufficient_data"
	IntelligenceHandoffUserAssumptionsResultStatusUnavailable      IntelligenceHandoffUserAssumptionsResultStatus = "unavailable"
	IntelligenceHandoffUserAssumptionsResultStatusStale            IntelligenceHandoffUserAssumptionsResultStatus = "stale"
	IntelligenceHandoffUserAssumptionsResultStatusError            IntelligenceHandoffUserAssumptionsResultStatus = "error"
)

// IntelligenceHandoffUserAssumptionsResultMetricsValue is generated from the OpenAPI spec.
type IntelligenceHandoffUserAssumptionsResultMetricsValue struct {
	Value *IntelligenceHandoffUserAssumptionsResultMetricsValueValue `json:"value,omitempty"`
	Unit  string                                                     `json:"unit"`
}

// IntelligenceHandoffUserAssumptionsResultMetricsValueValue is generated from the OpenAPI spec.
type IntelligenceHandoffUserAssumptionsResultMetricsValueValue struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
}

// IntelligenceHandoffUserAssumptionsResultContext is generated from the OpenAPI spec.
type IntelligenceHandoffUserAssumptionsResultContext struct {
	InputKind           IntelligenceHandoffUserAssumptionsResultContextInputKind       `json:"input_kind"`
	RequestedScope      *IntelligenceHandoffUserAssumptionsResultContextRequestedScope `json:"requested_scope,omitempty"`
	ActualScope         *IntelligenceHandoffUserAssumptionsResultContextActualScope    `json:"actual_scope,omitempty"`
	Query               *IntelligenceHandoffUserAssumptionsResultContextQuery          `json:"query,omitempty"`
	Period              json.RawMessage                                                `json:"period,omitempty"`
	MetricPeriods       map[string]json.RawMessage                                     `json:"metric_periods"`
	LatestObservationAt *string                                                        `json:"latest_observation_at,omitempty"`
	SourceAsOf          *string                                                        `json:"source_as_of,omitempty"`
	SourceVintage       *string                                                        `json:"source_vintage,omitempty"`
	HistoryComplete     *bool                                                          `json:"history_complete,omitempty"`
}

// IntelligenceHandoffUserAssumptionsResultContextInputKind is generated from the OpenAPI spec. It
// is a string; the IntelligenceHandoffUserAssumptionsResultContextInputKind* constants list the
// documented values.
type IntelligenceHandoffUserAssumptionsResultContextInputKind = string

// Documented values of IntelligenceHandoffUserAssumptionsResultContextInputKind.
const (
	IntelligenceHandoffUserAssumptionsResultContextInputKindObserved        IntelligenceHandoffUserAssumptionsResultContextInputKind = "observed"
	IntelligenceHandoffUserAssumptionsResultContextInputKindUserAssumptions IntelligenceHandoffUserAssumptionsResultContextInputKind = "user_assumptions"
	IntelligenceHandoffUserAssumptionsResultContextInputKindUnavailable     IntelligenceHandoffUserAssumptionsResultContextInputKind = "unavailable"
)

// IntelligenceHandoffUserAssumptionsResultContextRequestedScope is generated from the OpenAPI
// spec.
type IntelligenceHandoffUserAssumptionsResultContextRequestedScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceHandoffUserAssumptionsResultContextActualScope is generated from the OpenAPI spec.
type IntelligenceHandoffUserAssumptionsResultContextActualScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceHandoffUserAssumptionsResultContextQuery is generated from the OpenAPI spec.
type IntelligenceHandoffUserAssumptionsResultContextQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceScenarioInput: All five unique cost buckets and a matching currency are required;
// explicitly enter zero. No inferred defaults. Only fixed-dollar profit and purchase-independent
// carry.
type IntelligenceScenarioInput struct {
	RunID            string                               `json:"run_id"`
	Label            IntelligenceScenarioInputLabel       `json:"label"`
	ParentRevisionID *string                              `json:"parent_revision_id,omitempty"`
	Assumptions      IntelligenceScenarioInputAssumptions `json:"assumptions"`
}

// IntelligenceScenarioInputLabel is generated from the OpenAPI spec. It is a string; the
// IntelligenceScenarioInputLabel* constants list the documented values.
type IntelligenceScenarioInputLabel = string

// Documented values of IntelligenceScenarioInputLabel.
const (
	IntelligenceScenarioInputLabelBase     IntelligenceScenarioInputLabel = "base"
	IntelligenceScenarioInputLabelDownside IntelligenceScenarioInputLabel = "downside"
	IntelligenceScenarioInputLabelUpside   IntelligenceScenarioInputLabel = "upside"
)

// IntelligenceScenarioInputAssumptions is generated from the OpenAPI spec.
type IntelligenceScenarioInputAssumptions struct {
	Currency              string                                      `json:"currency"`
	GrossCompletedSale    string                                      `json:"gross_completed_sale"`
	SellingCosts          string                                      `json:"selling_costs"`
	Costs                 []IntelligenceScenarioInputAssumptionsCosts `json:"costs"`
	RequiredProfitDollars string                                      `json:"required_profit_dollars"`
	FixedAcquisitionCosts string                                      `json:"fixed_acquisition_costs"`
	AcquisitionCostRate   string                                      `json:"acquisition_cost_rate"`
	ProfitMode            string                                      `json:"profit_mode"`
	CarryMode             string                                      `json:"carry_mode"`
	InputSource           string                                      `json:"input_source"`
}

// IntelligenceScenarioInputAssumptionsCosts is generated from the OpenAPI spec.
type IntelligenceScenarioInputAssumptionsCosts struct {
	Bucket   IntelligenceScenarioInputAssumptionsCostsBucket `json:"bucket"`
	Amount   string                                          `json:"amount"`
	Currency string                                          `json:"currency"`
}

// IntelligenceScenarioInputAssumptionsCostsBucket is generated from the OpenAPI spec. It is a
// string; the IntelligenceScenarioInputAssumptionsCostsBucket* constants list the documented
// values.
type IntelligenceScenarioInputAssumptionsCostsBucket = string

// Documented values of IntelligenceScenarioInputAssumptionsCostsBucket.
const (
	IntelligenceScenarioInputAssumptionsCostsBucketHard         IntelligenceScenarioInputAssumptionsCostsBucket = "hard"
	IntelligenceScenarioInputAssumptionsCostsBucketSoft         IntelligenceScenarioInputAssumptionsCostsBucket = "soft"
	IntelligenceScenarioInputAssumptionsCostsBucketContingency  IntelligenceScenarioInputAssumptionsCostsBucket = "contingency"
	IntelligenceScenarioInputAssumptionsCostsBucketCarry        IntelligenceScenarioInputAssumptionsCostsBucket = "carry"
	IntelligenceScenarioInputAssumptionsCostsBucketOtherNonland IntelligenceScenarioInputAssumptionsCostsBucket = "other_nonland"
)

// ZillowMetric: Regional provider metric with its actual geography, variant, period, accepted
// snapshot and source attribution. Never a parcel value, achieved rent or automatic scenario
// input.
type ZillowMetric struct {
	Rights           *ZillowMetricRights           `json:"rights,omitempty"`
	Metric           ZillowMetricMetric            `json:"metric"`
	Status           ZillowMetricStatus            `json:"status"`
	Reason           *string                       `json:"reason,omitempty"`
	Definition       string                        `json:"definition"`
	Unit             ZillowMetricUnit              `json:"unit"`
	Value            *float64                      `json:"value,omitempty"`
	Period           *string                       `json:"period,omitempty"`
	Variant          *ZillowMetricVariant          `json:"variant,omitempty"`
	Geography        *ZillowMetricGeography        `json:"geography,omitempty"`
	Mapping          *ZillowMetricMapping          `json:"mapping,omitempty"`
	Snapshot         *ZillowMetricSnapshot         `json:"snapshot,omitempty"`
	AnnualChange     *ZillowMetricAnnualChange     `json:"annualChange,omitempty"`
	MonthlyChange    *ZillowMetricMonthlyChange    `json:"monthlyChange,omitempty"`
	RentAcceleration *ZillowMetricRentAcceleration `json:"rentAcceleration,omitempty"`
	Points           []ZillowMetricPoints          `json:"points"`
	SourceURL        string                        `json:"sourceUrl"`
	Attribution      string                        `json:"attribution"`
}

// UnmarshalJSON decodes ZillowMetric, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowMetric) UnmarshalJSON(data []byte) error {
	type plain ZillowMetric
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowMetricRights is generated from the OpenAPI spec.
type ZillowMetricRights struct {
	Version     string  `json:"version"`
	EvidenceURL *string `json:"evidenceUrl,omitempty"`
	ExpiresAt   *string `json:"expiresAt,omitempty"`
}

// ZillowMetricMetric is generated from the OpenAPI spec. It is a string; the ZillowMetricMetric*
// constants list the documented values.
type ZillowMetricMetric = string

// Documented values of ZillowMetricMetric.
const (
	ZillowMetricMetricZori                ZillowMetricMetric = "zori"
	ZillowMetricMetricZhvi                ZillowMetricMetric = "zhvi"
	ZillowMetricMetricInventory           ZillowMetricMetric = "inventory"
	ZillowMetricMetricPriceCutShare       ZillowMetricMetric = "price_cut_share"
	ZillowMetricMetricMedianDaysToPending ZillowMetricMetric = "median_days_to_pending"
)

// ZillowMetricStatus is generated from the OpenAPI spec. It is a string; the ZillowMetricStatus*
// constants list the documented values.
type ZillowMetricStatus = string

// Documented values of ZillowMetricStatus.
const (
	ZillowMetricStatusAvailable   ZillowMetricStatus = "available"
	ZillowMetricStatusUnavailable ZillowMetricStatus = "unavailable"
)

// ZillowMetricUnit is generated from the OpenAPI spec. It is a string; the ZillowMetricUnit*
// constants list the documented values.
type ZillowMetricUnit = string

// Documented values of ZillowMetricUnit.
const (
	ZillowMetricUnitUsd         ZillowMetricUnit = "usd"
	ZillowMetricUnitUsdPerMonth ZillowMetricUnit = "usd_per_month"
	ZillowMetricUnitCount       ZillowMetricUnit = "count"
	ZillowMetricUnitFraction    ZillowMetricUnit = "fraction"
	ZillowMetricUnitDays        ZillowMetricUnit = "days"
)

// ZillowMetricVariant is generated from the OpenAPI spec.
type ZillowMetricVariant struct {
	DatasetKey         string                                `json:"datasetKey"`
	RegistryVersion    int64                                 `json:"registryVersion"`
	Universe           string                                `json:"universe"`
	Frequency          string                                `json:"frequency"`
	Smoothing          string                                `json:"smoothing"`
	SeasonalAdjustment ZillowMetricVariantSeasonalAdjustment `json:"seasonalAdjustment"`
}

// UnmarshalJSON decodes ZillowMetricVariant, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowMetricVariant) UnmarshalJSON(data []byte) error {
	type plain ZillowMetricVariant
	aux := struct {
		*plain
		RegistryVersion lenientNumber[int64] `json:"registryVersion"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.RegistryVersion.assign(&r.RegistryVersion)
	return softTypeError(err)
}

// ZillowMetricVariantSeasonalAdjustment is generated from the OpenAPI spec. It is a string; the
// ZillowMetricVariantSeasonalAdjustment* constants list the documented values.
type ZillowMetricVariantSeasonalAdjustment = string

// Documented values of ZillowMetricVariantSeasonalAdjustment.
const (
	ZillowMetricVariantSeasonalAdjustmentSa        ZillowMetricVariantSeasonalAdjustment = "sa"
	ZillowMetricVariantSeasonalAdjustmentNotStated ZillowMetricVariantSeasonalAdjustment = "not_stated"
)

// ZillowMetricGeography is generated from the OpenAPI spec.
type ZillowMetricGeography struct {
	ProviderID string                    `json:"providerId"`
	Name       string                    `json:"name"`
	Type       ZillowMetricGeographyType `json:"type"`
}

// ZillowMetricGeographyType is generated from the OpenAPI spec. It is a string; the
// ZillowMetricGeographyType* constants list the documented values.
type ZillowMetricGeographyType = string

// Documented values of ZillowMetricGeographyType.
const (
	ZillowMetricGeographyTypeCountry ZillowMetricGeographyType = "country"
	ZillowMetricGeographyTypeMsa     ZillowMetricGeographyType = "msa"
	ZillowMetricGeographyTypeCounty  ZillowMetricGeographyType = "county"
	ZillowMetricGeographyTypeZip     ZillowMetricGeographyType = "zip"
)

// ZillowMetricMapping is generated from the OpenAPI spec.
type ZillowMetricMapping struct {
	Method         ZillowMetricMappingMethod `json:"method"`
	Version        string                    `json:"version"`
	Source         string                    `json:"source"`
	FallbackReason *string                   `json:"fallbackReason,omitempty"`
}

// ZillowMetricMappingMethod is generated from the OpenAPI spec. It is a string; the
// ZillowMetricMappingMethod* constants list the documented values.
type ZillowMetricMappingMethod = string

// Documented values of ZillowMetricMappingMethod.
const (
	ZillowMetricMappingMethodPostalZip              ZillowMetricMappingMethod = "postal_zip"
	ZillowMetricMappingMethodCountyFIPS             ZillowMetricMappingMethod = "county_fips"
	ZillowMetricMappingMethodVerifiedCrosswalk      ZillowMetricMappingMethod = "verified_crosswalk"
	ZillowMetricMappingMethodExplicitProviderRegion ZillowMetricMappingMethod = "explicit_provider_region"
)

// ZillowMetricSnapshot is generated from the OpenAPI spec.
type ZillowMetricSnapshot struct {
	ID           string `json:"id"`
	Sha256       string `json:"sha256"`
	RetrievedAt  string `json:"retrievedAt"`
	AcceptedAt   string `json:"acceptedAt"`
	LatestPeriod string `json:"latestPeriod"`
	Stale        bool   `json:"stale"`
}

// ZillowMetricAnnualChange is generated from the OpenAPI spec.
type ZillowMetricAnnualChange struct {
	Value  *float64                     `json:"value,omitempty"`
	Unit   ZillowMetricAnnualChangeUnit `json:"unit"`
	Reason *string                      `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowMetricAnnualChange, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowMetricAnnualChange) UnmarshalJSON(data []byte) error {
	type plain ZillowMetricAnnualChange
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowMetricAnnualChangeUnit is generated from the OpenAPI spec. It is a string; the
// ZillowMetricAnnualChangeUnit* constants list the documented values.
type ZillowMetricAnnualChangeUnit = string

// Documented values of ZillowMetricAnnualChangeUnit.
const (
	ZillowMetricAnnualChangeUnitPercent          ZillowMetricAnnualChangeUnit = "percent"
	ZillowMetricAnnualChangeUnitPercentagePoints ZillowMetricAnnualChangeUnit = "percentage_points"
	ZillowMetricAnnualChangeUnitDays             ZillowMetricAnnualChangeUnit = "days"
)

// ZillowMetricMonthlyChange is generated from the OpenAPI spec.
type ZillowMetricMonthlyChange struct {
	Value  *float64                      `json:"value,omitempty"`
	Unit   ZillowMetricMonthlyChangeUnit `json:"unit"`
	Reason *string                       `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowMetricMonthlyChange, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowMetricMonthlyChange) UnmarshalJSON(data []byte) error {
	type plain ZillowMetricMonthlyChange
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowMetricMonthlyChangeUnit is generated from the OpenAPI spec. It is a string; the
// ZillowMetricMonthlyChangeUnit* constants list the documented values.
type ZillowMetricMonthlyChangeUnit = string

// Documented values of ZillowMetricMonthlyChangeUnit.
const (
	ZillowMetricMonthlyChangeUnitPercent          ZillowMetricMonthlyChangeUnit = "percent"
	ZillowMetricMonthlyChangeUnitPercentagePoints ZillowMetricMonthlyChangeUnit = "percentage_points"
	ZillowMetricMonthlyChangeUnitDays             ZillowMetricMonthlyChangeUnit = "days"
)

// ZillowMetricRentAcceleration is generated from the OpenAPI spec.
type ZillowMetricRentAcceleration struct {
	Value  *float64                         `json:"value,omitempty"`
	Unit   ZillowMetricRentAccelerationUnit `json:"unit"`
	Reason *string                          `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowMetricRentAcceleration, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowMetricRentAcceleration) UnmarshalJSON(data []byte) error {
	type plain ZillowMetricRentAcceleration
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowMetricRentAccelerationUnit is generated from the OpenAPI spec. It is a string; the
// ZillowMetricRentAccelerationUnit* constants list the documented values.
type ZillowMetricRentAccelerationUnit = string

// Documented values of ZillowMetricRentAccelerationUnit.
const (
	ZillowMetricRentAccelerationUnitPercent          ZillowMetricRentAccelerationUnit = "percent"
	ZillowMetricRentAccelerationUnitPercentagePoints ZillowMetricRentAccelerationUnit = "percentage_points"
	ZillowMetricRentAccelerationUnitDays             ZillowMetricRentAccelerationUnit = "days"
)

// ZillowMetricPoints is generated from the OpenAPI spec.
type ZillowMetricPoints struct {
	Period string   `json:"period"`
	Value  *float64 `json:"value,omitempty"`
}

// UnmarshalJSON decodes ZillowMetricPoints, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowMetricPoints) UnmarshalJSON(data []byte) error {
	type plain ZillowMetricPoints
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowContext: Each metric carries its own actual geography and missing reason. canonical_id is
// omitted when source rights prevent parcel resolution.
type ZillowContext struct {
	SchemaVersion      string                          `json:"schemaVersion"`
	Status             ZillowContextStatus             `json:"status"`
	CanonicalID        *string                         `json:"canonical_id,omitempty"`
	RequestedGeography ZillowContextRequestedGeography `json:"requestedGeography"`
	Metrics            []ZillowContextMetrics          `json:"metrics"`
	Reason             *string                         `json:"reason,omitempty"`
	Rights             ZillowContextRights             `json:"rights"`
}

// ZillowContextStatus is generated from the OpenAPI spec. It is a string; the ZillowContextStatus*
// constants list the documented values.
type ZillowContextStatus = string

// Documented values of ZillowContextStatus.
const (
	ZillowContextStatusAvailable   ZillowContextStatus = "available"
	ZillowContextStatusPartial     ZillowContextStatus = "partial"
	ZillowContextStatusUnavailable ZillowContextStatus = "unavailable"
)

// ZillowContextRequestedGeography is generated from the OpenAPI spec.
type ZillowContextRequestedGeography struct {
	Zip5       *string `json:"zip5,omitempty"`
	CountyFIPS *string `json:"countyFips,omitempty"`
	State      *string `json:"state,omitempty"`
	CBSA       *string `json:"cbsa,omitempty"`
}

// ZillowContextMetrics is generated from the OpenAPI spec.
type ZillowContextMetrics struct {
	Rights           *ZillowContextMetricsRights           `json:"rights,omitempty"`
	Metric           ZillowContextMetricsMetric            `json:"metric"`
	Status           ZillowContextMetricsStatus            `json:"status"`
	Reason           *string                               `json:"reason,omitempty"`
	Definition       string                                `json:"definition"`
	Unit             ZillowContextMetricsUnit              `json:"unit"`
	Value            *float64                              `json:"value,omitempty"`
	Period           *string                               `json:"period,omitempty"`
	Variant          *ZillowContextMetricsVariant          `json:"variant,omitempty"`
	Geography        *ZillowContextMetricsGeography        `json:"geography,omitempty"`
	Mapping          *ZillowContextMetricsMapping          `json:"mapping,omitempty"`
	Snapshot         *ZillowContextMetricsSnapshot         `json:"snapshot,omitempty"`
	AnnualChange     *ZillowContextMetricsAnnualChange     `json:"annualChange,omitempty"`
	MonthlyChange    *ZillowContextMetricsMonthlyChange    `json:"monthlyChange,omitempty"`
	RentAcceleration *ZillowContextMetricsRentAcceleration `json:"rentAcceleration,omitempty"`
	Points           []ZillowContextMetricsPoints          `json:"points"`
	SourceURL        string                                `json:"sourceUrl"`
	Attribution      string                                `json:"attribution"`
}

// UnmarshalJSON decodes ZillowContextMetrics, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowContextMetrics) UnmarshalJSON(data []byte) error {
	type plain ZillowContextMetrics
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowContextMetricsRights is generated from the OpenAPI spec.
type ZillowContextMetricsRights struct {
	Version     string  `json:"version"`
	EvidenceURL *string `json:"evidenceUrl,omitempty"`
	ExpiresAt   *string `json:"expiresAt,omitempty"`
}

// ZillowContextMetricsMetric is generated from the OpenAPI spec. It is a string; the
// ZillowContextMetricsMetric* constants list the documented values.
type ZillowContextMetricsMetric = string

// Documented values of ZillowContextMetricsMetric.
const (
	ZillowContextMetricsMetricZori                ZillowContextMetricsMetric = "zori"
	ZillowContextMetricsMetricZhvi                ZillowContextMetricsMetric = "zhvi"
	ZillowContextMetricsMetricInventory           ZillowContextMetricsMetric = "inventory"
	ZillowContextMetricsMetricPriceCutShare       ZillowContextMetricsMetric = "price_cut_share"
	ZillowContextMetricsMetricMedianDaysToPending ZillowContextMetricsMetric = "median_days_to_pending"
)

// ZillowContextMetricsStatus is generated from the OpenAPI spec. It is a string; the
// ZillowContextMetricsStatus* constants list the documented values.
type ZillowContextMetricsStatus = string

// Documented values of ZillowContextMetricsStatus.
const (
	ZillowContextMetricsStatusAvailable   ZillowContextMetricsStatus = "available"
	ZillowContextMetricsStatusUnavailable ZillowContextMetricsStatus = "unavailable"
)

// ZillowContextMetricsUnit is generated from the OpenAPI spec. It is a string; the
// ZillowContextMetricsUnit* constants list the documented values.
type ZillowContextMetricsUnit = string

// Documented values of ZillowContextMetricsUnit.
const (
	ZillowContextMetricsUnitUsd         ZillowContextMetricsUnit = "usd"
	ZillowContextMetricsUnitUsdPerMonth ZillowContextMetricsUnit = "usd_per_month"
	ZillowContextMetricsUnitCount       ZillowContextMetricsUnit = "count"
	ZillowContextMetricsUnitFraction    ZillowContextMetricsUnit = "fraction"
	ZillowContextMetricsUnitDays        ZillowContextMetricsUnit = "days"
)

// ZillowContextMetricsVariant is generated from the OpenAPI spec.
type ZillowContextMetricsVariant struct {
	DatasetKey         string                                        `json:"datasetKey"`
	RegistryVersion    int64                                         `json:"registryVersion"`
	Universe           string                                        `json:"universe"`
	Frequency          string                                        `json:"frequency"`
	Smoothing          string                                        `json:"smoothing"`
	SeasonalAdjustment ZillowContextMetricsVariantSeasonalAdjustment `json:"seasonalAdjustment"`
}

// UnmarshalJSON decodes ZillowContextMetricsVariant, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowContextMetricsVariant) UnmarshalJSON(data []byte) error {
	type plain ZillowContextMetricsVariant
	aux := struct {
		*plain
		RegistryVersion lenientNumber[int64] `json:"registryVersion"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.RegistryVersion.assign(&r.RegistryVersion)
	return softTypeError(err)
}

// ZillowContextMetricsVariantSeasonalAdjustment is generated from the OpenAPI spec. It is a
// string; the ZillowContextMetricsVariantSeasonalAdjustment* constants list the documented values.
type ZillowContextMetricsVariantSeasonalAdjustment = string

// Documented values of ZillowContextMetricsVariantSeasonalAdjustment.
const (
	ZillowContextMetricsVariantSeasonalAdjustmentSa        ZillowContextMetricsVariantSeasonalAdjustment = "sa"
	ZillowContextMetricsVariantSeasonalAdjustmentNotStated ZillowContextMetricsVariantSeasonalAdjustment = "not_stated"
)

// ZillowContextMetricsGeography is generated from the OpenAPI spec.
type ZillowContextMetricsGeography struct {
	ProviderID string                            `json:"providerId"`
	Name       string                            `json:"name"`
	Type       ZillowContextMetricsGeographyType `json:"type"`
}

// ZillowContextMetricsGeographyType is generated from the OpenAPI spec. It is a string; the
// ZillowContextMetricsGeographyType* constants list the documented values.
type ZillowContextMetricsGeographyType = string

// Documented values of ZillowContextMetricsGeographyType.
const (
	ZillowContextMetricsGeographyTypeCountry ZillowContextMetricsGeographyType = "country"
	ZillowContextMetricsGeographyTypeMsa     ZillowContextMetricsGeographyType = "msa"
	ZillowContextMetricsGeographyTypeCounty  ZillowContextMetricsGeographyType = "county"
	ZillowContextMetricsGeographyTypeZip     ZillowContextMetricsGeographyType = "zip"
)

// ZillowContextMetricsMapping is generated from the OpenAPI spec.
type ZillowContextMetricsMapping struct {
	Method         ZillowContextMetricsMappingMethod `json:"method"`
	Version        string                            `json:"version"`
	Source         string                            `json:"source"`
	FallbackReason *string                           `json:"fallbackReason,omitempty"`
}

// ZillowContextMetricsMappingMethod is generated from the OpenAPI spec. It is a string; the
// ZillowContextMetricsMappingMethod* constants list the documented values.
type ZillowContextMetricsMappingMethod = string

// Documented values of ZillowContextMetricsMappingMethod.
const (
	ZillowContextMetricsMappingMethodPostalZip              ZillowContextMetricsMappingMethod = "postal_zip"
	ZillowContextMetricsMappingMethodCountyFIPS             ZillowContextMetricsMappingMethod = "county_fips"
	ZillowContextMetricsMappingMethodVerifiedCrosswalk      ZillowContextMetricsMappingMethod = "verified_crosswalk"
	ZillowContextMetricsMappingMethodExplicitProviderRegion ZillowContextMetricsMappingMethod = "explicit_provider_region"
)

// ZillowContextMetricsSnapshot is generated from the OpenAPI spec.
type ZillowContextMetricsSnapshot struct {
	ID           string `json:"id"`
	Sha256       string `json:"sha256"`
	RetrievedAt  string `json:"retrievedAt"`
	AcceptedAt   string `json:"acceptedAt"`
	LatestPeriod string `json:"latestPeriod"`
	Stale        bool   `json:"stale"`
}

// ZillowContextMetricsAnnualChange is generated from the OpenAPI spec.
type ZillowContextMetricsAnnualChange struct {
	Value  *float64                             `json:"value,omitempty"`
	Unit   ZillowContextMetricsAnnualChangeUnit `json:"unit"`
	Reason *string                              `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowContextMetricsAnnualChange, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowContextMetricsAnnualChange) UnmarshalJSON(data []byte) error {
	type plain ZillowContextMetricsAnnualChange
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowContextMetricsAnnualChangeUnit is generated from the OpenAPI spec. It is a string; the
// ZillowContextMetricsAnnualChangeUnit* constants list the documented values.
type ZillowContextMetricsAnnualChangeUnit = string

// Documented values of ZillowContextMetricsAnnualChangeUnit.
const (
	ZillowContextMetricsAnnualChangeUnitPercent          ZillowContextMetricsAnnualChangeUnit = "percent"
	ZillowContextMetricsAnnualChangeUnitPercentagePoints ZillowContextMetricsAnnualChangeUnit = "percentage_points"
	ZillowContextMetricsAnnualChangeUnitDays             ZillowContextMetricsAnnualChangeUnit = "days"
)

// ZillowContextMetricsMonthlyChange is generated from the OpenAPI spec.
type ZillowContextMetricsMonthlyChange struct {
	Value  *float64                              `json:"value,omitempty"`
	Unit   ZillowContextMetricsMonthlyChangeUnit `json:"unit"`
	Reason *string                               `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowContextMetricsMonthlyChange, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowContextMetricsMonthlyChange) UnmarshalJSON(data []byte) error {
	type plain ZillowContextMetricsMonthlyChange
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowContextMetricsMonthlyChangeUnit is generated from the OpenAPI spec. It is a string; the
// ZillowContextMetricsMonthlyChangeUnit* constants list the documented values.
type ZillowContextMetricsMonthlyChangeUnit = string

// Documented values of ZillowContextMetricsMonthlyChangeUnit.
const (
	ZillowContextMetricsMonthlyChangeUnitPercent          ZillowContextMetricsMonthlyChangeUnit = "percent"
	ZillowContextMetricsMonthlyChangeUnitPercentagePoints ZillowContextMetricsMonthlyChangeUnit = "percentage_points"
	ZillowContextMetricsMonthlyChangeUnitDays             ZillowContextMetricsMonthlyChangeUnit = "days"
)

// ZillowContextMetricsRentAcceleration is generated from the OpenAPI spec.
type ZillowContextMetricsRentAcceleration struct {
	Value  *float64                                 `json:"value,omitempty"`
	Unit   ZillowContextMetricsRentAccelerationUnit `json:"unit"`
	Reason *string                                  `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowContextMetricsRentAcceleration, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowContextMetricsRentAcceleration) UnmarshalJSON(data []byte) error {
	type plain ZillowContextMetricsRentAcceleration
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowContextMetricsRentAccelerationUnit is generated from the OpenAPI spec. It is a string; the
// ZillowContextMetricsRentAccelerationUnit* constants list the documented values.
type ZillowContextMetricsRentAccelerationUnit = string

// Documented values of ZillowContextMetricsRentAccelerationUnit.
const (
	ZillowContextMetricsRentAccelerationUnitPercent          ZillowContextMetricsRentAccelerationUnit = "percent"
	ZillowContextMetricsRentAccelerationUnitPercentagePoints ZillowContextMetricsRentAccelerationUnit = "percentage_points"
	ZillowContextMetricsRentAccelerationUnitDays             ZillowContextMetricsRentAccelerationUnit = "days"
)

// ZillowContextMetricsPoints is generated from the OpenAPI spec.
type ZillowContextMetricsPoints struct {
	Period string   `json:"period"`
	Value  *float64 `json:"value,omitempty"`
}

// UnmarshalJSON decodes ZillowContextMetricsPoints, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowContextMetricsPoints) UnmarshalJSON(data []byte) error {
	type plain ZillowContextMetricsPoints
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowContextRights is generated from the OpenAPI spec.
type ZillowContextRights struct {
	Status      ZillowContextRightsStatus `json:"status"`
	Use         ZillowContextRightsUse    `json:"use"`
	EvidenceURL *string                   `json:"evidenceUrl,omitempty"`
}

// ZillowContextRightsStatus is generated from the OpenAPI spec. It is a string; the
// ZillowContextRightsStatus* constants list the documented values.
type ZillowContextRightsStatus = string

// Documented values of ZillowContextRightsStatus.
const (
	ZillowContextRightsStatusApproved ZillowContextRightsStatus = "approved"
	ZillowContextRightsStatusUnknown  ZillowContextRightsStatus = "unknown"
	ZillowContextRightsStatusDenied   ZillowContextRightsStatus = "denied"
)

// ZillowContextRightsUse is generated from the OpenAPI spec. It is a string; the
// ZillowContextRightsUse* constants list the documented values.
type ZillowContextRightsUse = string

// Documented values of ZillowContextRightsUse.
const (
	ZillowContextRightsUseDisplay ZillowContextRightsUse = "display"
	ZillowContextRightsUseAgent   ZillowContextRightsUse = "agent"
	ZillowContextRightsUseExport  ZillowContextRightsUse = "export"
)

// ZillowComparison: Up to five explicit provider regions at a common period/accepted snapshot.
// Gaps have a nullable value and explicit nullable reason; missing values are not zero.
type ZillowComparison struct {
	SchemaVersion     string                    `json:"schemaVersion"`
	DatasetKey        string                    `json:"datasetKey"`
	Period            string                    `json:"period"`
	ReferenceRegionID string                    `json:"referenceRegionId"`
	Metrics           []ZillowComparisonMetrics `json:"metrics"`
	Gaps              []ZillowComparisonGaps    `json:"gaps"`
}

// ZillowComparisonMetrics is generated from the OpenAPI spec.
type ZillowComparisonMetrics struct {
	Rights           *ZillowComparisonMetricsRights           `json:"rights,omitempty"`
	Metric           ZillowComparisonMetricsMetric            `json:"metric"`
	Status           ZillowComparisonMetricsStatus            `json:"status"`
	Reason           *string                                  `json:"reason,omitempty"`
	Definition       string                                   `json:"definition"`
	Unit             ZillowComparisonMetricsUnit              `json:"unit"`
	Value            *float64                                 `json:"value,omitempty"`
	Period           *string                                  `json:"period,omitempty"`
	Variant          *ZillowComparisonMetricsVariant          `json:"variant,omitempty"`
	Geography        *ZillowComparisonMetricsGeography        `json:"geography,omitempty"`
	Mapping          *ZillowComparisonMetricsMapping          `json:"mapping,omitempty"`
	Snapshot         *ZillowComparisonMetricsSnapshot         `json:"snapshot,omitempty"`
	AnnualChange     *ZillowComparisonMetricsAnnualChange     `json:"annualChange,omitempty"`
	MonthlyChange    *ZillowComparisonMetricsMonthlyChange    `json:"monthlyChange,omitempty"`
	RentAcceleration *ZillowComparisonMetricsRentAcceleration `json:"rentAcceleration,omitempty"`
	Points           []ZillowComparisonMetricsPoints          `json:"points"`
	SourceURL        string                                   `json:"sourceUrl"`
	Attribution      string                                   `json:"attribution"`
}

// UnmarshalJSON decodes ZillowComparisonMetrics, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowComparisonMetrics) UnmarshalJSON(data []byte) error {
	type plain ZillowComparisonMetrics
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowComparisonMetricsRights is generated from the OpenAPI spec.
type ZillowComparisonMetricsRights struct {
	Version     string  `json:"version"`
	EvidenceURL *string `json:"evidenceUrl,omitempty"`
	ExpiresAt   *string `json:"expiresAt,omitempty"`
}

// ZillowComparisonMetricsMetric is generated from the OpenAPI spec. It is a string; the
// ZillowComparisonMetricsMetric* constants list the documented values.
type ZillowComparisonMetricsMetric = string

// Documented values of ZillowComparisonMetricsMetric.
const (
	ZillowComparisonMetricsMetricZori                ZillowComparisonMetricsMetric = "zori"
	ZillowComparisonMetricsMetricZhvi                ZillowComparisonMetricsMetric = "zhvi"
	ZillowComparisonMetricsMetricInventory           ZillowComparisonMetricsMetric = "inventory"
	ZillowComparisonMetricsMetricPriceCutShare       ZillowComparisonMetricsMetric = "price_cut_share"
	ZillowComparisonMetricsMetricMedianDaysToPending ZillowComparisonMetricsMetric = "median_days_to_pending"
)

// ZillowComparisonMetricsStatus is generated from the OpenAPI spec. It is a string; the
// ZillowComparisonMetricsStatus* constants list the documented values.
type ZillowComparisonMetricsStatus = string

// Documented values of ZillowComparisonMetricsStatus.
const (
	ZillowComparisonMetricsStatusAvailable   ZillowComparisonMetricsStatus = "available"
	ZillowComparisonMetricsStatusUnavailable ZillowComparisonMetricsStatus = "unavailable"
)

// ZillowComparisonMetricsUnit is generated from the OpenAPI spec. It is a string; the
// ZillowComparisonMetricsUnit* constants list the documented values.
type ZillowComparisonMetricsUnit = string

// Documented values of ZillowComparisonMetricsUnit.
const (
	ZillowComparisonMetricsUnitUsd         ZillowComparisonMetricsUnit = "usd"
	ZillowComparisonMetricsUnitUsdPerMonth ZillowComparisonMetricsUnit = "usd_per_month"
	ZillowComparisonMetricsUnitCount       ZillowComparisonMetricsUnit = "count"
	ZillowComparisonMetricsUnitFraction    ZillowComparisonMetricsUnit = "fraction"
	ZillowComparisonMetricsUnitDays        ZillowComparisonMetricsUnit = "days"
)

// ZillowComparisonMetricsVariant is generated from the OpenAPI spec.
type ZillowComparisonMetricsVariant struct {
	DatasetKey         string                                           `json:"datasetKey"`
	RegistryVersion    int64                                            `json:"registryVersion"`
	Universe           string                                           `json:"universe"`
	Frequency          string                                           `json:"frequency"`
	Smoothing          string                                           `json:"smoothing"`
	SeasonalAdjustment ZillowComparisonMetricsVariantSeasonalAdjustment `json:"seasonalAdjustment"`
}

// UnmarshalJSON decodes ZillowComparisonMetricsVariant, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowComparisonMetricsVariant) UnmarshalJSON(data []byte) error {
	type plain ZillowComparisonMetricsVariant
	aux := struct {
		*plain
		RegistryVersion lenientNumber[int64] `json:"registryVersion"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.RegistryVersion.assign(&r.RegistryVersion)
	return softTypeError(err)
}

// ZillowComparisonMetricsVariantSeasonalAdjustment is generated from the OpenAPI spec. It is a
// string; the ZillowComparisonMetricsVariantSeasonalAdjustment* constants list the documented
// values.
type ZillowComparisonMetricsVariantSeasonalAdjustment = string

// Documented values of ZillowComparisonMetricsVariantSeasonalAdjustment.
const (
	ZillowComparisonMetricsVariantSeasonalAdjustmentSa        ZillowComparisonMetricsVariantSeasonalAdjustment = "sa"
	ZillowComparisonMetricsVariantSeasonalAdjustmentNotStated ZillowComparisonMetricsVariantSeasonalAdjustment = "not_stated"
)

// ZillowComparisonMetricsGeography is generated from the OpenAPI spec.
type ZillowComparisonMetricsGeography struct {
	ProviderID string                               `json:"providerId"`
	Name       string                               `json:"name"`
	Type       ZillowComparisonMetricsGeographyType `json:"type"`
}

// ZillowComparisonMetricsGeographyType is generated from the OpenAPI spec. It is a string; the
// ZillowComparisonMetricsGeographyType* constants list the documented values.
type ZillowComparisonMetricsGeographyType = string

// Documented values of ZillowComparisonMetricsGeographyType.
const (
	ZillowComparisonMetricsGeographyTypeCountry ZillowComparisonMetricsGeographyType = "country"
	ZillowComparisonMetricsGeographyTypeMsa     ZillowComparisonMetricsGeographyType = "msa"
	ZillowComparisonMetricsGeographyTypeCounty  ZillowComparisonMetricsGeographyType = "county"
	ZillowComparisonMetricsGeographyTypeZip     ZillowComparisonMetricsGeographyType = "zip"
)

// ZillowComparisonMetricsMapping is generated from the OpenAPI spec.
type ZillowComparisonMetricsMapping struct {
	Method         ZillowComparisonMetricsMappingMethod `json:"method"`
	Version        string                               `json:"version"`
	Source         string                               `json:"source"`
	FallbackReason *string                              `json:"fallbackReason,omitempty"`
}

// ZillowComparisonMetricsMappingMethod is generated from the OpenAPI spec. It is a string; the
// ZillowComparisonMetricsMappingMethod* constants list the documented values.
type ZillowComparisonMetricsMappingMethod = string

// Documented values of ZillowComparisonMetricsMappingMethod.
const (
	ZillowComparisonMetricsMappingMethodPostalZip              ZillowComparisonMetricsMappingMethod = "postal_zip"
	ZillowComparisonMetricsMappingMethodCountyFIPS             ZillowComparisonMetricsMappingMethod = "county_fips"
	ZillowComparisonMetricsMappingMethodVerifiedCrosswalk      ZillowComparisonMetricsMappingMethod = "verified_crosswalk"
	ZillowComparisonMetricsMappingMethodExplicitProviderRegion ZillowComparisonMetricsMappingMethod = "explicit_provider_region"
)

// ZillowComparisonMetricsSnapshot is generated from the OpenAPI spec.
type ZillowComparisonMetricsSnapshot struct {
	ID           string `json:"id"`
	Sha256       string `json:"sha256"`
	RetrievedAt  string `json:"retrievedAt"`
	AcceptedAt   string `json:"acceptedAt"`
	LatestPeriod string `json:"latestPeriod"`
	Stale        bool   `json:"stale"`
}

// ZillowComparisonMetricsAnnualChange is generated from the OpenAPI spec.
type ZillowComparisonMetricsAnnualChange struct {
	Value  *float64                                `json:"value,omitempty"`
	Unit   ZillowComparisonMetricsAnnualChangeUnit `json:"unit"`
	Reason *string                                 `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowComparisonMetricsAnnualChange, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowComparisonMetricsAnnualChange) UnmarshalJSON(data []byte) error {
	type plain ZillowComparisonMetricsAnnualChange
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowComparisonMetricsAnnualChangeUnit is generated from the OpenAPI spec. It is a string; the
// ZillowComparisonMetricsAnnualChangeUnit* constants list the documented values.
type ZillowComparisonMetricsAnnualChangeUnit = string

// Documented values of ZillowComparisonMetricsAnnualChangeUnit.
const (
	ZillowComparisonMetricsAnnualChangeUnitPercent          ZillowComparisonMetricsAnnualChangeUnit = "percent"
	ZillowComparisonMetricsAnnualChangeUnitPercentagePoints ZillowComparisonMetricsAnnualChangeUnit = "percentage_points"
	ZillowComparisonMetricsAnnualChangeUnitDays             ZillowComparisonMetricsAnnualChangeUnit = "days"
)

// ZillowComparisonMetricsMonthlyChange is generated from the OpenAPI spec.
type ZillowComparisonMetricsMonthlyChange struct {
	Value  *float64                                 `json:"value,omitempty"`
	Unit   ZillowComparisonMetricsMonthlyChangeUnit `json:"unit"`
	Reason *string                                  `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowComparisonMetricsMonthlyChange, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowComparisonMetricsMonthlyChange) UnmarshalJSON(data []byte) error {
	type plain ZillowComparisonMetricsMonthlyChange
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowComparisonMetricsMonthlyChangeUnit is generated from the OpenAPI spec. It is a string; the
// ZillowComparisonMetricsMonthlyChangeUnit* constants list the documented values.
type ZillowComparisonMetricsMonthlyChangeUnit = string

// Documented values of ZillowComparisonMetricsMonthlyChangeUnit.
const (
	ZillowComparisonMetricsMonthlyChangeUnitPercent          ZillowComparisonMetricsMonthlyChangeUnit = "percent"
	ZillowComparisonMetricsMonthlyChangeUnitPercentagePoints ZillowComparisonMetricsMonthlyChangeUnit = "percentage_points"
	ZillowComparisonMetricsMonthlyChangeUnitDays             ZillowComparisonMetricsMonthlyChangeUnit = "days"
)

// ZillowComparisonMetricsRentAcceleration is generated from the OpenAPI spec.
type ZillowComparisonMetricsRentAcceleration struct {
	Value  *float64                                    `json:"value,omitempty"`
	Unit   ZillowComparisonMetricsRentAccelerationUnit `json:"unit"`
	Reason *string                                     `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowComparisonMetricsRentAcceleration, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowComparisonMetricsRentAcceleration) UnmarshalJSON(data []byte) error {
	type plain ZillowComparisonMetricsRentAcceleration
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowComparisonMetricsRentAccelerationUnit is generated from the OpenAPI spec. It is a string;
// the ZillowComparisonMetricsRentAccelerationUnit* constants list the documented values.
type ZillowComparisonMetricsRentAccelerationUnit = string

// Documented values of ZillowComparisonMetricsRentAccelerationUnit.
const (
	ZillowComparisonMetricsRentAccelerationUnitPercent          ZillowComparisonMetricsRentAccelerationUnit = "percent"
	ZillowComparisonMetricsRentAccelerationUnitPercentagePoints ZillowComparisonMetricsRentAccelerationUnit = "percentage_points"
	ZillowComparisonMetricsRentAccelerationUnitDays             ZillowComparisonMetricsRentAccelerationUnit = "days"
)

// ZillowComparisonMetricsPoints is generated from the OpenAPI spec.
type ZillowComparisonMetricsPoints struct {
	Period string   `json:"period"`
	Value  *float64 `json:"value,omitempty"`
}

// UnmarshalJSON decodes ZillowComparisonMetricsPoints, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowComparisonMetricsPoints) UnmarshalJSON(data []byte) error {
	type plain ZillowComparisonMetricsPoints
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowComparisonGaps is generated from the OpenAPI spec.
type ZillowComparisonGaps struct {
	RegionID string                   `json:"regionId"`
	Value    *float64                 `json:"value,omitempty"`
	Reason   *string                  `json:"reason,omitempty"`
	Unit     ZillowComparisonGapsUnit `json:"unit"`
}

// UnmarshalJSON decodes ZillowComparisonGaps, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowComparisonGaps) UnmarshalJSON(data []byte) error {
	type plain ZillowComparisonGaps
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowComparisonGapsUnit is generated from the OpenAPI spec. It is a string; the
// ZillowComparisonGapsUnit* constants list the documented values.
type ZillowComparisonGapsUnit = string

// Documented values of ZillowComparisonGapsUnit.
const (
	ZillowComparisonGapsUnitDays             ZillowComparisonGapsUnit = "days"
	ZillowComparisonGapsUnitPercentagePoints ZillowComparisonGapsUnit = "percentage_points"
)

// IntelligenceParcelID is generated from the OpenAPI spec.
type IntelligenceParcelID = string

// IntelligenceInstant is generated from the OpenAPI spec.
type IntelligenceInstant = string

// IntelligenceRetainedID is generated from the OpenAPI spec.
type IntelligenceRetainedID = string

// IntelligenceCalculation is generated from the OpenAPI spec.
type IntelligenceCalculation struct {
	Definition  string                                         `json:"definition"`
	Status      IntelligenceCalculationStatus                  `json:"status"`
	Reasons     []string                                       `json:"reasons"`
	Metrics     map[string]IntelligenceCalculationMetricsValue `json:"metrics"`
	Counts      map[string]float64                             `json:"counts"`
	Exclusions  map[string]float64                             `json:"exclusions"`
	EvidenceIDs []string                                       `json:"evidence_ids"`
	Context     IntelligenceCalculationContext                 `json:"context"`
}

// UnmarshalJSON decodes IntelligenceCalculation, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *IntelligenceCalculation) UnmarshalJSON(data []byte) error {
	type plain IntelligenceCalculation
	aux := struct {
		*plain
		Counts     lenientMap[float64] `json:"counts"`
		Exclusions lenientMap[float64] `json:"exclusions"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Counts.assign(&r.Counts)
	aux.Exclusions.assign(&r.Exclusions)
	return softTypeError(err)
}

// IntelligenceCalculationStatus is generated from the OpenAPI spec. It is a string; the
// IntelligenceCalculationStatus* constants list the documented values.
type IntelligenceCalculationStatus = string

// Documented values of IntelligenceCalculationStatus.
const (
	IntelligenceCalculationStatusAvailable        IntelligenceCalculationStatus = "available"
	IntelligenceCalculationStatusPartial          IntelligenceCalculationStatus = "partial"
	IntelligenceCalculationStatusInsufficientData IntelligenceCalculationStatus = "insufficient_data"
	IntelligenceCalculationStatusUnavailable      IntelligenceCalculationStatus = "unavailable"
	IntelligenceCalculationStatusStale            IntelligenceCalculationStatus = "stale"
	IntelligenceCalculationStatusError            IntelligenceCalculationStatus = "error"
)

// IntelligenceCalculationMetricsValue is generated from the OpenAPI spec.
type IntelligenceCalculationMetricsValue struct {
	Value *IntelligenceCalculationMetricsValueValue `json:"value,omitempty"`
	Unit  string                                    `json:"unit"`
}

// IntelligenceCalculationMetricsValueValue is generated from the OpenAPI spec.
type IntelligenceCalculationMetricsValueValue struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
}

// IntelligenceCalculationContext is generated from the OpenAPI spec.
type IntelligenceCalculationContext struct {
	InputKind           IntelligenceCalculationContextInputKind       `json:"input_kind"`
	RequestedScope      *IntelligenceCalculationContextRequestedScope `json:"requested_scope,omitempty"`
	ActualScope         *IntelligenceCalculationContextActualScope    `json:"actual_scope,omitempty"`
	Query               *IntelligenceCalculationContextQuery          `json:"query,omitempty"`
	Period              json.RawMessage                               `json:"period,omitempty"`
	MetricPeriods       map[string]json.RawMessage                    `json:"metric_periods"`
	LatestObservationAt *string                                       `json:"latest_observation_at,omitempty"`
	SourceAsOf          *string                                       `json:"source_as_of,omitempty"`
	SourceVintage       *string                                       `json:"source_vintage,omitempty"`
	HistoryComplete     *bool                                         `json:"history_complete,omitempty"`
}

// IntelligenceCalculationContextInputKind is generated from the OpenAPI spec. It is a string; the
// IntelligenceCalculationContextInputKind* constants list the documented values.
type IntelligenceCalculationContextInputKind = string

// Documented values of IntelligenceCalculationContextInputKind.
const (
	IntelligenceCalculationContextInputKindObserved        IntelligenceCalculationContextInputKind = "observed"
	IntelligenceCalculationContextInputKindUserAssumptions IntelligenceCalculationContextInputKind = "user_assumptions"
	IntelligenceCalculationContextInputKindUnavailable     IntelligenceCalculationContextInputKind = "unavailable"
)

// IntelligenceCalculationContextRequestedScope is generated from the OpenAPI spec.
type IntelligenceCalculationContextRequestedScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceCalculationContextActualScope is generated from the OpenAPI spec.
type IntelligenceCalculationContextActualScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceCalculationContextQuery is generated from the OpenAPI spec.
type IntelligenceCalculationContextQuery struct {
	AsOf            string `json:"as_of"`
	KnowledgeCutoff string `json:"knowledge_cutoff"`
}

// IntelligenceResidualAssumptions is generated from the OpenAPI spec.
type IntelligenceResidualAssumptions struct {
	Currency              string                                 `json:"currency"`
	GrossCompletedSale    string                                 `json:"gross_completed_sale"`
	SellingCosts          string                                 `json:"selling_costs"`
	Costs                 []IntelligenceResidualAssumptionsCosts `json:"costs"`
	RequiredProfitDollars string                                 `json:"required_profit_dollars"`
	FixedAcquisitionCosts string                                 `json:"fixed_acquisition_costs"`
	AcquisitionCostRate   string                                 `json:"acquisition_cost_rate"`
	ProfitMode            string                                 `json:"profit_mode"`
	CarryMode             string                                 `json:"carry_mode"`
	InputSource           string                                 `json:"input_source"`
}

// IntelligenceResidualAssumptionsCosts is generated from the OpenAPI spec.
type IntelligenceResidualAssumptionsCosts struct {
	Bucket   IntelligenceResidualAssumptionsCostsBucket `json:"bucket"`
	Amount   string                                     `json:"amount"`
	Currency string                                     `json:"currency"`
}

// IntelligenceResidualAssumptionsCostsBucket is generated from the OpenAPI spec. It is a string;
// the IntelligenceResidualAssumptionsCostsBucket* constants list the documented values.
type IntelligenceResidualAssumptionsCostsBucket = string

// Documented values of IntelligenceResidualAssumptionsCostsBucket.
const (
	IntelligenceResidualAssumptionsCostsBucketHard         IntelligenceResidualAssumptionsCostsBucket = "hard"
	IntelligenceResidualAssumptionsCostsBucketSoft         IntelligenceResidualAssumptionsCostsBucket = "soft"
	IntelligenceResidualAssumptionsCostsBucketContingency  IntelligenceResidualAssumptionsCostsBucket = "contingency"
	IntelligenceResidualAssumptionsCostsBucketCarry        IntelligenceResidualAssumptionsCostsBucket = "carry"
	IntelligenceResidualAssumptionsCostsBucketOtherNonland IntelligenceResidualAssumptionsCostsBucket = "other_nonland"
)

// IntelligenceScenarioRevision is generated from the OpenAPI spec.
type IntelligenceScenarioRevision struct {
	ID               IntelligenceRetainedID            `json:"id"`
	RunID            IntelligenceRetainedID            `json:"run_id"`
	ParentRevisionID *IntelligenceRetainedID           `json:"parent_revision_id,omitempty"`
	Label            IntelligenceScenarioRevisionLabel `json:"label"`
	CreatedAt        IntelligenceInstant               `json:"created_at"`
	Assumptions      IntelligenceResidualAssumptions   `json:"assumptions"`
	Result           IntelligenceCalculation           `json:"result"`
}

// IntelligenceScenarioRevisionLabel is generated from the OpenAPI spec. It is a string; the
// IntelligenceScenarioRevisionLabel* constants list the documented values.
type IntelligenceScenarioRevisionLabel = string

// Documented values of IntelligenceScenarioRevisionLabel.
const (
	IntelligenceScenarioRevisionLabelBase     IntelligenceScenarioRevisionLabel = "base"
	IntelligenceScenarioRevisionLabelDownside IntelligenceScenarioRevisionLabel = "downside"
	IntelligenceScenarioRevisionLabelUpside   IntelligenceScenarioRevisionLabel = "upside"
)

// IntelligenceRights is generated from the OpenAPI spec.
type IntelligenceRights struct {
	Version       string `json:"version"`
	Display       bool   `json:"display"`
	Derived       bool   `json:"derived"`
	Cache         bool   `json:"cache"`
	RetainHistory bool   `json:"retain_history"`
	Export        bool   `json:"export"`
	Ai            bool   `json:"ai"`
}

// IntelligenceScope is generated from the OpenAPI spec.
type IntelligenceScope struct {
	Geography    string `json:"geography"`
	PropertyType string `json:"property_type"`
	Currency     string `json:"currency"`
}

// IntelligenceSourceCapability is generated from the OpenAPI spec.
type IntelligenceSourceCapability struct {
	ID                 string              `json:"id"`
	Version            string              `json:"version"`
	SourceProduct      string              `json:"source_product"`
	SourceVersion      string              `json:"source_version"`
	Scope              IntelligenceScope   `json:"scope"`
	Rights             *IntelligenceRights `json:"rights,omitempty"`
	AdditiveComponents bool                `json:"additive_components"`
	Complete           bool                `json:"complete"`
	MatureThrough      *string             `json:"mature_through,omitempty"`
	SourceAsOf         *string             `json:"source_as_of,omitempty"`
	StaleAfter         *string             `json:"stale_after,omitempty"`
}

// IntelligenceAssessmentObservation is generated from the OpenAPI spec.
type IntelligenceAssessmentObservation struct {
	Source                 string                                 `json:"source"`
	RecordID               string                                 `json:"record_id"`
	Version                string                                 `json:"version"`
	EvidenceID             string                                 `json:"evidence_id"`
	EffectiveAt            string                                 `json:"effective_at"`
	CapturedAt             string                                 `json:"captured_at"`
	ObservedAt             *string                                `json:"observed_at,omitempty"`
	Land                   *string                                `json:"land,omitempty"`
	Improvement            *string                                `json:"improvement,omitempty"`
	Basis                  IntelligenceAssessmentObservationBasis `json:"basis"`
	Vintage                string                                 `json:"vintage"`
	SourceProduct          string                                 `json:"source_product"`
	SourceVersion          string                                 `json:"source_version"`
	ComponentBasisVerified bool                                   `json:"component_basis_verified"`
}

// IntelligenceAssessmentObservationBasis is generated from the OpenAPI spec. It is a string; the
// IntelligenceAssessmentObservationBasis* constants list the documented values.
type IntelligenceAssessmentObservationBasis = string

// Documented values of IntelligenceAssessmentObservationBasis.
const (
	IntelligenceAssessmentObservationBasisAssessed  IntelligenceAssessmentObservationBasis = "assessed"
	IntelligenceAssessmentObservationBasisAppraised IntelligenceAssessmentObservationBasis = "appraised"
	IntelligenceAssessmentObservationBasisMarket    IntelligenceAssessmentObservationBasis = "market"
	IntelligenceAssessmentObservationBasisTaxable   IntelligenceAssessmentObservationBasis = "taxable"
)

// IntelligenceEvidence is generated from the OpenAPI spec.
type IntelligenceEvidence struct {
	ID                 IntelligenceRetainedID                 `json:"id"`
	SourceProduct      string                                 `json:"source_product"`
	SourceVersion      string                                 `json:"source_version"`
	CapturedAt         IntelligenceInstant                    `json:"captured_at"`
	ObservedAt         *string                                `json:"observed_at,omitempty"`
	SourceAsOf         *string                                `json:"source_as_of,omitempty"`
	SourceURL          *string                                `json:"source_url,omitempty"`
	KnowledgeBasis     IntelligenceEvidenceKnowledgeBasis     `json:"knowledge_basis"`
	EffectiveTimeBasis IntelligenceEvidenceEffectiveTimeBasis `json:"effective_time_basis"`
	Capability         IntelligenceSourceCapability           `json:"capability"`
	Assessment         IntelligenceAssessmentObservation      `json:"assessment"`
	SourceRecord       IntelligenceEvidenceSourceRecord       `json:"source_record"`
}

// IntelligenceEvidenceKnowledgeBasis is generated from the OpenAPI spec. It is a string; the
// IntelligenceEvidenceKnowledgeBasis* constants list the documented values.
type IntelligenceEvidenceKnowledgeBasis = string

// Documented values of IntelligenceEvidenceKnowledgeBasis.
const (
	IntelligenceEvidenceKnowledgeBasisSourceObservedAt     IntelligenceEvidenceKnowledgeBasis = "source_observed_at"
	IntelligenceEvidenceKnowledgeBasisFirstRetainedCapture IntelligenceEvidenceKnowledgeBasis = "first_retained_capture"
)

// IntelligenceEvidenceEffectiveTimeBasis is generated from the OpenAPI spec. It is a string; the
// IntelligenceEvidenceEffectiveTimeBasis* constants list the documented values.
type IntelligenceEvidenceEffectiveTimeBasis = string

// Documented values of IntelligenceEvidenceEffectiveTimeBasis.
const (
	IntelligenceEvidenceEffectiveTimeBasisSourceEventTime          IntelligenceEvidenceEffectiveTimeBasis = "source_event_time"
	IntelligenceEvidenceEffectiveTimeBasisFirstRetainedCaptureOnly IntelligenceEvidenceEffectiveTimeBasis = "first_retained_capture_only"
)

// IntelligenceEvidenceSourceRecord is generated from the OpenAPI spec.
type IntelligenceEvidenceSourceRecord struct {
	AssessmentYear *int64  `json:"assessment_year,omitempty"`
	TaxYear        *int64  `json:"tax_year,omitempty"`
	VintageYear    *int64  `json:"vintage_year,omitempty"`
	ReportedTotal  *string `json:"reported_total,omitempty"`
	TaxAmount      *string `json:"tax_amount,omitempty"`
	TaxPaidAmount  *string `json:"tax_paid_amount,omitempty"`
}

// UnmarshalJSON decodes IntelligenceEvidenceSourceRecord, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *IntelligenceEvidenceSourceRecord) UnmarshalJSON(data []byte) error {
	type plain IntelligenceEvidenceSourceRecord
	aux := struct {
		*plain
		AssessmentYear lenientNumber[int64] `json:"assessment_year"`
		TaxYear        lenientNumber[int64] `json:"tax_year"`
		VintageYear    lenientNumber[int64] `json:"vintage_year"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.AssessmentYear.assignPtr(&r.AssessmentYear)
	aux.TaxYear.assignPtr(&r.TaxYear)
	aux.VintageYear.assignPtr(&r.VintageYear)
	return softTypeError(err)
}

// IntelligenceRunDetail is generated from the OpenAPI spec.
type IntelligenceRunDetail struct {
	Run      IntelligenceRun        `json:"run"`
	Evidence []IntelligenceEvidence `json:"evidence"`
}

// ZillowPropertyComparison is generated from the OpenAPI spec.
type ZillowPropertyComparison struct {
	SchemaVersion string                             `json:"schemaVersion"`
	CanonicalID   *string                            `json:"canonical_id,omitempty"`
	Metric        ZillowPropertyComparisonMetric     `json:"metric"`
	Period        string                             `json:"period"`
	AsOf          string                             `json:"asOf"`
	Reference     *ZillowPropertyComparisonReference `json:"reference,omitempty"`
	Items         []ZillowPropertyComparisonItems    `json:"items"`
}

// ZillowPropertyComparisonMetric is generated from the OpenAPI spec. It is a string; the
// ZillowPropertyComparisonMetric* constants list the documented values.
type ZillowPropertyComparisonMetric = string

// Documented values of ZillowPropertyComparisonMetric.
const (
	ZillowPropertyComparisonMetricZori                ZillowPropertyComparisonMetric = "zori"
	ZillowPropertyComparisonMetricZhvi                ZillowPropertyComparisonMetric = "zhvi"
	ZillowPropertyComparisonMetricInventory           ZillowPropertyComparisonMetric = "inventory"
	ZillowPropertyComparisonMetricPriceCutShare       ZillowPropertyComparisonMetric = "price_cut_share"
	ZillowPropertyComparisonMetricMedianDaysToPending ZillowPropertyComparisonMetric = "median_days_to_pending"
)

// ZillowPropertyComparisonReference is generated from the OpenAPI spec. It is a string; the
// ZillowPropertyComparisonReference* constants list the documented values.
type ZillowPropertyComparisonReference = string

// Documented values of ZillowPropertyComparisonReference.
const (
	ZillowPropertyComparisonReferenceCounty   ZillowPropertyComparisonReference = "county"
	ZillowPropertyComparisonReferenceMetro    ZillowPropertyComparisonReference = "metro"
	ZillowPropertyComparisonReferenceNational ZillowPropertyComparisonReference = "national"
)

// ZillowPropertyComparisonItems is generated from the OpenAPI spec.
type ZillowPropertyComparisonItems struct {
	Level      ZillowPropertyComparisonItemsLevel  `json:"level"`
	DatasetKey *string                             `json:"datasetKey,omitempty"`
	Result     ZillowPropertyComparisonItemsResult `json:"result"`
	Gap        ZillowPropertyComparisonItemsGap    `json:"gap"`
}

// ZillowPropertyComparisonItemsLevel is generated from the OpenAPI spec. It is a string; the
// ZillowPropertyComparisonItemsLevel* constants list the documented values.
type ZillowPropertyComparisonItemsLevel = string

// Documented values of ZillowPropertyComparisonItemsLevel.
const (
	ZillowPropertyComparisonItemsLevelCounty   ZillowPropertyComparisonItemsLevel = "county"
	ZillowPropertyComparisonItemsLevelMetro    ZillowPropertyComparisonItemsLevel = "metro"
	ZillowPropertyComparisonItemsLevelNational ZillowPropertyComparisonItemsLevel = "national"
)

// ZillowPropertyComparisonItemsResult is generated from the OpenAPI spec.
type ZillowPropertyComparisonItemsResult struct {
	Rights           *ZillowPropertyComparisonItemsResultRights           `json:"rights,omitempty"`
	Metric           ZillowPropertyComparisonItemsResultMetric            `json:"metric"`
	Status           ZillowPropertyComparisonItemsResultStatus            `json:"status"`
	Reason           *ZillowPropertyComparisonItemsResultReason           `json:"reason,omitempty"`
	Definition       string                                               `json:"definition"`
	Unit             ZillowPropertyComparisonItemsResultUnit              `json:"unit"`
	Value            *float64                                             `json:"value,omitempty"`
	Period           *string                                              `json:"period,omitempty"`
	Variant          *ZillowPropertyComparisonItemsResultVariant          `json:"variant,omitempty"`
	Geography        *ZillowPropertyComparisonItemsResultGeography        `json:"geography,omitempty"`
	Mapping          *ZillowPropertyComparisonItemsResultMapping          `json:"mapping,omitempty"`
	Snapshot         *ZillowPropertyComparisonItemsResultSnapshot         `json:"snapshot,omitempty"`
	AnnualChange     *ZillowPropertyComparisonItemsResultAnnualChange     `json:"annualChange,omitempty"`
	MonthlyChange    *ZillowPropertyComparisonItemsResultMonthlyChange    `json:"monthlyChange,omitempty"`
	RentAcceleration *ZillowPropertyComparisonItemsResultRentAcceleration `json:"rentAcceleration,omitempty"`
	Points           []ZillowPropertyComparisonItemsResultPoints          `json:"points"`
	SourceURL        string                                               `json:"sourceUrl"`
	Attribution      string                                               `json:"attribution"`
}

// UnmarshalJSON decodes ZillowPropertyComparisonItemsResult, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowPropertyComparisonItemsResult) UnmarshalJSON(data []byte) error {
	type plain ZillowPropertyComparisonItemsResult
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowPropertyComparisonItemsResultRights is generated from the OpenAPI spec.
type ZillowPropertyComparisonItemsResultRights struct {
	Version     string  `json:"version"`
	EvidenceURL *string `json:"evidenceUrl,omitempty"`
	ExpiresAt   *string `json:"expiresAt,omitempty"`
}

// ZillowPropertyComparisonItemsResultMetric is generated from the OpenAPI spec. It is a string;
// the ZillowPropertyComparisonItemsResultMetric* constants list the documented values.
type ZillowPropertyComparisonItemsResultMetric = string

// Documented values of ZillowPropertyComparisonItemsResultMetric.
const (
	ZillowPropertyComparisonItemsResultMetricZori                ZillowPropertyComparisonItemsResultMetric = "zori"
	ZillowPropertyComparisonItemsResultMetricZhvi                ZillowPropertyComparisonItemsResultMetric = "zhvi"
	ZillowPropertyComparisonItemsResultMetricInventory           ZillowPropertyComparisonItemsResultMetric = "inventory"
	ZillowPropertyComparisonItemsResultMetricPriceCutShare       ZillowPropertyComparisonItemsResultMetric = "price_cut_share"
	ZillowPropertyComparisonItemsResultMetricMedianDaysToPending ZillowPropertyComparisonItemsResultMetric = "median_days_to_pending"
)

// ZillowPropertyComparisonItemsResultStatus is generated from the OpenAPI spec. It is a string;
// the ZillowPropertyComparisonItemsResultStatus* constants list the documented values.
type ZillowPropertyComparisonItemsResultStatus = string

// Documented values of ZillowPropertyComparisonItemsResultStatus.
const (
	ZillowPropertyComparisonItemsResultStatusAvailable   ZillowPropertyComparisonItemsResultStatus = "available"
	ZillowPropertyComparisonItemsResultStatusUnavailable ZillowPropertyComparisonItemsResultStatus = "unavailable"
)

// ZillowPropertyComparisonItemsResultReason is generated from the OpenAPI spec. It is a string;
// the ZillowPropertyComparisonItemsResultReason* constants list the documented values.
type ZillowPropertyComparisonItemsResultReason = string

// Documented values of ZillowPropertyComparisonItemsResultReason.
const (
	ZillowPropertyComparisonItemsResultReasonMissingPeriod                ZillowPropertyComparisonItemsResultReason = "missing_period"
	ZillowPropertyComparisonItemsResultReasonNonpositiveDenominator       ZillowPropertyComparisonItemsResultReason = "nonpositive_denominator"
	ZillowPropertyComparisonItemsResultReasonNonFiniteResult              ZillowPropertyComparisonItemsResultReason = "non_finite_result"
	ZillowPropertyComparisonItemsResultReasonNotApplicable                ZillowPropertyComparisonItemsResultReason = "not_applicable"
	ZillowPropertyComparisonItemsResultReasonMappingUnavailable           ZillowPropertyComparisonItemsResultReason = "mapping_unavailable"
	ZillowPropertyComparisonItemsResultReasonNoCoverage                   ZillowPropertyComparisonItemsResultReason = "no_coverage"
	ZillowPropertyComparisonItemsResultReasonSuppressed                   ZillowPropertyComparisonItemsResultReason = "suppressed"
	ZillowPropertyComparisonItemsResultReasonHistoricalVintageUnavailable ZillowPropertyComparisonItemsResultReason = "historical_vintage_unavailable"
	ZillowPropertyComparisonItemsResultReasonRightsUnavailable            ZillowPropertyComparisonItemsResultReason = "rights_unavailable"
	ZillowPropertyComparisonItemsResultReasonDisabled                     ZillowPropertyComparisonItemsResultReason = "disabled"
	ZillowPropertyComparisonItemsResultReasonIncompatibleVariant          ZillowPropertyComparisonItemsResultReason = "incompatible_variant"
)

// ZillowPropertyComparisonItemsResultUnit is generated from the OpenAPI spec. It is a string; the
// ZillowPropertyComparisonItemsResultUnit* constants list the documented values.
type ZillowPropertyComparisonItemsResultUnit = string

// Documented values of ZillowPropertyComparisonItemsResultUnit.
const (
	ZillowPropertyComparisonItemsResultUnitUsd         ZillowPropertyComparisonItemsResultUnit = "usd"
	ZillowPropertyComparisonItemsResultUnitUsdPerMonth ZillowPropertyComparisonItemsResultUnit = "usd_per_month"
	ZillowPropertyComparisonItemsResultUnitCount       ZillowPropertyComparisonItemsResultUnit = "count"
	ZillowPropertyComparisonItemsResultUnitFraction    ZillowPropertyComparisonItemsResultUnit = "fraction"
	ZillowPropertyComparisonItemsResultUnitDays        ZillowPropertyComparisonItemsResultUnit = "days"
)

// ZillowPropertyComparisonItemsResultVariant is generated from the OpenAPI spec.
type ZillowPropertyComparisonItemsResultVariant struct {
	DatasetKey         string                                                       `json:"datasetKey"`
	RegistryVersion    int64                                                        `json:"registryVersion"`
	Universe           string                                                       `json:"universe"`
	Frequency          string                                                       `json:"frequency"`
	Smoothing          string                                                       `json:"smoothing"`
	SeasonalAdjustment ZillowPropertyComparisonItemsResultVariantSeasonalAdjustment `json:"seasonalAdjustment"`
}

// UnmarshalJSON decodes ZillowPropertyComparisonItemsResultVariant, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowPropertyComparisonItemsResultVariant) UnmarshalJSON(data []byte) error {
	type plain ZillowPropertyComparisonItemsResultVariant
	aux := struct {
		*plain
		RegistryVersion lenientNumber[int64] `json:"registryVersion"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.RegistryVersion.assign(&r.RegistryVersion)
	return softTypeError(err)
}

// ZillowPropertyComparisonItemsResultVariantSeasonalAdjustment is generated from the OpenAPI spec.
// It is a string; the ZillowPropertyComparisonItemsResultVariantSeasonalAdjustment* constants list
// the documented values.
type ZillowPropertyComparisonItemsResultVariantSeasonalAdjustment = string

// Documented values of ZillowPropertyComparisonItemsResultVariantSeasonalAdjustment.
const (
	ZillowPropertyComparisonItemsResultVariantSeasonalAdjustmentSa        ZillowPropertyComparisonItemsResultVariantSeasonalAdjustment = "sa"
	ZillowPropertyComparisonItemsResultVariantSeasonalAdjustmentNotStated ZillowPropertyComparisonItemsResultVariantSeasonalAdjustment = "not_stated"
)

// ZillowPropertyComparisonItemsResultGeography is generated from the OpenAPI spec.
type ZillowPropertyComparisonItemsResultGeography struct {
	ProviderID string                                           `json:"providerId"`
	Name       string                                           `json:"name"`
	Type       ZillowPropertyComparisonItemsResultGeographyType `json:"type"`
}

// ZillowPropertyComparisonItemsResultGeographyType is generated from the OpenAPI spec. It is a
// string; the ZillowPropertyComparisonItemsResultGeographyType* constants list the documented
// values.
type ZillowPropertyComparisonItemsResultGeographyType = string

// Documented values of ZillowPropertyComparisonItemsResultGeographyType.
const (
	ZillowPropertyComparisonItemsResultGeographyTypeCountry ZillowPropertyComparisonItemsResultGeographyType = "country"
	ZillowPropertyComparisonItemsResultGeographyTypeMsa     ZillowPropertyComparisonItemsResultGeographyType = "msa"
	ZillowPropertyComparisonItemsResultGeographyTypeCounty  ZillowPropertyComparisonItemsResultGeographyType = "county"
	ZillowPropertyComparisonItemsResultGeographyTypeZip     ZillowPropertyComparisonItemsResultGeographyType = "zip"
)

// ZillowPropertyComparisonItemsResultMapping is generated from the OpenAPI spec.
type ZillowPropertyComparisonItemsResultMapping struct {
	Method         ZillowPropertyComparisonItemsResultMappingMethod `json:"method"`
	Version        string                                           `json:"version"`
	Source         string                                           `json:"source"`
	FallbackReason *string                                          `json:"fallbackReason,omitempty"`
}

// ZillowPropertyComparisonItemsResultMappingMethod is generated from the OpenAPI spec. It is a
// string; the ZillowPropertyComparisonItemsResultMappingMethod* constants list the documented
// values.
type ZillowPropertyComparisonItemsResultMappingMethod = string

// Documented values of ZillowPropertyComparisonItemsResultMappingMethod.
const (
	ZillowPropertyComparisonItemsResultMappingMethodPostalZip              ZillowPropertyComparisonItemsResultMappingMethod = "postal_zip"
	ZillowPropertyComparisonItemsResultMappingMethodCountyFIPS             ZillowPropertyComparisonItemsResultMappingMethod = "county_fips"
	ZillowPropertyComparisonItemsResultMappingMethodVerifiedCrosswalk      ZillowPropertyComparisonItemsResultMappingMethod = "verified_crosswalk"
	ZillowPropertyComparisonItemsResultMappingMethodExplicitProviderRegion ZillowPropertyComparisonItemsResultMappingMethod = "explicit_provider_region"
)

// ZillowPropertyComparisonItemsResultSnapshot is generated from the OpenAPI spec.
type ZillowPropertyComparisonItemsResultSnapshot struct {
	ID           string `json:"id"`
	Sha256       string `json:"sha256"`
	RetrievedAt  string `json:"retrievedAt"`
	AcceptedAt   string `json:"acceptedAt"`
	LatestPeriod string `json:"latestPeriod"`
	Stale        bool   `json:"stale"`
}

// ZillowPropertyComparisonItemsResultAnnualChange is generated from the OpenAPI spec.
type ZillowPropertyComparisonItemsResultAnnualChange struct {
	Value  *float64                                               `json:"value,omitempty"`
	Unit   ZillowPropertyComparisonItemsResultAnnualChangeUnit    `json:"unit"`
	Reason *ZillowPropertyComparisonItemsResultAnnualChangeReason `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowPropertyComparisonItemsResultAnnualChange, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowPropertyComparisonItemsResultAnnualChange) UnmarshalJSON(data []byte) error {
	type plain ZillowPropertyComparisonItemsResultAnnualChange
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowPropertyComparisonItemsResultAnnualChangeUnit is generated from the OpenAPI spec. It is a
// string; the ZillowPropertyComparisonItemsResultAnnualChangeUnit* constants list the documented
// values.
type ZillowPropertyComparisonItemsResultAnnualChangeUnit = string

// Documented values of ZillowPropertyComparisonItemsResultAnnualChangeUnit.
const (
	ZillowPropertyComparisonItemsResultAnnualChangeUnitPercent          ZillowPropertyComparisonItemsResultAnnualChangeUnit = "percent"
	ZillowPropertyComparisonItemsResultAnnualChangeUnitPercentagePoints ZillowPropertyComparisonItemsResultAnnualChangeUnit = "percentage_points"
	ZillowPropertyComparisonItemsResultAnnualChangeUnitDays             ZillowPropertyComparisonItemsResultAnnualChangeUnit = "days"
)

// ZillowPropertyComparisonItemsResultAnnualChangeReason is generated from the OpenAPI spec. It is
// a string; the ZillowPropertyComparisonItemsResultAnnualChangeReason* constants list the
// documented values.
type ZillowPropertyComparisonItemsResultAnnualChangeReason = string

// Documented values of ZillowPropertyComparisonItemsResultAnnualChangeReason.
const (
	ZillowPropertyComparisonItemsResultAnnualChangeReasonMissingPeriod                ZillowPropertyComparisonItemsResultAnnualChangeReason = "missing_period"
	ZillowPropertyComparisonItemsResultAnnualChangeReasonNonpositiveDenominator       ZillowPropertyComparisonItemsResultAnnualChangeReason = "nonpositive_denominator"
	ZillowPropertyComparisonItemsResultAnnualChangeReasonNonFiniteResult              ZillowPropertyComparisonItemsResultAnnualChangeReason = "non_finite_result"
	ZillowPropertyComparisonItemsResultAnnualChangeReasonNotApplicable                ZillowPropertyComparisonItemsResultAnnualChangeReason = "not_applicable"
	ZillowPropertyComparisonItemsResultAnnualChangeReasonMappingUnavailable           ZillowPropertyComparisonItemsResultAnnualChangeReason = "mapping_unavailable"
	ZillowPropertyComparisonItemsResultAnnualChangeReasonNoCoverage                   ZillowPropertyComparisonItemsResultAnnualChangeReason = "no_coverage"
	ZillowPropertyComparisonItemsResultAnnualChangeReasonSuppressed                   ZillowPropertyComparisonItemsResultAnnualChangeReason = "suppressed"
	ZillowPropertyComparisonItemsResultAnnualChangeReasonHistoricalVintageUnavailable ZillowPropertyComparisonItemsResultAnnualChangeReason = "historical_vintage_unavailable"
	ZillowPropertyComparisonItemsResultAnnualChangeReasonRightsUnavailable            ZillowPropertyComparisonItemsResultAnnualChangeReason = "rights_unavailable"
	ZillowPropertyComparisonItemsResultAnnualChangeReasonDisabled                     ZillowPropertyComparisonItemsResultAnnualChangeReason = "disabled"
	ZillowPropertyComparisonItemsResultAnnualChangeReasonIncompatibleVariant          ZillowPropertyComparisonItemsResultAnnualChangeReason = "incompatible_variant"
)

// ZillowPropertyComparisonItemsResultMonthlyChange is generated from the OpenAPI spec.
type ZillowPropertyComparisonItemsResultMonthlyChange struct {
	Value  *float64                                                `json:"value,omitempty"`
	Unit   ZillowPropertyComparisonItemsResultMonthlyChangeUnit    `json:"unit"`
	Reason *ZillowPropertyComparisonItemsResultMonthlyChangeReason `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowPropertyComparisonItemsResultMonthlyChange, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowPropertyComparisonItemsResultMonthlyChange) UnmarshalJSON(data []byte) error {
	type plain ZillowPropertyComparisonItemsResultMonthlyChange
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowPropertyComparisonItemsResultMonthlyChangeUnit is generated from the OpenAPI spec. It is a
// string; the ZillowPropertyComparisonItemsResultMonthlyChangeUnit* constants list the documented
// values.
type ZillowPropertyComparisonItemsResultMonthlyChangeUnit = string

// Documented values of ZillowPropertyComparisonItemsResultMonthlyChangeUnit.
const (
	ZillowPropertyComparisonItemsResultMonthlyChangeUnitPercent          ZillowPropertyComparisonItemsResultMonthlyChangeUnit = "percent"
	ZillowPropertyComparisonItemsResultMonthlyChangeUnitPercentagePoints ZillowPropertyComparisonItemsResultMonthlyChangeUnit = "percentage_points"
	ZillowPropertyComparisonItemsResultMonthlyChangeUnitDays             ZillowPropertyComparisonItemsResultMonthlyChangeUnit = "days"
)

// ZillowPropertyComparisonItemsResultMonthlyChangeReason is generated from the OpenAPI spec. It is
// a string; the ZillowPropertyComparisonItemsResultMonthlyChangeReason* constants list the
// documented values.
type ZillowPropertyComparisonItemsResultMonthlyChangeReason = string

// Documented values of ZillowPropertyComparisonItemsResultMonthlyChangeReason.
const (
	ZillowPropertyComparisonItemsResultMonthlyChangeReasonMissingPeriod                ZillowPropertyComparisonItemsResultMonthlyChangeReason = "missing_period"
	ZillowPropertyComparisonItemsResultMonthlyChangeReasonNonpositiveDenominator       ZillowPropertyComparisonItemsResultMonthlyChangeReason = "nonpositive_denominator"
	ZillowPropertyComparisonItemsResultMonthlyChangeReasonNonFiniteResult              ZillowPropertyComparisonItemsResultMonthlyChangeReason = "non_finite_result"
	ZillowPropertyComparisonItemsResultMonthlyChangeReasonNotApplicable                ZillowPropertyComparisonItemsResultMonthlyChangeReason = "not_applicable"
	ZillowPropertyComparisonItemsResultMonthlyChangeReasonMappingUnavailable           ZillowPropertyComparisonItemsResultMonthlyChangeReason = "mapping_unavailable"
	ZillowPropertyComparisonItemsResultMonthlyChangeReasonNoCoverage                   ZillowPropertyComparisonItemsResultMonthlyChangeReason = "no_coverage"
	ZillowPropertyComparisonItemsResultMonthlyChangeReasonSuppressed                   ZillowPropertyComparisonItemsResultMonthlyChangeReason = "suppressed"
	ZillowPropertyComparisonItemsResultMonthlyChangeReasonHistoricalVintageUnavailable ZillowPropertyComparisonItemsResultMonthlyChangeReason = "historical_vintage_unavailable"
	ZillowPropertyComparisonItemsResultMonthlyChangeReasonRightsUnavailable            ZillowPropertyComparisonItemsResultMonthlyChangeReason = "rights_unavailable"
	ZillowPropertyComparisonItemsResultMonthlyChangeReasonDisabled                     ZillowPropertyComparisonItemsResultMonthlyChangeReason = "disabled"
	ZillowPropertyComparisonItemsResultMonthlyChangeReasonIncompatibleVariant          ZillowPropertyComparisonItemsResultMonthlyChangeReason = "incompatible_variant"
)

// ZillowPropertyComparisonItemsResultRentAcceleration is generated from the OpenAPI spec.
type ZillowPropertyComparisonItemsResultRentAcceleration struct {
	Value  *float64                                                   `json:"value,omitempty"`
	Unit   ZillowPropertyComparisonItemsResultRentAccelerationUnit    `json:"unit"`
	Reason *ZillowPropertyComparisonItemsResultRentAccelerationReason `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowPropertyComparisonItemsResultRentAcceleration, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowPropertyComparisonItemsResultRentAcceleration) UnmarshalJSON(data []byte) error {
	type plain ZillowPropertyComparisonItemsResultRentAcceleration
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowPropertyComparisonItemsResultRentAccelerationUnit is generated from the OpenAPI spec. It
// is a string; the ZillowPropertyComparisonItemsResultRentAccelerationUnit* constants list the
// documented values.
type ZillowPropertyComparisonItemsResultRentAccelerationUnit = string

// Documented values of ZillowPropertyComparisonItemsResultRentAccelerationUnit.
const (
	ZillowPropertyComparisonItemsResultRentAccelerationUnitPercent          ZillowPropertyComparisonItemsResultRentAccelerationUnit = "percent"
	ZillowPropertyComparisonItemsResultRentAccelerationUnitPercentagePoints ZillowPropertyComparisonItemsResultRentAccelerationUnit = "percentage_points"
	ZillowPropertyComparisonItemsResultRentAccelerationUnitDays             ZillowPropertyComparisonItemsResultRentAccelerationUnit = "days"
)

// ZillowPropertyComparisonItemsResultRentAccelerationReason is generated from the OpenAPI spec. It
// is a string; the ZillowPropertyComparisonItemsResultRentAccelerationReason* constants list the
// documented values.
type ZillowPropertyComparisonItemsResultRentAccelerationReason = string

// Documented values of ZillowPropertyComparisonItemsResultRentAccelerationReason.
const (
	ZillowPropertyComparisonItemsResultRentAccelerationReasonMissingPeriod                ZillowPropertyComparisonItemsResultRentAccelerationReason = "missing_period"
	ZillowPropertyComparisonItemsResultRentAccelerationReasonNonpositiveDenominator       ZillowPropertyComparisonItemsResultRentAccelerationReason = "nonpositive_denominator"
	ZillowPropertyComparisonItemsResultRentAccelerationReasonNonFiniteResult              ZillowPropertyComparisonItemsResultRentAccelerationReason = "non_finite_result"
	ZillowPropertyComparisonItemsResultRentAccelerationReasonNotApplicable                ZillowPropertyComparisonItemsResultRentAccelerationReason = "not_applicable"
	ZillowPropertyComparisonItemsResultRentAccelerationReasonMappingUnavailable           ZillowPropertyComparisonItemsResultRentAccelerationReason = "mapping_unavailable"
	ZillowPropertyComparisonItemsResultRentAccelerationReasonNoCoverage                   ZillowPropertyComparisonItemsResultRentAccelerationReason = "no_coverage"
	ZillowPropertyComparisonItemsResultRentAccelerationReasonSuppressed                   ZillowPropertyComparisonItemsResultRentAccelerationReason = "suppressed"
	ZillowPropertyComparisonItemsResultRentAccelerationReasonHistoricalVintageUnavailable ZillowPropertyComparisonItemsResultRentAccelerationReason = "historical_vintage_unavailable"
	ZillowPropertyComparisonItemsResultRentAccelerationReasonRightsUnavailable            ZillowPropertyComparisonItemsResultRentAccelerationReason = "rights_unavailable"
	ZillowPropertyComparisonItemsResultRentAccelerationReasonDisabled                     ZillowPropertyComparisonItemsResultRentAccelerationReason = "disabled"
	ZillowPropertyComparisonItemsResultRentAccelerationReasonIncompatibleVariant          ZillowPropertyComparisonItemsResultRentAccelerationReason = "incompatible_variant"
)

// ZillowPropertyComparisonItemsResultPoints is generated from the OpenAPI spec.
type ZillowPropertyComparisonItemsResultPoints struct {
	Period string   `json:"period"`
	Value  *float64 `json:"value,omitempty"`
}

// UnmarshalJSON decodes ZillowPropertyComparisonItemsResultPoints, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowPropertyComparisonItemsResultPoints) UnmarshalJSON(data []byte) error {
	type plain ZillowPropertyComparisonItemsResultPoints
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowPropertyComparisonItemsGap is generated from the OpenAPI spec.
type ZillowPropertyComparisonItemsGap struct {
	Value  *float64                                `json:"value,omitempty"`
	Unit   ZillowPropertyComparisonItemsGapUnit    `json:"unit"`
	Reason *ZillowPropertyComparisonItemsGapReason `json:"reason,omitempty"`
}

// UnmarshalJSON decodes ZillowPropertyComparisonItemsGap, accepting numeric fields sent as JSON numbers or as
// quoted decimal strings.
func (r *ZillowPropertyComparisonItemsGap) UnmarshalJSON(data []byte) error {
	type plain ZillowPropertyComparisonItemsGap
	aux := struct {
		*plain
		Value lenientNumber[float64] `json:"value"`
	}{plain: (*plain)(r)}
	err := json.Unmarshal(data, &aux)
	aux.Value.assignPtr(&r.Value)
	return softTypeError(err)
}

// ZillowPropertyComparisonItemsGapUnit is generated from the OpenAPI spec. It is a string; the
// ZillowPropertyComparisonItemsGapUnit* constants list the documented values.
type ZillowPropertyComparisonItemsGapUnit = string

// Documented values of ZillowPropertyComparisonItemsGapUnit.
const (
	ZillowPropertyComparisonItemsGapUnitPercent          ZillowPropertyComparisonItemsGapUnit = "percent"
	ZillowPropertyComparisonItemsGapUnitPercentagePoints ZillowPropertyComparisonItemsGapUnit = "percentage_points"
	ZillowPropertyComparisonItemsGapUnitDays             ZillowPropertyComparisonItemsGapUnit = "days"
)

// ZillowPropertyComparisonItemsGapReason is generated from the OpenAPI spec. It is a string; the
// ZillowPropertyComparisonItemsGapReason* constants list the documented values.
type ZillowPropertyComparisonItemsGapReason = string

// Documented values of ZillowPropertyComparisonItemsGapReason.
const (
	ZillowPropertyComparisonItemsGapReasonMissingPeriod                ZillowPropertyComparisonItemsGapReason = "missing_period"
	ZillowPropertyComparisonItemsGapReasonNonpositiveDenominator       ZillowPropertyComparisonItemsGapReason = "nonpositive_denominator"
	ZillowPropertyComparisonItemsGapReasonNonFiniteResult              ZillowPropertyComparisonItemsGapReason = "non_finite_result"
	ZillowPropertyComparisonItemsGapReasonNotApplicable                ZillowPropertyComparisonItemsGapReason = "not_applicable"
	ZillowPropertyComparisonItemsGapReasonMappingUnavailable           ZillowPropertyComparisonItemsGapReason = "mapping_unavailable"
	ZillowPropertyComparisonItemsGapReasonNoCoverage                   ZillowPropertyComparisonItemsGapReason = "no_coverage"
	ZillowPropertyComparisonItemsGapReasonSuppressed                   ZillowPropertyComparisonItemsGapReason = "suppressed"
	ZillowPropertyComparisonItemsGapReasonHistoricalVintageUnavailable ZillowPropertyComparisonItemsGapReason = "historical_vintage_unavailable"
	ZillowPropertyComparisonItemsGapReasonRightsUnavailable            ZillowPropertyComparisonItemsGapReason = "rights_unavailable"
	ZillowPropertyComparisonItemsGapReasonDisabled                     ZillowPropertyComparisonItemsGapReason = "disabled"
	ZillowPropertyComparisonItemsGapReasonIncompatibleVariant          ZillowPropertyComparisonItemsGapReason = "incompatible_variant"
)

// ZillowComparisonResponse is generated from the OpenAPI spec.
type ZillowComparisonResponse = json.RawMessage
