package propraven

// String returns a pointer to v, for optional string fields.
func String(v string) *string { return &v }

// Int returns a pointer to v, for optional integer fields.
func Int(v int64) *int64 { return &v }

// Float returns a pointer to v, for optional number fields.
func Float(v float64) *float64 { return &v }

// Bool returns a pointer to v, for optional boolean fields.
func Bool(v bool) *bool { return &v }

// Ptr returns a pointer to any value (for optional struct fields).
func Ptr[T any](v T) *T { return &v }
