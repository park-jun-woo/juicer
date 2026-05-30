//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyNestedFields_Scalar 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyNestedFields_Scalar(t *testing.T) {

	obj, src := firstObject(t, `{ type: "string" }`)
	f := &scanner.Field{Type: "string"}
	applyNestedFields(f, "string", obj, src)
	if f.Type != "string" || len(f.Fields) != 0 {
		t.Fatalf("expected unchanged for scalar, got %v", f)
	}
}
