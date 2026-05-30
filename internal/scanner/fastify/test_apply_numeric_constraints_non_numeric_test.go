//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyNumericConstraints_NonNumeric 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyNumericConstraints_NonNumeric(t *testing.T) {

	obj, src := firstObject(t, `{ type: "integer", minimum: "abc" }`)
	f := &scanner.Field{}
	applyNumericConstraints(f, obj, src)
	if f.Minimum != nil {
		t.Fatalf("expected nil for non-numeric, got %v", *f.Minimum)
	}
}
