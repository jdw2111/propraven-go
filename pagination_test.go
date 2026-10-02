package propraven

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

// offsetServer serves n rows; withTotal adds "total", hasMore adds "has_more".
func offsetServer(n int, withTotal, hasMore bool) func(w http.ResponseWriter, r *http.Request, attempt int) {
	return func(w http.ResponseWriter, r *http.Request, attempt int) {
		q := r.URL.Query()
		limit, _ := strconv.Atoi(q.Get("limit"))
		offset, _ := strconv.Atoi(q.Get("offset"))
		var rows []string
		for i := offset; i < offset+limit && i < n; i++ {
			rows = append(rows, fmt.Sprintf(`{"parcel_id":"%d","county_fips":"37119"}`, i))
		}
		body := `{"data":[` + strings.Join(rows, ",") + `],"limit":` + strconv.Itoa(limit) + `,"offset":` + strconv.Itoa(offset)
		if withTotal {
			body += `,"total":` + strconv.Itoa(n)
		}
		if hasMore {
			body += `,"has_more":` + strconv.FormatBool(offset+limit < n)
		}
		jsonReply(200, body+"}")(w, r, attempt)
	}
}

func collect[T any](t *testing.T, it *Iter[T]) []T {
	t.Helper()
	var out []T
	for it.Next() {
		out = append(out, it.Current())
	}
	if err := it.Err(); err != nil {
		t.Fatal(err)
	}
	return out
}

func TestOffsetPaginationStopsOnShortPage(t *testing.T) {
	env := newEnv(t, offsetServer(25, false, false))
	it := env.client.Deals.AbsenteeIter(context.Background(), &DealsAbsenteeParams{CountyFIPS: String("37119")}, IterOptions{PageSize: 10})
	rows := collect(t, it)
	if len(rows) != 25 || it.Pages() != 3 || env.rec.count() != 3 {
		t.Fatalf("rows %d pages %d requests %d", len(rows), it.Pages(), env.rec.count())
	}
	if rows[0].ParcelID != "0" || rows[24].ParcelID != "24" {
		t.Fatalf("rows %v .. %v", rows[0], rows[24])
	}
	var offsets []string
	for _, r := range env.rec.requests {
		q, _ := url.ParseQuery(r.Query)
		if q.Get("county_fips") != "37119" || q.Get("limit") != "10" {
			t.Fatalf("query %q", r.Query)
		}
		offsets = append(offsets, q.Get("offset"))
	}
	if strings.Join(offsets, ",") != "0,10,20" {
		t.Fatalf("offsets %v", offsets)
	}
}

func TestOffsetPaginationStopsOnTotal(t *testing.T) {
	// 30 rows in pages of 10: the third page is full, but offset reaches
	// total, so no fourth (empty) request is made.
	env := newEnv(t, offsetServer(30, true, false))
	rows := collect(t, env.client.Deals.AbsenteeIter(context.Background(), nil, IterOptions{PageSize: 10}))
	if len(rows) != 30 || env.rec.count() != 3 {
		t.Fatalf("rows %d requests %d", len(rows), env.rec.count())
	}
}

func TestOffsetPaginationStopsOnHasMore(t *testing.T) {
	env := newEnv(t, offsetServer(20, false, true))
	rows := collect(t, env.client.Deals.AbsenteeIter(context.Background(), nil, IterOptions{PageSize: 10}))
	if len(rows) != 20 || env.rec.count() != 2 {
		t.Fatalf("rows %d requests %d", len(rows), env.rec.count())
	}
}

func TestOffsetPaginationMaxItemsAndStartOffset(t *testing.T) {
	env := newEnv(t, offsetServer(100, true, false))
	rows := collect(t, env.client.Deals.AbsenteeIter(context.Background(), &DealsAbsenteeParams{Offset: Int(5), Limit: Int(4)}, IterOptions{MaxItems: 6}))
	if len(rows) != 6 || env.rec.count() != 2 || rows[0].ParcelID != "5" || rows[5].ParcelID != "10" {
		t.Fatalf("rows %d requests %d first %v", len(rows), env.rec.count(), rows)
	}
}

func TestOffsetPaginationFollowsEchoedLimit(t *testing.T) {
	// The server clamps limit 1000 to 3 and echoes it; the iterator must
	// not mistake a clamped page for a short (final) one.
	env := newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		q := r.URL.Query()
		offset, _ := strconv.Atoi(q.Get("offset"))
		var rows []string
		for i := offset; i < offset+3 && i < 7; i++ {
			rows = append(rows, fmt.Sprintf(`{"parcel_id":"%d"}`, i))
		}
		jsonReply(200, `{"data":[`+strings.Join(rows, ",")+`],"limit":3,"offset":`+strconv.Itoa(offset)+`}`)(w, r, attempt)
	})
	rows := collect(t, env.client.Deals.AbsenteeIter(context.Background(), nil, IterOptions{PageSize: 1000}))
	if len(rows) != 7 || env.rec.count() != 3 {
		t.Fatalf("rows %d requests %d", len(rows), env.rec.count())
	}
}

func TestOffsetPaginationInJSONBody(t *testing.T) {
	var env *testEnv
	env = newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		var body struct {
			Limit  int `json:"limit"`
			Offset int `json:"offset"`
		}
		_ = json.Unmarshal(env.rec.last().Body, &body)
		var rows []string
		for i := body.Offset; i < body.Offset+body.Limit && i < 5; i++ {
			rows = append(rows, fmt.Sprintf(`{"parcel_id":"%d"}`, i))
		}
		jsonReply(200, `{"data":[`+strings.Join(rows, ",")+`],"total":5,"has_more":`+strconv.FormatBool(body.Offset+body.Limit < 5)+`,"limit":`+strconv.Itoa(body.Limit)+`,"offset":`+strconv.Itoa(body.Offset)+`}`)(w, r, attempt)
	})
	rows := collect(t, env.client.Search.ParcelsIter(context.Background(),
		&SearchParcelsParams{Bounds: &SearchParcelsParamsBounds{North: 1, South: 0, East: 1, West: 0}}, IterOptions{PageSize: 2}))
	if len(rows) != 5 || env.rec.count() != 3 {
		t.Fatalf("rows %d requests %d", len(rows), env.rec.count())
	}
	last := string(env.rec.last().Body)
	if !strings.Contains(last, `"limit":2`) || !strings.Contains(last, `"offset":4`) || !strings.Contains(last, `"bounds"`) {
		t.Fatalf("last body %s", last)
	}
	if env.rec.last().Query != "" {
		t.Fatalf("POST pagination leaked into the query: %q", env.rec.last().Query)
	}
}

func TestCursorPagination(t *testing.T) {
	pages := map[string]string{
		"":   `{"results":[{"parcel_id":"a"},{"parcel_id":"b"}],"hasMore":true,"nextCursor":"c1","total":5,"page":1,"pages":3}`,
		"c1": `{"results":[{"parcel_id":"c"},{"parcel_id":"d"}],"hasMore":true,"nextCursor":"c2","total":5,"page":2,"pages":3}`,
		"c2": `{"results":[{"parcel_id":"e"}],"hasMore":false,"nextCursor":null,"total":5,"page":3,"pages":3}`,
	}
	env := newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		jsonReply(200, pages[r.URL.Query().Get("after")])(w, r, attempt)
	})
	it := env.client.Search.FullIter(context.Background(), &SearchFullParams{Q: String("main st")}, IterOptions{PageSize: 2})
	rows := collect(t, it)
	var ids []string
	for _, r := range rows {
		ids = append(ids, r.ParcelID)
	}
	if strings.Join(ids, "") != "abcde" || env.rec.count() != 3 {
		t.Fatalf("ids %v requests %d", ids, env.rec.count())
	}
	q, _ := url.ParseQuery(env.rec.requests[1].Query)
	if q.Get("after") != "c1" || q.Get("limit") != "2" || q.Get("q") != "main st" {
		t.Fatalf("second query %q", env.rec.requests[1].Query)
	}
	if strings.Contains(env.rec.requests[0].Query, "after=") {
		t.Fatalf("first request must not send a cursor: %q", env.rec.requests[0].Query)
	}
}

func TestCursorPaginationStopsWhenCursorAbsent(t *testing.T) {
	env := newEnv(t, jsonReply(200, `{"results":[{"parcel_id":"a"}],"total":1}`))
	rows := collect(t, env.client.Search.FullIter(context.Background(), &SearchFullParams{Q: String("zz")}, IterOptions{}))
	if len(rows) != 1 || env.rec.count() != 1 {
		t.Fatalf("rows %d requests %d", len(rows), env.rec.count())
	}
}

func TestIteratorSurfacesErrors(t *testing.T) {
	env := newEnv(t, func(w http.ResponseWriter, r *http.Request, attempt int) {
		if r.URL.Query().Get("offset") == "0" {
			offsetServer(10, false, false)(w, r, attempt)
			return
		}
		jsonReply(403, `{"code":"forbidden","detail":"no"}`)(w, r, attempt)
	})
	it := env.client.Deals.AbsenteeIter(context.Background(), nil, IterOptions{PageSize: 5})
	n := 0
	for it.Next() {
		n++
	}
	if n != 5 || !IsPermissionDenied(it.Err()) {
		t.Fatalf("n %d err %v", n, it.Err())
	}
	if it.Next() {
		t.Fatal("Next after error must be false")
	}
}
