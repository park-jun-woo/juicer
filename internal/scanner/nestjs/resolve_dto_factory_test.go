//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what resolveDTOFactory 테스트 (round5)
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
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
	// PartialType makes all fields optional
	for _, f := range got {
		if !f.optional {
			t.Errorf("field %s should be optional under PartialType", f.name)
		}
	}
}

func TestResolveDTOFactory_OmitType_Round5(t *testing.T) {
	b := []byte("class C extends OmitType(BaseDto, ['name']) {}")
	root, _ := parseTypeScript(b)
	calls := findAllByType(root, "call_expression")
	cache := map[string][]scanner.Field{
		"BaseDto": {
			{Name: "id", Type: "number"},
			{Name: "name", Type: "string"},
		},
	}
	got := resolveDTOFactory(calls[0], b, "/proj/c.dto.ts", map[string]string{}, "/proj", cache)
	for _, f := range got {
		if f.name == "name" {
			t.Errorf("name should be omitted: %+v", got)
		}
	}
}

func TestResolveDTOFactory_NoBase_Round5(t *testing.T) {
	// factory with no identifier base -> nil
	b := []byte("class C extends PartialType('x') {}")
	root, _ := parseTypeScript(b)
	calls := findAllByType(root, "call_expression")
	got := resolveDTOFactory(calls[0], b, "/proj/c.dto.ts", map[string]string{}, "/proj", map[string][]scanner.Field{})
	if got != nil {
		t.Fatalf("expected nil when base unresolved, got %+v", got)
	}
}
