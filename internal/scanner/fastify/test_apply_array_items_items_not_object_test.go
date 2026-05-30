//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyArrayItems_ItemsNotObject 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyArrayItems_ItemsNotObject(t *testing.T) {

	obj, src := firstObject(t, `{ type: "array", items: "nope" }`)
	f := &scanner.Field{Type: "orig"}
	applyArrayItems(f, obj, src)
	if f.Type != "orig" {
		t.Fatalf("expected unchanged when items not object, got %q", f.Type)
	}
}
