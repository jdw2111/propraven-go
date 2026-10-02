package propraven

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// Generated structs decode their numeric fields through these helpers so a
// number arrives intact whether the server sends 19388200 or "19388200.00".
// Older API deployments serialised some decimals as strings. A value that is
// not numeric at all (say "NC" where the spec promises a number) leaves the
// field unset instead of failing the whole response; the raw body is still
// available through WithRawResponse.

type number interface{ int64 | float64 }

func parseLenient[T number](b []byte) (v T, isNull bool, err error) {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || string(b) == "null" {
		return v, true, nil
	}
	s := string(b)
	if b[0] == '"' {
		if err := json.Unmarshal(b, &s); err != nil {
			return v, false, err
		}
		s = strings.TrimSpace(s)
		if s == "" {
			return v, true, nil
		}
	}
	var zero T
	switch any(zero).(type) {
	case int64:
		if n, err := strconv.ParseInt(s, 10, 64); err == nil {
			return T(n), false, nil
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil || math.IsNaN(f) || math.IsInf(f, 0) || f != math.Trunc(f) {
			return v, false, fmt.Errorf("propraven: %q is not an integer", s)
		}
		return T(int64(f)), false, nil
	default:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return v, false, fmt.Errorf("propraven: %q is not a number", s)
		}
		return T(f), false, nil
	}
}

type lenientNumber[T number] struct {
	v          T
	seen, null bool
}

func (l *lenientNumber[T]) UnmarshalJSON(b []byte) error {
	v, isNull, err := parseLenient[T](b)
	if err != nil {
		// Not numeric: leave the destination field untouched.
		*l = lenientNumber[T]{}
		return nil
	}
	l.v, l.null, l.seen = v, isNull, true
	return nil
}

func (l lenientNumber[T]) assign(dst *T) {
	if l.seen && !l.null {
		*dst = l.v
	}
}

func (l lenientNumber[T]) assignPtr(dst **T) {
	if !l.seen {
		return
	}
	if l.null {
		*dst = nil
		return
	}
	v := l.v
	*dst = &v
}

type lenientSlice[T number] struct {
	v    []T
	seen bool
}

func (l *lenientSlice[T]) UnmarshalJSON(b []byte) error {
	l.seen = true
	if string(bytes.TrimSpace(b)) == "null" {
		l.v = nil
		return nil
	}
	var raw []json.RawMessage
	if err := json.Unmarshal(b, &raw); err != nil {
		*l = lenientSlice[T]{}
		return nil
	}
	l.v = make([]T, len(raw))
	for i, r := range raw {
		v, _, _ := parseLenient[T](r)
		l.v[i] = v
	}
	return nil
}

func (l lenientSlice[T]) assign(dst *[]T) {
	if l.seen {
		*dst = l.v
	}
}

type lenientMap[T number] struct {
	v    map[string]T
	seen bool
}

func (l *lenientMap[T]) UnmarshalJSON(b []byte) error {
	l.seen = true
	if string(bytes.TrimSpace(b)) == "null" {
		l.v = nil
		return nil
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err != nil {
		*l = lenientMap[T]{}
		return nil
	}
	l.v = make(map[string]T, len(raw))
	for k, r := range raw {
		v, _, _ := parseLenient[T](r)
		l.v[k] = v
	}
	return nil
}

func (l lenientMap[T]) assign(dst *map[string]T) {
	if l.seen {
		*dst = l.v
	}
}

// softTypeError drops *json.UnmarshalTypeError: encoding/json has already
// decoded every other field, and one field whose JSON type drifted from the
// spec should not fail the whole response.
func softTypeError(err error) error {
	var te *json.UnmarshalTypeError
	if errors.As(err, &te) {
		return nil
	}
	return err
}

// jsonKindOf classifies a JSON value by its first byte: 'o' object, 'a'
// array, 's' string, 'n' number, 'b' boolean, 'z' null, 0 unknown.
func jsonKindOf(b []byte) byte {
	b = bytes.TrimSpace(b)
	if len(b) == 0 {
		return 0
	}
	switch c := b[0]; {
	case c == '{':
		return 'o'
	case c == '[':
		return 'a'
	case c == '"':
		return 's'
	case c == 't' || c == 'f':
		return 'b'
	case c == 'n':
		return 'z'
	case c == '-' || (c >= '0' && c <= '9'):
		return 'n'
	}
	return 0
}

// decodeMixed handles operations that answer either JSON or another media
// type (for example CSV): JSON goes through decode, anything else to text.
func decodeMixed(contentType, text *string, decode func([]byte) error) decodeFunc {
	return func(h http.Header, body []byte) error {
		*contentType = h.Get("Content-Type")
		if isJSONContentType(h) {
			if len(bytes.TrimSpace(body)) == 0 {
				return nil
			}
			if err := decode(body); err != nil {
				return fmt.Errorf("propraven: decoding response: %w", err)
			}
			return nil
		}
		*text = string(body)
		return nil
	}
}
