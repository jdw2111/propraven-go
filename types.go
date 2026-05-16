package propraven

// Parcel is the canonical parcel surface exposed by /v1/parcels.
// Fields mirror the OpenAPI 3.1 spec at https://api.propraven.com/openapi.json
// and the columns documented at META.ENTITY_DICTIONARY on the Snowflake share.
type Parcel struct {
	ParcelID       string   `json:"parcel_id"`
	APN            string   `json:"apn,omitempty"`
	StateFIPS      string   `json:"state_fips"`
	CountyFIPS     string   `json:"county_fips"`
	LotSizeAcres   *float64 `json:"lot_size_acres,omitempty"`
	YearBuilt      *int     `json:"year_built,omitempty"`
	BuildingSqft   *int     `json:"building_sqft,omitempty"`
	PropertyType   string   `json:"property_type_class,omitempty"`
	OwnerName      string   `json:"owner_name,omitempty"`
	OwnerEntityID  string   `json:"owner_entity_id,omitempty"`
	Latitude       *float64 `json:"latitude,omitempty"`
	Longitude      *float64 `json:"longitude,omitempty"`
	AbsenteeOwner  *bool    `json:"absentee_owner,omitempty"`
	LastRefreshed  string   `json:"last_refreshed_at,omitempty"`
}

// Owner is the consolidated entity surface exposed by /v1/owners.
type Owner struct {
	OwnerEntityID string   `json:"owner_entity_id"`
	OwnerName     string   `json:"owner_name"`
	OwnerType     string   `json:"owner_type,omitempty"`
	ParcelCount   *int     `json:"parcel_count,omitempty"`
	Ticker        string   `json:"ticker,omitempty"`
	ParentEntity  string   `json:"parent_entity_id,omitempty"`
	StatesActive  []string `json:"states_active,omitempty"`
}

// Deed is a matched deed transfer.
type Deed struct {
	DeedID         string   `json:"deed_id"`
	ParcelID       string   `json:"parcel_id"`
	StateFIPS      string   `json:"state_fips"`
	CountyFIPS     string   `json:"county_fips"`
	SaleDate       string   `json:"sale_date,omitempty"`
	SalePriceUSD   *float64 `json:"sale_price_usd,omitempty"`
	IsArmLength    *bool    `json:"is_arm_length,omitempty"`
	Grantor        string   `json:"grantor,omitempty"`
	Grantee        string   `json:"grantee,omitempty"`
	DocumentType   string   `json:"document_type,omitempty"`
	DocumentNumber string   `json:"document_number,omitempty"`
	RecordedDate   string   `json:"recorded_date,omitempty"`
}

// Permit is a matched building permit.
type Permit struct {
	PermitID       string   `json:"permit_id"`
	ParcelID       string   `json:"parcel_id"`
	StateFIPS      string   `json:"state_fips"`
	CountyFIPS     string   `json:"county_fips"`
	PermitNumber   string   `json:"permit_number,omitempty"`
	PermitType     string   `json:"permit_type,omitempty"`
	PermitStatus   string   `json:"permit_status,omitempty"`
	FiledDate      string   `json:"filed_date,omitempty"`
	IssuedDate     string   `json:"issued_date,omitempty"`
	Description    string   `json:"description,omitempty"`
	EstimatedCost  *float64 `json:"estimated_cost,omitempty"`
	ContractorName string   `json:"contractor_name,omitempty"`
}

// Webhook is a registered outbound webhook endpoint.
type Webhook struct {
	ID          string   `json:"id"`
	URL         string   `json:"url"`
	EventTypes  []string `json:"event_types"`
	FilterKind  string   `json:"filter_kind"`
	FilterValue any      `json:"filter_value"`
	IsActive    bool     `json:"is_active"`
	CreatedAt   string   `json:"created_at"`
}

// HealthStatus is the response from /v1/health.
type HealthStatus struct {
	Status        string `json:"status"`
	APIVersion    string `json:"api_version,omitempty"`
	SpecRevision  string `json:"spec_revision,omitempty"`
}
