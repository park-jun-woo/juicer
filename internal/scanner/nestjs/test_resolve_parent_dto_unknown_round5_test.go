//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveParentDTO_Unknown_Round5 테스트
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveParentDTO_Unknown_Round5(t *testing.T) {
	got := resolveParentDTO("Missing", "/proj/child.dto.ts", map[string]string{}, "/proj", map[string][]scanner.Field{})
	if got != nil {
		t.Fatalf("expected nil for unknown parent, got %+v", got)
	}
}
