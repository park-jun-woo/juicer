//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyNestedFields_Object 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyNestedFields_Object(t *testing.T) {
	obj, src := firstObject(t, `{ type: "object", properties: { id: { type: "integer" } } }`)
	f := &scanner.Field{}
	applyNestedFields(f, "object", obj, src)
	if len(f.Fields) != 1 || f.Fields[0].Name != "id" {
		t.Fatalf("expected nested id field, got %v", f.Fields)
	}
}
