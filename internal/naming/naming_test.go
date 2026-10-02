package naming

import "testing"

func TestPascal(t *testing.T) {
	cases := map[string]string{
		"parcel_id":        "ParcelID",
		"county_fips":      "CountyFIPS",
		"trafficHistory":   "TrafficHistory",
		"highLandRatio":    "HighLandRatio",
		"x402Version":      "X402Version",
		"nextCursor":       "NextCursor",
		"has_more":         "HasMore",
		"10-15yr":          "V1015yr",
		"30yr+":            "V30yrPlus",
		"VERY_EXPENSIVE":   "VeryExpensive",
		"parcel.sold":      "ParcelSold",
		"cmbs":             "CMBS",
		"retryDelivery":    "RetryDelivery",
		"zoningCategories": "ZoningCategories",
		"HTTPServer":       "HTTPServer",
		"canonical_ids":    "CanonicalIDs",
		"":                 "Empty",
	}
	for in, want := range cases {
		if got := Pascal(in); got != want {
			t.Errorf("Pascal(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestCamelAndHeaders(t *testing.T) {
	if got := Camel("deliveryId"); got != "deliveryID" {
		t.Errorf("Camel(deliveryId) = %q", got)
	}
	if got := Camel("type"); got != "type_" {
		t.Errorf("Camel(type) = %q", got)
	}
	if got := Camel("id"); got != "id" {
		t.Errorf("Camel(id) = %q", got)
	}
	if got := HeaderField("X-CREDIT-TOKEN"); got != "CreditToken" {
		t.Errorf("HeaderField = %q", got)
	}
	if got := HeaderField("X-PAYMENT"); got != "Payment" {
		t.Errorf("HeaderField = %q", got)
	}
}
