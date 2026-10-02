package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

const repoRoot = "../../.."

// TestCommittedFilesUpToDate regenerates from the vendored openapi.json and
// compares with the committed gen_*.go files (and checks determinism).
func TestCommittedFilesUpToDate(t *testing.T) {
	spec, err := os.ReadFile(filepath.Join(repoRoot, "openapi.json"))
	if err != nil {
		t.Fatal(err)
	}
	a, err := Generate(spec, "propraven")
	if err != nil {
		t.Fatal(err)
	}
	b, err := Generate(spec, "propraven")
	if err != nil {
		t.Fatal(err)
	}
	for name, src := range a.Files {
		if !bytes.Equal(src, b.Files[name]) {
			t.Errorf("%s: output is not deterministic", name)
		}
		disk, err := os.ReadFile(filepath.Join(repoRoot, name))
		if err != nil || !bytes.Equal(disk, src) {
			t.Errorf("%s is stale: run `go run ./internal/cmd/generate`", name)
		}
	}
	existing, _ := filepath.Glob(filepath.Join(repoRoot, "gen_*.go"))
	for _, p := range existing {
		if _, ok := a.Files[filepath.Base(p)]; !ok {
			t.Errorf("%s is not produced by the generator any more", p)
		}
	}
	readme, err := os.ReadFile(filepath.Join(repoRoot, "README.md"))
	if err == nil && !strings.Contains(string(readme), a.MethodTable) {
		t.Errorf("README method table is stale: run `go run ./internal/cmd/generate`")
	}
}

// stressSpec exercises schema shapes the real spec may grow into.
const stressSpec = `{
  "openapi": "3.1.0",
  "info": {"title": "stress", "version": "0"},
  "paths": {
    "/api/v1/things/{thing_id}/parts/{n}": {
      "parameters": [
        {"name": "thing_id", "in": "path", "required": true, "schema": {"type": "string"}},
        {"$ref": "#/components/parameters/N"}
      ],
      "get": {
        "x-sdk-group": "things", "x-sdk-method": "parts",
        "parameters": [
          {"name": "tags", "in": "query", "schema": {"type": "array", "items": {"type": "string"}}},
          {"name": "ids", "in": "query", "explode": true, "schema": {"type": "array", "items": {"type": "integer"}}},
          {"name": "filter", "in": "query", "schema": {"type": "object", "properties": {"a": {"type": "string"}}}},
          {"name": "mode", "in": "query", "schema": {"type": ["string", "null"], "enum": ["a-b", "a_b", "30yr+", "", "type"]}},
          {"name": "X-Request-Thing", "in": "header", "schema": {"type": "string"}},
          {"name": "session", "in": "cookie", "schema": {"type": "string"}},
          {"name": "limit", "in": "query", "schema": {"type": "integer"}},
          {"name": "offset", "in": "query", "schema": {"type": "integer"}}
        ],
        "x-sdk-pagination": {"style": "offset", "items": "rows"},
        "responses": {"200": {"$ref": "#/components/responses/Parts"}}
      }
    },
    "/api/v1/things": {
      "post": {
        "x-sdk-group": "things", "x-sdk-method": "createMany",
        "requestBody": {"$ref": "#/components/requestBodies/Many"},
        "responses": {"201": {"description": "ok", "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/Thing"}}}}}}
      },
      "get": {
        "operationId": "listThings",
        "parameters": [
          {"name": "cursor", "in": "query", "schema": {"type": "string"}},
          {"name": "limit", "in": "query", "schema": {"type": "number"}}
        ],
        "x-sdk-pagination": {"style": "cursor", "items": "items", "cursor_param": "cursor", "next": "next"},
        "responses": {"200": {"description": "ok", "content": {"application/json": {"schema": {
          "type": "object", "properties": {"items": {"type": "array", "items": {"$ref": "#/components/schemas/Thing"}}, "next": {"type": ["string", "null"]}}}}}}}
      }
    },
    "/api/v1/things/search": {
      "post": {
        "x-sdk-group": "things", "x-sdk-method": "search",
        "x-sdk-pagination": {"style": "offset", "items": "data"},
        "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [
          {"type": "object", "properties": {"limit": {"type": "integer"}, "offset": {"type": "integer"}}},
          {"type": "object", "required": ["q"], "properties": {"q": {"type": "string"}, "range": {"type": "object", "properties": {"min": {"type": ["number", "null"]}}}}}
        ]}}}},
        "responses": {"200": {"description": "ok", "content": {"application/json": {"schema": {"anyOf": [
          {"type": "object", "properties": {"data": {"type": "array", "items": {"type": "object", "properties": {"x": {"type": "number"}}}}}},
          {"type": "object", "properties": {"data": {"type": "array", "items": {"type": "string"}}}}
        ]}}}}}
      }
    },
    "/api/v1/things/{id}": {
      "delete": {
        "x-sdk-group": "things", "x-sdk-method": "delete",
        "parameters": [{"name": "id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}],
        "responses": {"204": {"description": "gone"}}
      },
      "put": {
        "x-sdk-group": "things", "x-sdk-method": "type",
        "parameters": [{"name": "id", "in": "path", "required": true, "schema": {"type": "string"}}],
        "requestBody": {"content": {"application/json": {"schema": {"type": "array", "items": {"type": "string"}}}}},
        "responses": {"200": {"description": "ok", "content": {"text/csv": {"schema": {"type": "string"}}, "application/json": {"schema": {"type": "array", "items": {"type": "integer"}}}}}}
      }
    },
    "/api/v1/raw": {
      "get": {
        "x-sdk-group": "client", "x-sdk-method": "get",
        "responses": {"200": {"description": "ok", "content": {"text/plain": {"schema": {"type": "string"}}}}}
      }
    }
  },
  "components": {
    "parameters": {"N": {"name": "n", "in": "path", "required": true, "schema": {"type": "integer"}}},
    "requestBodies": {"Many": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {
      "things": {"type": "array", "items": {"$ref": "#/components/schemas/Thing"}},
      "meta": {"type": "object", "additionalProperties": {"type": "number"}},
      "free": {"type": "object", "additionalProperties": true},
      "anything": {},
      "flag": true
    }}}}}},
    "responses": {"Parts": {"description": "ok", "content": {"application/json": {"schema": {"type": "object", "properties": {
      "rows": {"type": "array", "items": {"$ref": "#/components/schemas/Part"}},
      "total": {"type": "integer"}
    }}}}}},
    "schemas": {
      "Error": {"type": "object", "properties": {"code": {"type": "string"}}},
      "Client": {"type": "string", "enum": ["x", "y"]},
      "Thing": {"type": "object", "required": ["id", "child"], "properties": {
        "id": {"type": "string"},
        "child": {"$ref": "#/components/schemas/Thing"},
        "children": {"type": "array", "items": {"$ref": "#/components/schemas/Thing"}},
        "grid": {"type": "array", "items": {"type": "array", "items": {"type": ["number", "null"]}}},
        "shape": {"oneOf": [{"type": "array", "items": {"type": "string"}}, {"type": "object", "properties": {"a": {"type": "integer"}}}, {"type": "string"}, {"type": "number"}, {"type": "boolean"}, {"type": "null"}]},
        "ambiguous": {"oneOf": [{"type": "object", "properties": {"a": {"type": "string"}}}, {"type": "object", "properties": {"b": {"type": "string"}}}]},
        "maybe": {"anyOf": [{"$ref": "#/components/schemas/Part"}, {"type": "null"}]},
        "wrapped": {"allOf": [{"$ref": "#/components/schemas/Part"}], "description": "single allOf"},
        "merged": {"allOf": [{"$ref": "#/components/schemas/Part"}, {"type": "object", "properties": {"extra": {"type": "boolean"}}}]},
        "mixed": {"type": ["string", "number"]},
        "counts": {"type": "object", "additionalProperties": {"type": "integer"}},
        "objs": {"type": "object", "additionalProperties": {"$ref": "#/components/schemas/Part"}},
        "kind": {"const": "thing"},
        "nothing": {"type": "null"},
        "has_more": {"type": "boolean"},
        "hasMore": {"type": "boolean"},
        "legacy": {"type": "string", "nullable": true, "deprecated": true}
      }},
      "Part": {"type": "object", "properties": {"weight": {"type": "number"}, "count": {"type": "integer"}, "sizes": {"type": "array", "items": {"type": "integer"}}}},
      "Parts": {"type": "array", "items": {"type": "object", "properties": {"z": {"type": "string"}}}},
      "Empty": {},
      "Free": {"type": "object"},
      "Status": {"type": "string", "enum": ["on", "off"]},
      "Alias": {"$ref": "#/components/schemas/Part"}
    }
  }
}`

// TestStressSpecCompiles generates from stressSpec and builds the result
// together with the real hand-written core in a scratch module.
func TestStressSpecCompiles(t *testing.T) {
	if testing.Short() {
		t.Skip("builds a scratch module")
	}
	res, err := Generate([]byte(stressSpec), "propraven")
	if err != nil {
		t.Fatal(err)
	}
	if res.Operations != 7 {
		t.Fatalf("operations = %d", res.Operations)
	}
	dir := t.TempDir()
	must := func(err error) {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
	}
	must(os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module github.com/jdw2111/propraven-go\n\ngo 1.22\n"), 0o644))
	must(os.MkdirAll(filepath.Join(dir, "internal"), 0o755))
	v, err := os.ReadFile(filepath.Join(repoRoot, "internal", "version.go"))
	must(err)
	must(os.WriteFile(filepath.Join(dir, "internal", "version.go"), v, 0o644))
	core, _ := filepath.Glob(filepath.Join(repoRoot, "*.go"))
	for _, p := range core {
		base := filepath.Base(p)
		if strings.HasPrefix(base, "gen_") || strings.HasSuffix(base, "_test.go") {
			continue
		}
		src, err := os.ReadFile(p)
		must(err)
		must(os.WriteFile(filepath.Join(dir, base), src, 0o644))
	}
	for name, src := range res.Files {
		must(os.WriteFile(filepath.Join(dir, name), src, 0o644))
	}
	// A usage file pins down the shapes the generator should produce.
	usage := `package propraven

import (
	"context"
	"encoding/json"
)

var _ = func(c *Client) {
	ctx := context.Background()
	r, _ := c.Things.Parts(ctx, "t1", 3, &ThingsPartsParams{Tags: []string{"a"}, IDs: []int64{1, 2}, Mode: String(ThingsPartsParamsModeAB), RequestThing: String("h")})
	var _ []Part = r.Rows
	var _ *int64 = r.Total
	_ = c.Things.PartsIter(ctx, "t1", 3, nil, IterOptions{})
	created, _ := c.Things.CreateMany(ctx, &ThingsCreateManyParams{Things: []Thing{{ID: "x"}}, Meta: map[string]float64{"a": 1}})
	var _ []Thing = *created
	var _ *Thing = Thing{}.Child
	var _ []Thing = Thing{}.Children
	var _ [][]*float64 = Thing{}.Grid
	var sh ThingShape
	var _ []string = sh.Array
	var _ *ThingShapeObject = sh.Object
	var _ *string = sh.String
	var _ *float64 = sh.Number
	var _ *bool = sh.Bool
	var _ json.RawMessage = Thing{}.Ambiguous
	var _ *Part = Thing{}.Maybe
	var _ *Part = Thing{}.Wrapped
	var _ *bool = Thing{}.Merged.Extra
	var _ json.RawMessage = Thing{}.Mixed
	var _ map[string]int64 = Thing{}.Counts
	var _ map[string]Part = Thing{}.Objs
	var _ *string = Thing{}.Kind
	var _ *bool = Thing{}.HasMore
	var _ *bool = Thing{}.HasMore2
	var _ *string = Thing{}.Legacy
	var _ ErrorSchema = ErrorSchema{}
	var _ string = ClientSchemaX
	var _ Status = StatusOn
	var _ Part = Alias{}
	var _ []PartsItem = Parts{}
	var _ json.RawMessage = Empty{}
	var _ map[string]any = Free{}
	s, _ := c.Things.Search(ctx, &ThingsSearchParams{Q: "q", Range: &ThingsSearchParamsRange{Min: Float(1)}})
	var _ json.RawMessage = *s
	_ = c.Things.SearchIter(ctx, nil, IterOptions{PageSize: 5})
	var _ error = c.Things.Delete(ctx, "id", nil)
	m, _ := c.Things.Type(ctx, "id", &ThingsTypeParams{Body: []string{"a"}})
	var _ string = m.Text
	var _ []int64 = m.JSON
	var _ string = must(c.Client.Get(ctx, nil))
	l, _ := c.API.ListThings(ctx, &APIListThingsParams{Cursor: String("c")})
	var _ *string = l.Next
	_ = c.API.ListThingsIter(ctx, nil, IterOptions{})
}

func must(s string, _ error) string { return s }
`
	must(os.WriteFile(filepath.Join(dir, "usage_stress.go"), []byte(usage), 0o644))
	goBin := filepath.Join(runtime.GOROOT(), "bin", "go")
	cmd := exec.Command(goBin, "vet", "./...")
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOFLAGS=-mod=mod", "GOWORK=off")
	if out, err := cmd.CombinedOutput(); err != nil {
		var listing strings.Builder
		for name, src := range res.Files {
			listing.WriteString("==== " + name + "\n" + numbered(string(src)))
		}
		t.Fatalf("go vet failed: %v\n%s\n%s", err, out, listing.String())
	}
	if len(res.Warnings) == 0 {
		t.Error("expected warnings (cookie param, missing x-sdk fields, number limit)")
	}
	for _, w := range res.Warnings {
		t.Log("warning:", w)
	}
}
