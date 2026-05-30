//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyEnum_NoEnum 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyEnum_NoEnum(t *testing.T) {
	obj, src := firstObject(t, `{ type: "string" }`)
	f := &scanner.Field{}
	applyEnum(f, obj, src)
	if len(f.Enum) != 0 {
		t.Fatalf("expected no enum, got %v", f.Enum)
	}
}
