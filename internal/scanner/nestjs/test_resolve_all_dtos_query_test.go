//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveAllDTOs_Query 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestResolveAllDTOs_Query(t *testing.T) {
	dir := t.TempDir()
	dtoFile := "dto.ts"
	writeFile(t, dir, dtoFile, `
export class PageOptionsDto {
  page: number;
  limit: number;
  q?: string;
}
`)
	imports := map[string]string{"PageOptionsDto": "./dto"}
	reqs := []dtoRequest{{
		typeName:    "PageOptionsDto",
		imports:     imports,
		referrer:    dir + "/controller.ts",
		projectRoot: dir,
		epIdx:       0,
		isQuery:     true,
	}}
	eps := []scanner.Endpoint{{Method: "GET", Path: "/users"}}
	resolveAllDTOs(reqs, eps)
	if eps[0].Request == nil {
		t.Fatal("expected non-nil Request")
	}
	if len(eps[0].Request.Query) != 3 {
		t.Fatalf("expected 3 query params, got %d", len(eps[0].Request.Query))
	}
	q := eps[0].Request.Query
	if q[0].Name != "page" {
		t.Errorf("param 0: want page, got %s", q[0].Name)
	}
	if q[1].Name != "limit" {
		t.Errorf("param 1: want limit, got %s", q[1].Name)
	}
	if q[2].Name != "q" {
		t.Errorf("param 2: want q, got %s", q[2].Name)
	}
}
