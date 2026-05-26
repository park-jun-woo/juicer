//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveDTOFields_NoImport 테스트
package nestjs

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestResolveDTOFields_NoImport(t *testing.T) {
	cache := make(map[string][]scanner.Field)
	dr := dtoRequest{typeName: "Dto", imports: map[string]string{}}
	fields, err := resolveDTOFields(dr, cache)
	if err != nil || fields != nil {
		t.Fatalf("expected nil, got %v err=%v", fields, err)
	}
}
