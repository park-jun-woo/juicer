//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyArrayItems_NoItems 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyArrayItems_NoItems(t *testing.T) {
	obj, src := firstObject(t, `{ type: "array" }`)
	f := &scanner.Field{Type: "orig"}
	applyArrayItems(f, obj, src)
	if f.Type != "orig" || len(f.Fields) != 0 {
		t.Fatalf("expected unchanged when items missing, got %v", f)
	}
}
