//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestResolveStructFields_CacheHit 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveStructFields_CacheHit(t *testing.T) {
	cache := map[string][]scanner.Field{
		"X": {{Name: "a", Type: "string"}},
	}
	fields := resolveStructFields("X", nil, cache)
	if len(fields) != 1 || fields[0].Name != "a" {
		t.Fatalf("expected cached fields, got %+v", fields)
	}
}
