//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyEnum_Array 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyEnum_Array(t *testing.T) {
	obj, src := firstObject(t, `{ type: "string", enum: ["a", "b", "c"] }`)
	f := &scanner.Field{}
	applyEnum(f, obj, src)
	if len(f.Enum) != 3 || f.Enum[0] != "a" || f.Enum[2] != "c" {
		t.Fatalf("expected [a b c], got %v", f.Enum)
	}
}
