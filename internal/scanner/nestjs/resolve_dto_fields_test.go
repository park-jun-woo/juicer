//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveDTOFields_Cached 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestResolveDTOFields_Cached(t *testing.T) {
	cache := map[string][]scanner.Field{
		"Dto": {{Name: "name", Type: "string"}},
	}
	dr := dtoRequest{typeName: "Dto"}
	fields, err := resolveDTOFields(dr, cache)
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 1 {
		t.Fatalf("expected 1, got %d", len(fields))
	}
}
