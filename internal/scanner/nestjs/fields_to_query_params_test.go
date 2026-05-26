//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFieldsToQueryParams 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestFieldsToQueryParams(t *testing.T) {
	fields := []scanner.Field{
		{Name: "page", Type: "number"},
		{Name: "q", Type: "string"},
		{Name: "noType"},
	}
	params := fieldsToQueryParams(fields)
	if len(params) != 3 {
		t.Fatalf("expected 3, got %d", len(params))
	}
	if params[0].Name != "page" || params[0].Type != "number" {
		t.Fatalf("param 0: want page/number, got %s/%s", params[0].Name, params[0].Type)
	}
	if params[2].Type != "string" {
		t.Fatalf("empty type should default to string, got %q", params[2].Type)
	}
}
