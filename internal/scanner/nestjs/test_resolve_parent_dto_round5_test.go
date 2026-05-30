//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveParentDTO_Round5 테스트
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
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
