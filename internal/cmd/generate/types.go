package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jdw2111/propraven-go/internal/naming"
)

type kind int

const (
	kScalar kind = iota // string, int64, float64, bool, enum aliases
	kStruct             // generated structs, unions
	kSlice
	kMap
	kRaw // json.RawMessage
	kAny // any
)

type goType struct {
	expr     string
	kind     kind
	nullable bool
}

var rawType = goType{expr: "json.RawMessage", kind: kRaw}

// structField describes one field of a generated struct.
type structField struct {
	name string
	wire string
	t    goType
	doc  string
	tag  string
	expr string
}

// typeOf maps a schema to a Go type, declaring named types (structs, unions,
// enums) as needed. hint names any type it has to declare.
func (g *gen) typeOf(s *node, hint string) goType {
	if s == nil {
		return rawType
	}
	if s.kind == 'b' { // boolean schema: true = anything
		return rawType
	}
	if !s.isObject() {
		return rawType
	}
	if ref := s.s("$ref"); ref != "" {
		if name, ok := componentSchemaName(ref); ok && g.compSchemas.has(name) {
			return g.compType(name)
		}
		target, err := resolvePointer(g.root, ref)
		if err != nil {
			g.fail(err)
			return rawType
		}
		return g.typeOf(target, hint)
	}
	types, nullable := schemaTypes(s)
	desc := s.s("description")

	for _, key := range []string{"oneOf", "anyOf"} {
		alts := s.get(key)
		if alts == nil || alts.kind != 'a' {
			continue
		}
		var variants []*node
		for _, v := range alts.items {
			if g.isNullSchema(v) {
				nullable = true
				continue
			}
			variants = append(variants, v)
		}
		if len(variants) == 0 {
			return goType{expr: "json.RawMessage", kind: kRaw, nullable: nullable}
		}
		if len(variants) == 1 && !s.has("properties") {
			t := g.typeOf(variants[0], hint)
			t.nullable = t.nullable || nullable
			return t
		}
		if t, ok := g.declareUnion(hint, variants, desc); ok {
			t.nullable = nullable
			return t
		}
		return goType{expr: "json.RawMessage", kind: kRaw, nullable: nullable}
	}

	if all := s.get("allOf"); all != nil && all.kind == 'a' {
		if len(all.items) == 1 && !s.has("properties") {
			t := g.typeOf(all.items[0], hint)
			t.nullable = t.nullable || nullable
			return t
		}
		props, req, ok := g.mergeAllOf(s)
		if !ok {
			return goType{expr: "json.RawMessage", kind: kRaw, nullable: nullable}
		}
		if len(props) == 0 {
			return goType{expr: "map[string]any", kind: kMap, nullable: nullable}
		}
		t := g.declareStruct(hint, props, req, desc)
		t.nullable = nullable
		return t
	}

	if len(types) > 1 {
		return goType{expr: "json.RawMessage", kind: kRaw, nullable: nullable}
	}
	t := ""
	if len(types) == 1 {
		t = types[0]
	}
	if enum := s.get("enum"); enum != nil && enum.kind == 'a' && (t == "" || t == "string") && allStrings(enum) {
		e := g.declareEnum(hint, enum, desc)
		e.nullable = nullable
		return e
	}
	if t == "" {
		switch {
		case s.has("properties") || s.has("additionalProperties"):
			t = "object"
		case s.has("items"):
			t = "array"
		case s.has("const"):
			switch s.get("const").kind {
			case 's':
				t = "string"
			case 'n':
				t = "number"
			case 'b':
				t = "boolean"
			}
		case s.has("enum"):
			t = "string"
		}
	}
	switch t {
	case "string":
		return goType{expr: "string", kind: kScalar, nullable: nullable}
	case "integer":
		return goType{expr: "int64", kind: kScalar, nullable: nullable}
	case "number":
		return goType{expr: "float64", kind: kScalar, nullable: nullable}
	case "boolean":
		return goType{expr: "bool", kind: kScalar, nullable: nullable}
	case "null":
		return goType{expr: "any", kind: kAny, nullable: true}
	case "array":
		elemHint := hint
		if g.preclaimed[hint] {
			// The name is reserved for the array itself (an alias).
			elemHint = hint + "Item"
		}
		elem := g.typeOf(s.get("items"), elemHint)
		return goType{expr: "[]" + elemExpr(elem), kind: kSlice, nullable: nullable}
	case "object":
		if props := s.get("properties"); props.isObject() && len(props.keys) > 0 {
			var ps []prop
			for _, k := range props.keys {
				ps = append(ps, prop{name: k, schema: props.m[k]})
			}
			st := g.declareStruct(hint, ps, requiredSet(s), desc)
			st.nullable = nullable
			return st
		}
		if ap := s.get("additionalProperties"); ap.isObject() && len(ap.keys) > 0 {
			v := g.typeOf(ap, hint+"Value")
			return goType{expr: "map[string]" + elemExpr(v), kind: kMap, nullable: nullable}
		}
		return goType{expr: "map[string]any", kind: kMap, nullable: nullable}
	}
	return goType{expr: "json.RawMessage", kind: kRaw, nullable: nullable}
}

// elemExpr is the type of a slice element or map value.
func elemExpr(t goType) string {
	if t.nullable && (t.kind == kScalar || t.kind == kStruct) {
		return "*" + t.expr
	}
	return t.expr
}

type prop struct {
	name   string
	schema *node
}

func schemaTypes(s *node) ([]string, bool) {
	tn := s.get("type")
	nullable := s.truthy("nullable") // OpenAPI 3.0 style
	var out []string
	switch {
	case tn == nil:
	case tn.kind == 's':
		if tn.str == "null" {
			return []string{"null"}, true
		}
		out = append(out, tn.str)
	case tn.kind == 'a':
		for _, it := range tn.items {
			if it.kind != 's' {
				continue
			}
			if it.str == "null" {
				nullable = true
				continue
			}
			out = append(out, it.str)
		}
		if len(out) == 0 && nullable {
			return []string{"null"}, true
		}
	}
	return out, nullable
}

func (g *gen) isNullSchema(s *node) bool {
	if !s.isObject() {
		return false
	}
	if ref := s.s("$ref"); ref != "" {
		t, err := resolvePointer(g.root, ref)
		if err != nil {
			return false
		}
		s = t
	}
	types, _ := schemaTypes(s)
	return len(types) == 1 && types[0] == "null"
}

func allStrings(a *node) bool {
	if len(a.items) == 0 {
		return false
	}
	for _, it := range a.items {
		if it.kind != 's' {
			return false
		}
	}
	return true
}

func requiredSet(s *node) map[string]bool {
	out := map[string]bool{}
	if r := s.get("required"); r != nil && r.kind == 'a' {
		for _, it := range r.items {
			if it.kind == 's' {
				out[it.str] = true
			}
		}
	}
	return out
}

func componentSchemaName(ref string) (string, bool) {
	const p = "#/components/schemas/"
	if strings.HasPrefix(ref, p) && !strings.Contains(ref[len(p):], "/") {
		return strings.ReplaceAll(strings.ReplaceAll(ref[len(p):], "~1", "/"), "~0", "~"), true
	}
	return "", false
}

// mergeAllOf flattens allOf members (objects and refs to objects) into one
// property list. ok is false when a member is not object-like.
func (g *gen) mergeAllOf(s *node) ([]prop, map[string]bool, bool) {
	var props []prop
	idx := map[string]int{}
	req := map[string]bool{}
	var add func(n *node, depth int) bool
	add = func(n *node, depth int) bool {
		if depth > 20 || !n.isObject() {
			return false
		}
		if ref := n.s("$ref"); ref != "" {
			t, err := resolvePointer(g.root, ref)
			if err != nil {
				return false
			}
			return add(t, depth+1)
		}
		if all := n.get("allOf"); all != nil && all.kind == 'a' {
			for _, m := range all.items {
				if !add(m, depth+1) {
					return false
				}
			}
		}
		if n.has("oneOf") || n.has("anyOf") {
			return false
		}
		types, _ := schemaTypes(n)
		if len(types) > 0 && types[0] != "object" {
			return false
		}
		if p := n.get("properties"); p.isObject() {
			for _, k := range p.keys {
				if i, ok := idx[k]; ok {
					props[i].schema = p.m[k]
					continue
				}
				idx[k] = len(props)
				props = append(props, prop{name: k, schema: p.m[k]})
			}
		}
		for k := range requiredSet(n) {
			req[k] = true
		}
		return true
	}
	for _, m := range s.get("allOf").items {
		if !add(m, 0) {
			return nil, nil, false
		}
	}
	if p := s.get("properties"); p.isObject() {
		for _, k := range p.keys {
			if i, ok := idx[k]; ok {
				props[i].schema = p.m[k]
				continue
			}
			idx[k] = len(props)
			props = append(props, prop{name: k, schema: p.m[k]})
		}
	}
	for k := range requiredSet(s) {
		req[k] = true
	}
	return props, req, true
}

// fieldDecl decides a struct field's Go type and json tag.
func fieldDecl(t goType, required bool, owner string) (expr, tagOpts string) {
	optional := !required || t.nullable
	expr = t.expr
	switch t.kind {
	case kScalar, kStruct:
		if optional || expr == owner {
			expr = "*" + expr
		}
	}
	if optional {
		tagOpts = ",omitempty"
	}
	return expr, tagOpts
}

func (g *gen) buildFields(owner string, props []prop, req map[string]bool) []structField {
	var fields []structField
	used := map[string]bool{}
	for _, p := range props {
		t := g.typeOf(p.schema, owner+naming.Pascal(p.name))
		fname := naming.Pascal(p.name)
		for i := 2; used[fname]; i++ {
			fname = naming.Pascal(p.name) + fmt.Sprint(i)
		}
		used[fname] = true
		expr, opts := fieldDecl(t, req[p.name], owner)
		doc := schemaDoc(p.schema)
		fields = append(fields, structField{
			name: fname, wire: p.name, t: t, doc: doc, expr: expr,
			tag: fmt.Sprintf("`json:%q`", p.name+opts),
		})
	}
	return fields
}

func (g *gen) declareStruct(hint string, props []prop, req map[string]bool, desc string) goType {
	name := g.claim(hint)
	slot := g.reserve()
	fields := g.buildFields(name, props, req)
	g.recordFields(name, fields)
	var b strings.Builder
	writeDoc(&b, "", name, g.docOr(name, desc), "")
	fmt.Fprintf(&b, "type %s struct {\n", name)
	for i, f := range fields {
		if f.doc != "" {
			if i > 0 {
				b.WriteString("\n")
			}
			writeComment(&b, "\t", f.doc)
		}
		fmt.Fprintf(&b, "\t%s %s %s\n", f.name, f.expr, f.tag)
	}
	b.WriteString("}\n")
	writeLenientUnmarshal(&b, name, fields)
	g.fill(slot, b.String())
	return goType{expr: name, kind: kStruct}
}

// writeLenientUnmarshal gives structs with numeric fields an UnmarshalJSON
// that accepts each number as a JSON number or a quoted decimal string
// ("19388200.00"), so decoding works against old and new server versions.
// A field whose JSON type does not match the spec is left unset.
func writeLenientUnmarshal(b *strings.Builder, name string, fields []structField) {
	type nf struct {
		f        structField
		helper   string
		assignFn string
	}
	var nums []nf
	for _, f := range fields {
		base := strings.TrimPrefix(f.expr, "*")
		ptr := strings.HasPrefix(f.expr, "*")
		switch base {
		case "float64", "int64":
			fn := "assign"
			if ptr {
				fn = "assignPtr"
			}
			nums = append(nums, nf{f, "lenientNumber[" + base + "]", fn})
		case "[]float64", "[]int64":
			nums = append(nums, nf{f, "lenientSlice[" + base[2:] + "]", "assign"})
		case "map[string]float64", "map[string]int64":
			nums = append(nums, nf{f, "lenientMap[" + base[len("map[string]"):] + "]", "assign"})
		}
	}
	if len(nums) == 0 {
		return
	}
	fmt.Fprintf(b, "\n// UnmarshalJSON decodes %s, accepting numeric fields sent as JSON numbers or as\n// quoted decimal strings.\n", name)
	fmt.Fprintf(b, "func (r *%s) UnmarshalJSON(data []byte) error {\n\ttype plain %s\n\taux := struct {\n\t\t*plain\n", name, name)
	for _, n := range nums {
		fmt.Fprintf(b, "\t\t%s %s `json:%q`\n", n.f.name, n.helper, n.f.wire)
	}
	b.WriteString("\t}{plain: (*plain)(r)}\n\terr := json.Unmarshal(data, &aux)\n")
	for _, n := range nums {
		fmt.Fprintf(b, "\taux.%s.%s(&r.%s)\n", n.f.name, n.assignFn, n.f.name)
	}
	b.WriteString("\treturn softTypeError(err)\n}\n")
}

func (g *gen) recordFields(name string, fields []structField) {
	m := map[string]goType{}
	for _, f := range fields {
		m[f.wire] = f.t
	}
	g.structFields[name] = m
}

// jsonKind classifies the JSON value a schema produces, or "" if unknown.
func (g *gen) jsonKind(s *node, depth int) string {
	if depth > 20 || !s.isObject() {
		return ""
	}
	if ref := s.s("$ref"); ref != "" {
		t, err := resolvePointer(g.root, ref)
		if err != nil {
			return ""
		}
		return g.jsonKind(t, depth+1)
	}
	if s.has("oneOf") || s.has("anyOf") {
		return ""
	}
	if all := s.get("allOf"); all != nil && all.kind == 'a' && len(all.items) > 0 {
		return g.jsonKind(all.items[0], depth+1)
	}
	types, _ := schemaTypes(s)
	if len(types) > 1 {
		return ""
	}
	if len(types) == 1 {
		switch types[0] {
		case "integer", "number":
			return "number"
		case "null":
			return ""
		}
		return types[0]
	}
	switch {
	case s.has("properties") || s.has("additionalProperties"):
		return "object"
	case s.has("items"):
		return "array"
	case s.has("enum") && allStrings(s.get("enum")):
		return "string"
	}
	return ""
}

var unionFieldNames = map[string]string{
	"array": "Array", "object": "Object", "string": "String", "number": "Number", "boolean": "Bool",
}
var unionKindBytes = map[string]string{
	"array": "'a'", "object": "'o'", "string": "'s'", "number": "'n'", "boolean": "'b'",
}

// declareUnion handles oneOf/anyOf whose variants are distinguishable by
// JSON kind (e.g. a bare array vs an envelope object).
func (g *gen) declareUnion(hint string, variants []*node, desc string) (goType, bool) {
	seen := map[string]bool{}
	var kinds []string
	for _, v := range variants {
		k := g.jsonKind(v, 0)
		if k == "" || seen[k] {
			return goType{}, false
		}
		seen[k] = true
		kinds = append(kinds, k)
	}
	name := g.claim(hint)
	slot := g.reserve()
	type uv struct {
		field, kind, expr string
		t                 goType
	}
	var vs []uv
	for i, v := range variants {
		f := unionFieldNames[kinds[i]]
		t := g.typeOf(v, name+f)
		expr := t.expr
		if t.kind == kScalar || t.kind == kStruct {
			expr = "*" + expr
		}
		vs = append(vs, uv{field: f, kind: kinds[i], expr: expr, t: t})
	}
	var b strings.Builder
	var shapes []string
	for _, v := range vs {
		shapes = append(shapes, fmt.Sprintf("%s (%s)", strings.TrimPrefix(v.expr, "*"), v.field))
	}
	writeDoc(&b, "", name, g.docOr(name, desc), "It is one of: "+strings.Join(shapes, ", ")+". Exactly one variant field is set after decoding; Raw keeps the original JSON.")
	fmt.Fprintf(&b, "type %s struct {\n", name)
	for _, v := range vs {
		fmt.Fprintf(&b, "\t// %s is set when the value is a JSON %s.\n\t%s %s\n", v.field, v.kind, v.field, v.expr)
	}
	b.WriteString("\t// Raw is the undecoded JSON value.\n\tRaw json.RawMessage\n}\n\n")
	fmt.Fprintf(&b, "// UnmarshalJSON decodes whichever variant matches the JSON kind.\nfunc (u *%s) UnmarshalJSON(b []byte) error {\n", name)
	fmt.Fprintf(&b, "\t*u = %s{Raw: append(json.RawMessage(nil), b...)}\n\tswitch jsonKindOf(b) {\n", name)
	for _, v := range vs {
		fmt.Fprintf(&b, "\tcase %s:\n\t\treturn json.Unmarshal(b, &u.%s)\n", unionKindBytes[v.kind], v.field)
	}
	b.WriteString("\t}\n\treturn nil\n}\n\n")
	fmt.Fprintf(&b, "// MarshalJSON encodes the set variant (or Raw).\nfunc (u %s) MarshalJSON() ([]byte, error) {\n\tswitch {\n", name)
	for _, v := range vs {
		fmt.Fprintf(&b, "\tcase u.%s != nil:\n\t\treturn json.Marshal(u.%s)\n", v.field, v.field)
	}
	b.WriteString("\tcase u.Raw != nil:\n\t\treturn u.Raw, nil\n\t}\n\treturn []byte(\"null\"), nil\n}\n")
	g.fill(slot, b.String())
	return goType{expr: name, kind: kStruct}, true
}

func (g *gen) declareEnum(hint string, enum *node, desc string) goType {
	name := g.claim(hint)
	slot := g.reserve()
	var b strings.Builder
	writeDoc(&b, "", name, g.docOr(name, desc), "It is a string; the "+name+"* constants list the documented values.")
	fmt.Fprintf(&b, "type %s = string\n\n// Documented values of %s.\nconst (\n", name, name)
	seen := map[string]bool{}
	for _, it := range enum.items {
		if seen[it.str] {
			continue
		}
		seen[it.str] = true
		c := g.unique(name + naming.Pascal(it.str))
		fmt.Fprintf(&b, "\t%s %s = %q\n", c, name, it.str)
	}
	b.WriteString(")\n")
	g.fill(slot, b.String())
	return goType{expr: name, kind: kScalar}
}

// compType returns (declaring on first use) the Go type of a component schema.
func (g *gen) compType(name string) goType {
	if t, ok := g.compTypes[name]; ok {
		return t
	}
	goName := g.compNames[name]
	if g.compBusy[name] {
		return goType{expr: goName, kind: kStruct}
	}
	g.compBusy[name] = true
	t := g.declareNamed(goName, g.compSchemas.m[name], schemaDoc(g.compSchemas.m[name]))
	g.compTypes[name] = t
	return t
}

// declareNamed makes goName (already reserved) a type for schema s: a
// struct/union/enum declaration, or an alias of whatever s maps to.
func (g *gen) declareNamed(goName string, s *node, doc string) goType {
	g.preclaimed[goName] = true
	g.fallbackDoc[goName] = doc
	t := g.typeOf(s, goName)
	if t.expr == goName {
		return t
	}
	delete(g.preclaimed, goName)
	var b strings.Builder
	writeDoc(&b, "", goName, doc, "")
	fmt.Fprintf(&b, "type %s = %s\n", goName, t.expr)
	g.fill(g.reserve(), b.String())
	if _, isNamed := g.structFields[t.expr]; isNamed {
		g.aliasOf[goName] = t.expr
	} else if target, ok := g.aliasOf[t.expr]; ok {
		g.aliasOf[goName] = target
	}
	return goType{expr: goName, kind: t.kind, nullable: t.nullable}
}

// claim returns hint itself when it was reserved for this declaration,
// otherwise a unique variant of it.
func (g *gen) claim(hint string) string {
	if g.preclaimed[hint] {
		delete(g.preclaimed, hint)
		return hint
	}
	return g.unique(hint)
}

// docOr returns desc, or the fallback doc registered for a reserved name
// (an operation's summary for its response type).
func (g *gen) docOr(name, desc string) string {
	if strings.TrimSpace(desc) != "" {
		return desc
	}
	return g.fallbackDoc[name]
}

func (g *gen) unique(base string) string {
	if !g.used[base] {
		g.used[base] = true
		return base
	}
	for i := 2; ; i++ {
		n := fmt.Sprintf("%s%d", base, i)
		if !g.used[n] {
			g.used[n] = true
			return n
		}
	}
}

func sortedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
