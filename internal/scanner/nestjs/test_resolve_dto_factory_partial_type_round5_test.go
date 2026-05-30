//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestResolveDTOFactory_PartialType_Round5 테스트
package nestjs

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveDTOFactory_PartialType_Round5(t *testing.T) {
	b := []byte("class C extends PartialType(BaseDto) {}")
	root, _ := parseTypeScript(b)
	calls := findAllByType(root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no call_expression")
	}
	cache := map[string][]scanner.Field{
		"BaseDto": {
			{Name: "id", Type: "number", Validate: "required"},
			{Name: "name", Type: "string", Validate: "required"},
		},
	}
	got := resolveDTOFactory(calls[0], b, "/proj/c.dto.ts", map[string]string{}, "/proj", cache)
	if len(got) != 2 {
		t.Fatalf("expected 2 fields, got %d: %+v", len(got), got)
	}

	for _, f := range got {
		if !f.optional {
			t.Errorf("field %s should be optional under PartialType", f.name)
		}
	}
}
