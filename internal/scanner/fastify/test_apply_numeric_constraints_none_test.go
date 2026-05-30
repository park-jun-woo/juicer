//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyNumericConstraints_None 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyNumericConstraints_None(t *testing.T) {
	obj, src := firstObject(t, `{ type: "integer" }`)
	f := &scanner.Field{}
	applyNumericConstraints(f, obj, src)
	if f.Minimum != nil || f.Maximum != nil || f.MinLength != nil || f.MaxLength != nil {
		t.Fatalf("expected all nil, got %v", f)
	}
}
