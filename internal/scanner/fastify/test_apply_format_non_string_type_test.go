//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyFormat_NonStringType 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyFormat_NonStringType(t *testing.T) {

	obj, src := firstObject(t, `{ type: "integer", format: "int64" }`)
	f := &scanner.Field{Type: "integer"}
	applyFormat(f, obj, src)
	if f.Type != "integer" {
		t.Fatalf("expected unchanged integer, got %q", f.Type)
	}
}
