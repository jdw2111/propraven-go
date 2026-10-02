package main

import (
	"errors"
	"fmt"
	"go/format"
	"regexp"
	"sort"
	"strings"

	"github.com/jdw2111/propraven-go/internal/naming"
)

const header = "// Code generated from openapi.json by internal/cmd/generate. DO NOT EDIT.\n\n"

// coreIdentifiers are package-level names of the hand-written core.
var coreIdentifiers = []string{
	"Version", "DefaultBaseURL", "Option", "RequestOption", "WithAPIKey", "WithBaseURL",
	"WithHTTPClient", "WithMaxRetries", "WithTimeout", "WithHeader", "WithLogger",
	"RawResponse", "WithRawResponse", "RateLimit", "NewClient", "Client", "Error",
	"FieldError", "ConnectionError", "AsError", "IsBadRequest", "IsAuthentication",
	"IsPaymentRequired", "IsPermissionDenied", "IsNotFound", "IsConflict", "IsRateLimited",
	"IsServerError", "IsConnectionError", "IsTimeout", "IterOptions", "Iter",
	"WebhookSignatureHeader", "DefaultWebhookTolerance", "ErrWebhookSignature",
	"WebhookOption", "WithWebhookTolerance", "WithWebhookNow", "VerifyWebhook",
	"SignWebhook", "String", "Int", "Float", "Bool", "Ptr",
}

// Result is the generator's output.
type Result struct {
	Files       map[string][]byte
	MethodTable string
	Operations  int
	Warnings    []string
}

type fileOut struct {
	decls []string
}

type gen struct {
	root        *node
	compSchemas *node
	compNames   map[string]string // schema name -> Go name
	compTypes   map[string]goType
	compBusy    map[string]bool

	used         map[string]bool
	preclaimed   map[string]bool
	structFields map[string]map[string]goType
	aliasOf      map[string]string
	fallbackDoc  map[string]string

	cur      *fileOut
	errs     []error
	warnings []string
}

func (g *gen) fail(err error) { g.errs = append(g.errs, err) }

func (g *gen) warn(format string, args ...any) {
	g.warnings = append(g.warnings, fmt.Sprintf(format, args...))
}

func (g *gen) reserve() int {
	g.cur.decls = append(g.cur.decls, "")
	return len(g.cur.decls) - 1
}

func (g *gen) fill(slot int, code string) { g.cur.decls[slot] = code }

type param struct {
	name     string
	in       string
	required bool
	schema   *node
	desc     string
	explode  bool
	deprec   bool
}

type operation struct {
	group, method    string
	httpMethod, path string
	op               *node
	params           []param
	summary          string
}

type paramsField struct {
	name, wire, in string // in: query, header, body
	expr, tag, doc string
	t              goType
	pointer        bool
	list           bool
	explode        bool
}

var httpMethods = []string{"get", "put", "post", "delete", "options", "head", "patch", "trace"}

// Generate produces the generated files for spec.
func Generate(spec []byte, pkg string) (*Result, error) {
	root, err := parseJSON(spec)
	if err != nil {
		return nil, fmt.Errorf("parsing spec: %w", err)
	}
	g := &gen{
		root:         root,
		compSchemas:  root.get("components").get("schemas"),
		compNames:    map[string]string{},
		compTypes:    map[string]goType{},
		compBusy:     map[string]bool{},
		used:         map[string]bool{},
		preclaimed:   map[string]bool{},
		structFields: map[string]map[string]goType{},
		aliasOf:      map[string]string{},
		fallbackDoc:  map[string]string{},
	}
	if g.compSchemas == nil {
		g.compSchemas = &node{kind: 'o', m: map[string]*node{}}
	}
	for _, id := range coreIdentifiers {
		g.used[id] = true
	}

	ops, err := g.collectOperations()
	if err != nil {
		return nil, err
	}

	// Reserve names: components first, then services, params and responses.
	for _, name := range g.compSchemas.keys {
		base := naming.Pascal(name)
		goName := base
		if g.used[goName] {
			goName = base + "Schema"
		}
		g.compNames[name] = g.unique(goName)
	}
	var groups []string
	byGroup := map[string][]*operation{}
	for _, op := range ops {
		if _, ok := byGroup[op.group]; !ok {
			groups = append(groups, op.group)
			g.unique(naming.Service(op.group))
		}
		byGroup[op.group] = append(byGroup[op.group], op)
	}
	opNames := map[*operation][2]string{}
	for _, op := range ops {
		base := naming.Pascal(op.group) + naming.Method(op.method)
		opNames[op] = [2]string{g.unique(base + "Params"), g.unique(base + "Response")}
	}

	files := map[string]*fileOut{}
	// Components.
	types := &fileOut{}
	files["gen_types.go"] = types
	g.cur = types
	for _, name := range g.compSchemas.keys {
		g.compType(name)
	}

	var table strings.Builder
	sortedGroups := append([]string(nil), groups...)
	sort.Strings(sortedGroups)
	for _, grp := range groups {
		f := &fileOut{}
		fname := "gen_" + strings.ToLower(naming.Pascal(grp))
		if fname == "gen_client" || fname == "gen_types" {
			fname += "_service"
		}
		files[fname+".go"] = f
		g.cur = f
		svc := naming.Service(grp)
		f.decls = append(f.decls, fmt.Sprintf("// %s groups the %s operations. Use it as client.%s.\ntype %s struct {\n\tclient *Client\n}\n",
			svc, grp, naming.Field(grp), svc))
		fmt.Fprintf(&table, "\n### client.%s\n\n| Method | HTTP | Summary |\n|---|---|---|\n", naming.Field(grp))
		for _, op := range byGroup[grp] {
			names := opNames[op]
			rows := g.genOperation(op, names[0], names[1])
			table.WriteString(rows)
		}
	}

	// Client.
	var cb strings.Builder
	cb.WriteString("// Client is the PropRaven API client. Create it with [NewClient]; each\n// field groups the operations of one API namespace (x-sdk-group).\ntype Client struct {\n")
	for _, grp := range sortedGroups {
		fmt.Fprintf(&cb, "\t// %s holds the %s operations.\n\t%s *%s\n", naming.Field(grp), grp, naming.Field(grp), naming.Service(grp))
	}
	cb.WriteString("\n\tcore *clientCore\n}\n\nfunc (c *Client) initServices() {\n")
	for _, grp := range sortedGroups {
		fmt.Fprintf(&cb, "\tc.%s = &%s{client: c}\n", naming.Field(grp), naming.Service(grp))
	}
	cb.WriteString("}\n")
	files["gen_client.go"] = &fileOut{decls: []string{cb.String()}}

	if len(g.errs) > 0 {
		return nil, errors.Join(g.errs...)
	}

	out := map[string][]byte{}
	for name, f := range files {
		body := strings.Join(nonEmpty(f.decls), "\n")
		var imports []string
		if strings.Contains(body, "context.") {
			imports = append(imports, `"context"`)
		}
		if strings.Contains(body, "json.") {
			imports = append(imports, `"encoding/json"`)
		}
		src := header + "package " + pkg + "\n\n"
		if len(imports) > 0 {
			src += "import (\n\t" + strings.Join(imports, "\n\t") + "\n)\n\n"
		}
		src += body
		formatted, err := format.Source([]byte(src))
		if err != nil {
			return nil, fmt.Errorf("formatting %s: %w\n%s", name, err, numbered(src))
		}
		out[name] = formatted
	}
	return &Result{Files: out, MethodTable: table.String(), Operations: len(ops), Warnings: g.warnings}, nil
}

func nonEmpty(ss []string) []string {
	var out []string
	for _, s := range ss {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func numbered(src string) string {
	var b strings.Builder
	for i, l := range strings.Split(src, "\n") {
		fmt.Fprintf(&b, "%5d %s\n", i+1, l)
	}
	return b.String()
}

func (g *gen) collectOperations() ([]*operation, error) {
	paths := g.root.get("paths")
	if !paths.isObject() {
		return nil, fmt.Errorf("spec has no paths")
	}
	var ops []*operation
	seen := map[string]string{}
	for _, p := range paths.keys {
		item := paths.m[p]
		if ref := item.s("$ref"); ref != "" {
			t, err := resolvePointer(g.root, ref)
			if err != nil {
				return nil, err
			}
			item = t
		}
		for _, m := range httpMethods {
			op := item.get(m)
			if !op.isObject() {
				continue
			}
			group := op.s("x-sdk-group")
			method := op.s("x-sdk-method")
			if group == "" {
				if tags := op.get("tags"); tags != nil && len(tags.items) > 0 {
					group = strings.ToLower(naming.Camel(tags.items[0].text()))
				} else {
					group = "api"
				}
				g.warn("%s %s has no x-sdk-group; using %q", strings.ToUpper(m), p, group)
			}
			if method == "" {
				method = op.s("operationId")
				if method == "" {
					method = m + naming.Pascal(p)
				}
				method = naming.Camel(method)
				g.warn("%s %s has no x-sdk-method; using %q", strings.ToUpper(m), p, method)
			}
			key := naming.Field(group) + "." + naming.Method(method)
			if prev, dup := seen[key]; dup {
				return nil, fmt.Errorf("%s %s and %s both map to %s", strings.ToUpper(m), p, prev, key)
			}
			seen[key] = strings.ToUpper(m) + " " + p
			params, err := g.mergeParams(item.get("parameters"), op.get("parameters"))
			if err != nil {
				return nil, fmt.Errorf("%s %s: %w", strings.ToUpper(m), p, err)
			}
			ops = append(ops, &operation{
				group: group, method: method, httpMethod: strings.ToUpper(m), path: p,
				op: op, params: params, summary: oneLine(op.s("summary")),
			})
		}
	}
	return ops, nil
}

func (g *gen) mergeParams(lists ...*node) ([]param, error) {
	var out []param
	idx := map[string]int{}
	for _, l := range lists {
		if l == nil || l.kind != 'a' {
			continue
		}
		for _, pn := range l.items {
			if ref := pn.s("$ref"); ref != "" {
				t, err := resolvePointer(g.root, ref)
				if err != nil {
					return nil, err
				}
				pn = t
			}
			p := param{
				name: pn.s("name"), in: pn.s("in"), required: pn.truthy("required"),
				schema: pn.get("schema"), desc: pn.s("description"), explode: pn.truthy("explode"),
				deprec: pn.truthy("deprecated"),
			}
			if p.schema == nil {
				if c := pn.get("content"); c.isObject() && len(c.keys) > 0 {
					p.schema = c.m[c.keys[0]].get("schema")
				}
			}
			if p.name == "" || p.in == "" {
				return nil, fmt.Errorf("parameter without name/in")
			}
			k := p.in + ":" + p.name
			if i, ok := idx[k]; ok {
				out[i] = p
				continue
			}
			idx[k] = len(out)
			out = append(out, p)
		}
	}
	return out, nil
}

type pathArg struct{ name, wire, goType string }

var pathVar = regexp.MustCompile(`\{([^}]+)\}`)

// scalarGo is the Go type of a scalar parameter schema, or "" if the schema
// is not a scalar (or a list of scalars).
func (g *gen) scalarKind(s *node) string {
	return g.jsonKind(s, 0)
}

// genOperation emits one operation's method, iterator, request builder,
// params and response types into g.cur, and returns its README rows.
func (g *gen) genOperation(op *operation, paramsName, respName string) string {
	method := naming.Method(op.method)
	svc := naming.Service(op.group)
	methodSlot := g.reserve()

	// Path params, in path order.
	var pathArgs []pathArg
	argUsed := map[string]bool{}
	byName := map[string]param{}
	for _, p := range op.params {
		if p.in == "path" {
			byName[p.name] = p
		}
	}
	for _, m := range pathVar.FindAllStringSubmatch(op.path, -1) {
		wire := m[1]
		arg := naming.Camel(wire)
		for i := 2; argUsed[arg]; i++ {
			arg = naming.Camel(wire) + fmt.Sprint(i)
		}
		argUsed[arg] = true
		gt := "string"
		if p, ok := byName[wire]; ok {
			switch g.typeOfScalarName(p.schema) {
			case "int64":
				gt = "int64"
			case "float64":
				gt = "float64"
			case "bool":
				gt = "bool"
			}
		}
		pathArgs = append(pathArgs, pathArg{name: arg, wire: wire, goType: gt})
	}

	// Params struct.
	paramsSlot := g.reserve()
	var fields []paramsField
	fieldUsed := map[string]bool{}
	uniqueField := func(base, suffix string) string {
		n := base
		if fieldUsed[n] {
			n = base + suffix
		}
		for i := 2; fieldUsed[n]; i++ {
			n = fmt.Sprintf("%s%s%d", base, suffix, i)
		}
		fieldUsed[n] = true
		return n
	}
	for _, p := range op.params {
		if p.in != "query" && p.in != "header" {
			if p.in == "cookie" {
				g.warn("%s %s: cookie parameter %q is not supported; skipped", op.httpMethod, op.path, p.name)
			}
			continue
		}
		var fname string
		if p.in == "header" {
			fname = uniqueField(naming.HeaderField(p.name), "Header")
		} else {
			fname = uniqueField(naming.Pascal(p.name), "Query")
		}
		f := paramsField{name: fname, wire: p.name, in: p.in, explode: p.explode}
		doc := p.desc
		if p.required {
			doc = strings.TrimSpace(doc + "\n\nRequired.")
		}
		if p.deprec {
			doc = strings.TrimSpace(doc + "\n\nDeprecated: the API marks this parameter deprecated.")
		}
		f.doc = doc
		switch g.scalarKind(p.schema) {
		case "string", "number", "boolean":
			t := g.typeOf(p.schema, paramsName+fname)
			f.t = t
			f.expr = "*" + t.expr
			f.pointer = true
		case "array":
			items := resolveItems(g, p.schema)
			switch g.scalarKind(items) {
			case "string", "number", "boolean":
				t := g.typeOf(p.schema, paramsName+fname)
				f.t = t
				f.expr = t.expr
				f.list = true
			default:
				f.t = goType{expr: "string", kind: kScalar}
				f.expr = "*string"
				f.pointer = true
				f.doc = strings.TrimSpace(f.doc + "\n\nSent verbatim (the spec's item type is not a scalar).")
			}
		default:
			f.t = goType{expr: "string", kind: kScalar}
			f.expr = "*string"
			f.pointer = true
		}
		f.tag = fmt.Sprintf("`%s:%q json:\"-\"`", p.in, p.name)
		fields = append(fields, f)
	}

	// Request body.
	const (
		bodyNone = iota
		bodyObject
		bodyWhole
	)
	bodyMode := bodyNone
	bodyRequired := false
	if rb := g.deref(op.op.get("requestBody")); rb != nil {
		bodyRequired = rb.truthy("required")
		var js *node
		if c := rb.get("content"); c.isObject() {
			for _, ct := range c.keys {
				if isJSONMedia(ct) {
					js = c.m[ct]
					break
				}
			}
			if js == nil && len(c.keys) > 0 {
				g.warn("%s %s: request body has no JSON media type; sending params.Body as JSON", op.httpMethod, op.path)
				js = c.m[c.keys[0]]
			}
		}
		if js != nil {
			bs := js.get("schema")
			if props, req, ok := g.objectProps(bs); ok && len(props) > 0 {
				bodyMode = bodyObject
				for _, pr := range props {
					fname := uniqueField(naming.Pascal(pr.name), "Body")
					t := g.typeOf(pr.schema, paramsName+fname)
					expr, opts := fieldDecl(t, req[pr.name], paramsName)
					doc := schemaDoc(pr.schema)
					if req[pr.name] {
						doc = strings.TrimSpace(doc + "\n\nRequired (JSON body).")
					}
					fields = append(fields, paramsField{
						name: fname, wire: pr.name, in: "body", expr: expr, t: t, doc: doc,
						pointer: strings.HasPrefix(expr, "*"),
						tag:     fmt.Sprintf("`json:%q`", pr.name+opts),
					})
				}
			} else {
				bodyMode = bodyWhole
				t := g.typeOf(bs, paramsName+"Body")
				fname := uniqueField("Body", "_")
				expr := t.expr
				if t.kind == kScalar || t.kind == kStruct {
					expr = "*" + expr
				}
				fields = append(fields, paramsField{
					name: fname, wire: "", in: "bodywhole", expr: expr, t: t,
					doc: "Body is sent as the JSON request body.", pointer: strings.HasPrefix(expr, "*"), tag: "`json:\"-\"`",
				})
			}
		}
	}

	{
		var b strings.Builder
		writeDoc(&b, "", paramsName, "", fmt.Sprintf("%s holds the query, header and JSON-body parameters of [%s.%s]. Pass nil when you need none.", paramsName, svc, method))
		fmt.Fprintf(&b, "type %s struct {\n", paramsName)
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
		g.fill(paramsSlot, b.String())
	}

	// Response.
	const (
		respNone = iota
		respJSON
		respText
		respMixed
	)
	respMode := respNone
	var accept []string
	var respT goType
	if resp := g.successResponse(op.op); resp != nil {
		c := resp.get("content")
		var jsonCT string
		var textCTs []string
		if c.isObject() {
			for _, ct := range c.keys {
				if isJSONMedia(ct) {
					if jsonCT == "" {
						jsonCT = ct
					}
				} else {
					textCTs = append(textCTs, ct)
				}
			}
		}
		switch {
		case jsonCT != "" && len(textCTs) == 0:
			respMode = respJSON
			respT = g.declareNamed(respName, c.m[jsonCT].get("schema"), op.summary)
		case jsonCT == "" && len(textCTs) > 0:
			respMode = respText
			accept = append(textCTs, "application/json")
		case jsonCT != "" && len(textCTs) > 0:
			respMode = respMixed
			accept = append(append([]string{}, textCTs...), jsonCT)
			slot := g.reserve()
			jt := g.typeOf(c.m[jsonCT].get("schema"), respName+"JSON")
			jexpr := jt.expr
			if jt.kind == kScalar || jt.kind == kStruct {
				jexpr = "*" + jexpr
			}
			var b strings.Builder
			writeDoc(&b, "", respName, op.summary, fmt.Sprintf("The server answers %s or JSON; Text holds a non-JSON body and JSON a decoded JSON one.", strings.Join(textCTs, ", ")))
			fmt.Fprintf(&b, "type %s struct {\n\t// ContentType is the response Content-Type.\n\tContentType string\n\t// Text is the body when it is not JSON (%s).\n\tText string\n\t// JSON is the decoded body when the server answered with JSON.\n\tJSON %s\n}\n",
				respName, strings.Join(textCTs, ", "), jexpr)
			g.fill(slot, b.String())
			respT = goType{expr: respName, kind: kStruct}
		}
	}

	// Method, iterator and request builder.
	var b strings.Builder
	sigArgs := []string{"ctx context.Context"}
	callArgs := []string{}
	builderParams := []string{}
	for _, a := range pathArgs {
		sigArgs = append(sigArgs, a.name+" "+a.goType)
		callArgs = append(callArgs, a.name)
		builderParams = append(builderParams, a.name+" "+a.goType)
	}
	sigArgs = append(sigArgs, "params *"+paramsName, "opts ...RequestOption")
	builderParams = append(builderParams, "params *"+paramsName)
	builder := "build" + naming.Pascal(op.group) + method + "Request"
	callBuild := builder + "(" + strings.Join(append(append([]string{}, callArgs...), "params"), ", ") + ")"

	httpLine := op.httpMethod + " " + op.path
	doc := op.summary
	if d := strings.TrimSpace(op.op.s("description")); d != "" && d != op.summary {
		doc = strings.TrimSpace(doc + "\n\n" + d)
	}
	extra := "HTTP: " + httpLine
	if op.op.truthy("deprecated") {
		extra += "\n\nDeprecated: the API marks this operation deprecated."
	}
	writeMethodDoc(&b, method, doc, extra)
	switch respMode {
	case respJSON:
		fmt.Fprintf(&b, "func (s *%s) %s(%s) (*%s, error) {\n\tvar out %s\n\tif err := s.client.do(ctx, %s, opts, decodeJSON(&out)); err != nil {\n\t\treturn nil, err\n\t}\n\treturn &out, nil\n}\n",
			svc, method, strings.Join(sigArgs, ", "), respT.expr, respT.expr, callBuild)
	case respText:
		fmt.Fprintf(&b, "func (s *%s) %s(%s) (string, error) {\n\tvar out string\n\tif err := s.client.do(ctx, %s, opts, decodeText(&out)); err != nil {\n\t\treturn \"\", err\n\t}\n\treturn out, nil\n}\n",
			svc, method, strings.Join(sigArgs, ", "), callBuild)
	case respMixed:
		fmt.Fprintf(&b, "func (s *%s) %s(%s) (*%s, error) {\n\tvar out %s\n\tif err := s.client.do(ctx, %s, opts, decodeMixed(&out.ContentType, &out.Text, func(b []byte) error {\n\t\treturn json.Unmarshal(b, &out.JSON)\n\t})); err != nil {\n\t\treturn nil, err\n\t}\n\treturn &out, nil\n}\n",
			svc, method, strings.Join(sigArgs, ", "), respName, respName, callBuild)
	default:
		fmt.Fprintf(&b, "func (s *%s) %s(%s) error {\n\treturn s.client.do(ctx, %s, opts, nil)\n}\n",
			svc, method, strings.Join(sigArgs, ", "), callBuild)
	}

	rows := fmt.Sprintf("| `%s(%s)` | `%s` | %s |\n", method, readmeArgs(pathArgs, "params"), httpLine, mdCell(op.summary))

	// Pagination.
	if pg := op.op.get("x-sdk-pagination"); pg.isObject() {
		code, row := g.iterator(op, pg, svc, method, paramsName, respT, respMode == respJSON, pathArgs, builder, fields)
		if code != "" {
			b.WriteString("\n")
			b.WriteString(code)
			rows += row
		}
	}

	// Builder.
	path := pathExpr(op.path, pathArgs)
	fmt.Fprintf(&b, "\nfunc %s(%s) *apiRequest {\n\treq := newRequest(%q, %s)\n", builder, strings.Join(builderParams, ", "), op.httpMethod, path)
	if len(accept) > 0 {
		fmt.Fprintf(&b, "\treq.accept = %q\n", strings.Join(dedupe(accept), ", "))
	}
	var enc []string
	for _, f := range fields {
		switch {
		case f.in == "query" && f.list:
			enc = append(enc, fmt.Sprintf("addQueryList(req.query, %q, params.%s, %t)", f.wire, f.name, f.explode))
		case f.in == "query":
			enc = append(enc, fmt.Sprintf("addQuery(req.query, %q, params.%s)", f.wire, f.name))
		case f.in == "header":
			enc = append(enc, fmt.Sprintf("setHeader(req.header, %q, params.%s)", f.wire, f.name))
		}
	}
	switch bodyMode {
	case bodyObject:
		enc = append(enc, "req.body = params")
	case bodyWhole:
		for _, f := range fields {
			if f.in == "bodywhole" {
				enc = append(enc, fmt.Sprintf("req.body = params.%s", f.name))
			}
		}
	}
	if len(enc) > 0 {
		b.WriteString("\tif params != nil {\n")
		for _, e := range enc {
			b.WriteString("\t\t" + e + "\n")
		}
		b.WriteString("\t}")
		if bodyMode != bodyNone && bodyRequired {
			b.WriteString(" else {\n\t\treq.body = struct{}{}\n\t}")
		}
		b.WriteString("\n")
	}
	b.WriteString("\treturn req\n}\n")
	g.fill(methodSlot, b.String())
	return rows
}

func (g *gen) typeOfScalarName(s *node) string {
	switch g.jsonKind(s, 0) {
	case "number":
		r := g.deref(s)
		if types, _ := schemaTypes(r); len(types) == 1 && types[0] == "integer" {
			return "int64"
		}
		return "float64"
	case "boolean":
		return "bool"
	}
	return "string"
}

func resolveItems(g *gen, s *node) *node {
	return g.deref(s).get("items")
}

// deref follows $ref chains.
func (g *gen) deref(n *node) *node {
	for i := 0; i < 20 && n.isObject(); i++ {
		ref := n.s("$ref")
		if ref == "" {
			return n
		}
		t, err := resolvePointer(g.root, ref)
		if err != nil {
			g.fail(err)
			return nil
		}
		n = t
	}
	return n
}

// objectProps returns the properties of an object body schema.
func (g *gen) objectProps(s *node) ([]prop, map[string]bool, bool) {
	s = g.deref(s)
	if !s.isObject() || s.has("oneOf") || s.has("anyOf") {
		return nil, nil, false
	}
	if s.has("allOf") {
		return g.mergeAllOf(s)
	}
	types, _ := schemaTypes(s)
	if len(types) > 0 && types[0] != "object" {
		return nil, nil, false
	}
	p := s.get("properties")
	if !p.isObject() {
		return nil, nil, false
	}
	var props []prop
	for _, k := range p.keys {
		props = append(props, prop{name: k, schema: p.m[k]})
	}
	return props, requiredSet(s), true
}

func (g *gen) successResponse(op *node) *node {
	rs := op.get("responses")
	if !rs.isObject() {
		return nil
	}
	codes := append([]string(nil), rs.keys...)
	sort.Strings(codes)
	for _, c := range codes {
		if strings.HasPrefix(c, "2") {
			r := g.deref(rs.m[c])
			if r.get("content").isObject() && len(r.get("content").keys) > 0 {
				return r
			}
		}
	}
	return nil
}

func isJSONMedia(ct string) bool {
	ct = strings.ToLower(strings.TrimSpace(strings.SplitN(ct, ";", 2)[0]))
	return ct == "application/json" || strings.HasSuffix(ct, "+json")
}

func pathExpr(path string, args []pathArg) string {
	var parts []string
	rest := path
	i := 0
	for {
		loc := pathVar.FindStringIndex(rest)
		if loc == nil {
			break
		}
		if loc[0] > 0 {
			parts = append(parts, fmt.Sprintf("%q", rest[:loc[0]]))
		}
		parts = append(parts, "pathParam("+args[i].name+")")
		i++
		rest = rest[loc[1]:]
	}
	if rest != "" || len(parts) == 0 {
		parts = append(parts, fmt.Sprintf("%q", rest))
	}
	return strings.Join(parts, "+")
}

func readmeArgs(args []pathArg, last ...string) string {
	out := []string{"ctx"}
	for _, a := range args {
		out = append(out, a.name)
	}
	return strings.Join(append(out, last...), ", ")
}

func dedupe(ss []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, s := range ss {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}

func mdCell(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "|", `\|`), "\n", " ")
}

// iterator emits the XxxIter method for a paginated operation.
func (g *gen) iterator(op *operation, pg *node, svc, method, paramsName string, respT goType, isJSON bool,
	pathArgs []pathArg, builder string, fields []paramsField) (string, string) {
	style := pg.s("style")
	itemsKey := pg.s("items")
	find := func(wire string) *paramsField {
		for i := range fields {
			if fields[i].wire == wire && (fields[i].in == "query" || fields[i].in == "body") {
				return &fields[i]
			}
		}
		return nil
	}
	isInt := func(f *paramsField) bool { return f != nil && strings.TrimPrefix(f.expr, "*") == "int64" }
	itemType := "json.RawMessage"
	if isJSON {
		name := respT.expr
		if a, ok := g.aliasOf[name]; ok {
			name = a
		}
		if fs, ok := g.structFields[name]; ok {
			if t, ok := fs[itemsKey]; ok && t.kind == kSlice {
				itemType = strings.TrimPrefix(t.expr, "[]")
			}
		}
	}
	iterName := naming.IterMethod(op.method)
	sig := []string{"ctx context.Context"}
	args := []string{}
	for _, a := range pathArgs {
		sig = append(sig, a.name+" "+a.goType)
		args = append(args, a.name)
	}
	sig = append(sig, "params *"+paramsName, "iter IterOptions", "opts ...RequestOption")
	buildCall := builder + "(" + strings.Join(append(args, "&q"), ", ") + ")"
	ref := func(f *paramsField) string {
		if f == nil {
			return "nil"
		}
		if f.pointer {
			return "p." + f.name
		}
		return "&p." + f.name
	}
	assign := func(f *paramsField, v string) string {
		if f.pointer {
			return fmt.Sprintf("q.%s = &%s", f.name, v)
		}
		return fmt.Sprintf("q.%s = %s", f.name, v)
	}
	var b strings.Builder
	switch style {
	case "offset":
		if itemsKey == "" {
			itemsKey = "data"
		}
		lim, off := find("limit"), find("offset")
		if !isInt(lim) || !isInt(off) {
			g.warn("%s %s: offset pagination needs integer limit and offset parameters; no %s generated", op.httpMethod, op.path, iterName)
			return "", ""
		}
		writeComment(&b, "", fmt.Sprintf("%s iterates every item of [%s.%s] across pages (offset pagination over %q): it advances offset by the page size and stops on a short page, when offset reaches total, or when has_more is false. Use [IterOptions] for the page size and an item cap.", iterName, svc, method, itemsKey))
		fmt.Fprintf(&b, "func (s *%s) %s(%s) *Iter[%s] {\n\tvar p %s\n\tif params != nil {\n\t\tp = *params\n\t}\n", svc, iterName, strings.Join(sig, ", "), itemType, paramsName)
		fmt.Fprintf(&b, "\treturn newOffsetIter[%s](ctx, s.client, iter, %q, %s, %s, func(limit, offset int64) *apiRequest {\n\t\tq := p\n\t\t%s\n\t\t%s\n\t\treturn %s\n\t}, opts)\n}\n",
			itemType, itemsKey, ref(lim), ref(off), assign(lim, "limit"), assign(off, "offset"), buildCall)
	case "cursor":
		if itemsKey == "" {
			itemsKey = "results"
		}
		cp := pg.s("cursor_param")
		if cp == "" {
			cp = "after"
		}
		next := pg.s("next")
		if next == "" {
			next = "nextCursor"
		}
		cur := find(cp)
		if cur == nil || cur.expr != "*string" {
			g.warn("%s %s: cursor pagination needs a string %q parameter; no %s generated", op.httpMethod, op.path, cp, iterName)
			return "", ""
		}
		lim := find("limit")
		if lim != nil && !isInt(lim) {
			lim = nil
		}
		writeComment(&b, "", fmt.Sprintf("%s iterates every item of [%s.%s] across pages (cursor pagination over %q): it passes the previous page's %q as %q until the cursor is null/absent or hasMore is false. Use [IterOptions] for the page size and an item cap.", iterName, svc, method, itemsKey, next, cp))
		fmt.Fprintf(&b, "func (s *%s) %s(%s) *Iter[%s] {\n\tvar p %s\n\tif params != nil {\n\t\tp = *params\n\t}\n", svc, iterName, strings.Join(sig, ", "), itemType, paramsName)
		limAssign := ""
		if lim != nil {
			if lim.pointer {
				limAssign = fmt.Sprintf("\t\tif limit != nil {\n\t\t\tq.%s = limit\n\t\t}\n", lim.name)
			} else {
				limAssign = fmt.Sprintf("\t\tif limit != nil {\n\t\t\tq.%s = *limit\n\t\t}\n", lim.name)
			}
		}
		fmt.Fprintf(&b, "\treturn newCursorIter[%s](ctx, s.client, iter, %q, %q, %s, func(limit *int64, cursor *string) *apiRequest {\n\t\tq := p\n%s\t\tif cursor != nil {\n\t\t\tq.%s = cursor\n\t\t}\n\t\treturn %s\n\t}, opts)\n}\n",
			itemType, itemsKey, next, ref(lim), limAssign, cur.name, buildCall)
		if lim == nil {
			// The closure ignores limit when the operation has none.
			s := b.String()
			s = strings.Replace(s, "func(limit *int64, cursor *string)", "func(_ *int64, cursor *string)", 1)
			b.Reset()
			b.WriteString(s)
		}
	default:
		g.warn("%s %s: unknown x-sdk-pagination style %q; no iterator generated", op.httpMethod, op.path, style)
		return "", ""
	}
	row := fmt.Sprintf("| `%s(%s)` | %s pages of `%s` | iterator over `%s` |\n", iterName, readmeArgs(pathArgs, "params", "iterOpts"), style, itemsKey, method)
	return b.String(), row
}

// ---- doc comments ----

func oneLine(s string) string { return strings.Join(strings.Fields(s), " ") }

func schemaDoc(s *node) string {
	if !s.isObject() {
		return ""
	}
	d := strings.TrimSpace(s.s("description"))
	if d == "" {
		d = strings.TrimSpace(s.s("title"))
	}
	if s.truthy("deprecated") {
		d = strings.TrimSpace(d + "\n\nDeprecated: the API marks this deprecated.")
	}
	return d
}

// writeDoc writes a type comment that starts with the type name.
func writeDoc(b *strings.Builder, indent, name, desc, extra string) {
	desc = strings.TrimSpace(desc)
	var text string
	switch {
	case extra != "" && strings.HasPrefix(extra, name+" "):
		text = extra
		if desc != "" {
			text += "\n\n" + desc
		}
	case desc != "":
		text = name + ": " + desc
		if extra != "" {
			text += "\n\n" + extra
		}
	default:
		text = name + " is generated from the OpenAPI spec."
		if extra != "" {
			text += " " + extra
		}
	}
	writeComment(b, indent, text)
}

func writeMethodDoc(b *strings.Builder, name, desc, extra string) {
	text := name + " calls the API."
	if desc != "" {
		text = name + ": " + desc
	}
	if extra != "" {
		text += "\n\n" + extra
	}
	writeComment(b, "", text)
}

// writeComment writes text as // lines wrapped near 100 columns.
func writeComment(b *strings.Builder, indent, text string) {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	paras := strings.Split(strings.TrimSpace(text), "\n")
	for i, line := range paras {
		line = strings.TrimRight(line, " \t")
		if strings.TrimSpace(line) == "" {
			if i > 0 && i < len(paras)-1 {
				b.WriteString(indent + "//\n")
			}
			continue
		}
		words := strings.Fields(line)
		cur := ""
		for _, w := range words {
			if cur != "" && len(cur)+1+len(w) > 96 {
				b.WriteString(indent + "// " + cur + "\n")
				cur = w
				continue
			}
			if cur == "" {
				cur = w
			} else {
				cur += " " + w
			}
		}
		if cur != "" {
			b.WriteString(indent + "// " + cur + "\n")
		}
	}
}
