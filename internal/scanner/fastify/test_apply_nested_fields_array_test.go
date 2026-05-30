//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyNestedFields_Array 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyNestedFields_Array(t *testing.T) {
	obj, src := firstObject(t, `{ type: "array", items: { type: "string" } }`)
	f := &scanner.Field{}
	applyNestedFields(f, "array", obj, src)
	if f.Type != "string[]" {
		t.Fatalf("expected string[], got %q", f.Type)
	}
}
