//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what resolveParentDTO 테스트 (round5)
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestResolveParentDTO_Round5(t *testing.T) {
	cache := map[string][]scanner.Field{
		"BaseDto": {
			{Name: "id", Type: "number", Validate: "required"},
			{Name: "name", Type: "string"},
		},
	}
	got := resolveParentDTO("BaseDto", "/proj/child.dto.ts", map[string]string{}, "/proj", cache)
	if len(got) != 2 {
		t.Fatalf("expected 2 fields, got %d: %+v", len(got), got)
	}
	if got[0].name != "id" || got[1].name != "name" {
		t.Fatalf("unexpected fields: %+v", got)
	}
}

func TestResolveParentDTO_Unknown_Round5(t *testing.T) {
	got := resolveParentDTO("Missing", "/proj/child.dto.ts", map[string]string{}, "/proj", map[string][]scanner.Field{})
	if got != nil {
		t.Fatalf("expected nil for unknown parent, got %+v", got)
	}
}
