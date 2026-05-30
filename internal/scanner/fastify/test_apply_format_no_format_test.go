//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyFormat_NoFormat 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyFormat_NoFormat(t *testing.T) {
	obj, src := firstObject(t, `{ type: "string" }`)
	f := &scanner.Field{Type: "string"}
	applyFormat(f, obj, src)
	if f.Type != "string" {
		t.Fatalf("expected unchanged string, got %q", f.Type)
	}
}
