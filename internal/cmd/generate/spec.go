package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// node is an order-preserving JSON value, so generated code follows the
// spec's property order and output is deterministic.
type node struct {
	kind  byte // 'o' object, 'a' array, 's' string, 'n' number, 'b' bool, 'z' null
	keys  []string
	m     map[string]*node
	items []*node
	str   string // string value, or the number's text
	b     bool
}

func parseJSON(data []byte) (*node, error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	n, err := parseValue(dec)
	if err != nil {
		return nil, err
	}
	if _, err := dec.Token(); err != io.EOF {
		return nil, fmt.Errorf("trailing data after JSON value")
	}
	return n, nil
}

func parseValue(dec *json.Decoder) (*node, error) {
	tok, err := dec.Token()
	if err != nil {
		return nil, err
	}
	switch t := tok.(type) {
	case json.Delim:
		switch t {
		case '{':
			n := &node{kind: 'o', m: map[string]*node{}}
			for dec.More() {
				kt, err := dec.Token()
				if err != nil {
					return nil, err
				}
				k, ok := kt.(string)
				if !ok {
					return nil, fmt.Errorf("object key is not a string")
				}
				v, err := parseValue(dec)
				if err != nil {
					return nil, err
				}
				if _, dup := n.m[k]; !dup {
					n.keys = append(n.keys, k)
				}
				n.m[k] = v
			}
			if _, err := dec.Token(); err != nil {
				return nil, err
			}
			return n, nil
		case '[':
			n := &node{kind: 'a'}
			for dec.More() {
				v, err := parseValue(dec)
				if err != nil {
					return nil, err
				}
				n.items = append(n.items, v)
			}
			if _, err := dec.Token(); err != nil {
				return nil, err
			}
			return n, nil
		}
		return nil, fmt.Errorf("unexpected delimiter %v", t)
	case string:
		return &node{kind: 's', str: t}, nil
	case json.Number:
		return &node{kind: 'n', str: t.String()}, nil
	case bool:
		return &node{kind: 'b', b: t}, nil
	case nil:
		return &node{kind: 'z'}, nil
	}
	return nil, fmt.Errorf("unexpected token %v", tok)
}

func (n *node) get(k string) *node {
	if n == nil || n.kind != 'o' {
		return nil
	}
	return n.m[k]
}

func (n *node) has(k string) bool { return n.get(k) != nil }

func (n *node) s(k string) string {
	if v := n.get(k); v != nil && v.kind == 's' {
		return v.str
	}
	return ""
}

func (n *node) truthy(k string) bool {
	v := n.get(k)
	return v != nil && v.kind == 'b' && v.b
}

func (n *node) isObject() bool { return n != nil && n.kind == 'o' }

// text renders a scalar node for doc comments.
func (n *node) text() string {
	if n == nil {
		return ""
	}
	switch n.kind {
	case 's':
		return n.str
	case 'n':
		return n.str
	case 'b':
		if n.b {
			return "true"
		}
		return "false"
	case 'z':
		return "null"
	}
	return ""
}

// resolvePointer follows a local JSON pointer ("#/components/schemas/X").
func resolvePointer(root *node, ref string) (*node, error) {
	if !strings.HasPrefix(ref, "#/") {
		return nil, fmt.Errorf("unsupported $ref %q (only local refs)", ref)
	}
	cur := root
	for _, seg := range strings.Split(ref[2:], "/") {
		seg = strings.ReplaceAll(strings.ReplaceAll(seg, "~1", "/"), "~0", "~")
		switch {
		case cur.isObject():
			cur = cur.m[seg]
		default:
			cur = nil
		}
		if cur == nil {
			return nil, fmt.Errorf("unresolvable $ref %q", ref)
		}
	}
	return cur, nil
}
