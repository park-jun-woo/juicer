//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyFormat_StringWithFormat 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyFormat_StringWithFormat(t *testing.T) {
	obj, src := firstObject(t, `{ type: "string", format: "email" }`)
	f := &scanner.Field{Type: "string"}
	applyFormat(f, obj, src)
	if f.Type != "email" {
		t.Fatalf("expected email, got %q", f.Type)
	}
}
