//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what applyNestedFields 테스트
package fastify

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyNestedFields_Object(t *testing.T) {
	obj, src := firstObject(t, `{ type: "object", properties: { id: { type: "integer" } } }`)
	f := &scanner.Field{}
	applyNestedFields(f, "object", obj, src)
	if len(f.Fields) != 1 || f.Fields[0].Name != "id" {
		t.Fatalf("expected nested id field, got %v", f.Fields)
	}
}

func TestApplyNestedFields_Array(t *testing.T) {
	obj, src := firstObject(t, `{ type: "array", items: { type: "string" } }`)
	f := &scanner.Field{}
	applyNestedFields(f, "array", obj, src)
	if f.Type != "string[]" {
		t.Fatalf("expected string[], got %q", f.Type)
	}
}

func TestApplyNestedFields_Scalar(t *testing.T) {
	// neither object nor array -> no change
	obj, src := firstObject(t, `{ type: "string" }`)
	f := &scanner.Field{Type: "string"}
	applyNestedFields(f, "string", obj, src)
	if f.Type != "string" || len(f.Fields) != 0 {
		t.Fatalf("expected unchanged for scalar, got %v", f)
	}
}
