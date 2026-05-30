//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveParentDTOFields_NoImport 테스트
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveParentDTOFields_NoImport(t *testing.T) {
	got := resolveParentDTOFields("Unknown", "x.ts", map[string]string{}, "", map[string][]scanner.Field{})
	if got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}
