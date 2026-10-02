package propraven

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/jdw2111/propraven-go/internal/naming"
)

type specOp struct {
	httpMethod, path, group, method string
	paginated                       bool
}

func loadSpecOps(t *testing.T) []specOp {
	t.Helper()
	data, err := os.ReadFile("openapi.json")
	if err != nil {
		t.Fatal(err)
	}
	var spec struct {
		Paths map[string]map[string]json.RawMessage `json:"paths"`
	}
	if err := json.Unmarshal(data, &spec); err != nil {
		t.Fatal(err)
	}
	var ops []specOp
	for path, item := range spec.Paths {
		for m, raw := range item {
			switch m {
			case "get", "put", "post", "delete", "patch", "head", "options":
			default:
				continue
			}
			var op struct {
				Group      string          `json:"x-sdk-group"`
				Method     string          `json:"x-sdk-method"`
				Pagination json.RawMessage `json:"x-sdk-pagination"`
			}
			if err := json.Unmarshal(raw, &op); err != nil {
				t.Fatal(err)
			}
			ops = append(ops, specOp{strings.ToUpper(m), path, op.Group, op.Method, len(op.Pagination) > 0})
		}
	}
	sort.Slice(ops, func(i, j int) bool { return ops[i].path+ops[i].httpMethod < ops[j].path+ops[j].httpMethod })
	return ops
}

var pathVarRE = regexp.MustCompile(`\{[^}]+\}`)

// TestGeneratedSurface asserts every spec operation has a method on the
// right namespace, that calling it hits the spec's method and path, and that
// every paginated operation has an Iter sibling.
func TestGeneratedSurface(t *testing.T) {
	ops := loadSpecOps(t)
	if len(ops) == 0 {
		t.Fatal("no operations in openapi.json")
	}
	env := newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}, WithMaxRetries(0))
	cv := reflect.ValueOf(env.client).Elem()
	ctx := reflect.ValueOf(context.Background())
	for _, op := range ops {
		svc := cv.FieldByName(naming.Field(op.group))
		if !svc.IsValid() || svc.IsNil() {
			t.Errorf("%s %s: client.%s missing", op.httpMethod, op.path, naming.Field(op.group))
			continue
		}
		m := svc.MethodByName(naming.Method(op.method))
		if !m.IsValid() {
			t.Errorf("%s %s: client.%s.%s missing", op.httpMethod, op.path, naming.Field(op.group), naming.Method(op.method))
			continue
		}
		mt := m.Type()
		nPath := len(pathVarRE.FindAllString(op.path, -1))
		// ctx, path params..., params, opts...
		if mt.NumIn() != nPath+3 || !mt.IsVariadic() {
			t.Errorf("%s %s: signature %s", op.httpMethod, op.path, mt)
			continue
		}
		args := []reflect.Value{ctx}
		want := op.path
		for i := 0; i < nPath; i++ {
			v := "seg" + string(rune('A'+i))
			args = append(args, reflect.ValueOf(v))
			loc := pathVarRE.FindStringIndex(want)
			want = want[:loc[0]] + v + want[loc[1]:]
		}
		args = append(args, reflect.Zero(mt.In(nPath+1)))
		before := env.rec.count()
		m.Call(args)
		if env.rec.count() != before+1 {
			t.Errorf("%s %s: made %d requests", op.httpMethod, op.path, env.rec.count()-before)
			continue
		}
		got := env.rec.last()
		if got.Method != op.httpMethod || got.Path != want {
			t.Errorf("%s %s: sent %s %s, want %s", op.httpMethod, op.path, got.Method, got.Path, want)
		}
		if got.Header.Get("X-Payment") != "" {
			t.Errorf("%s %s: sent X-PAYMENT with nil params", op.httpMethod, op.path)
		}
		if op.paginated && !svc.MethodByName(naming.IterMethod(op.method)).IsValid() {
			t.Errorf("%s %s: paginated but client.%s.%s missing", op.httpMethod, op.path, naming.Field(op.group), naming.IterMethod(op.method))
		}
	}
	t.Logf("checked %d operations", len(ops))
}
