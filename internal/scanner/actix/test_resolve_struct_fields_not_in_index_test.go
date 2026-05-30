//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestResolveStructFields_NotInIndex 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveStructFields_NotInIndex(t *testing.T) {
	idx := structIndex{}
	if f := resolveStructFields("Missing", idx, map[string][]scanner.Field{}); f != nil {
		t.Fatalf("expected nil for unknown type, got %+v", f)
	}
}
