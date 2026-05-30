//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyEnum_NotArray 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyEnum_NotArray(t *testing.T) {

	obj, src := firstObject(t, `{ type: "string", enum: "x" }`)
	f := &scanner.Field{}
	applyEnum(f, obj, src)
	if len(f.Enum) != 0 {
		t.Fatalf("expected no enum when not array, got %v", f.Enum)
	}
}
