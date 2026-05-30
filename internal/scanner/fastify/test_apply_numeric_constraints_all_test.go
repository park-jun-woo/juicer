//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplyNumericConstraints_All 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyNumericConstraints_All(t *testing.T) {
	obj, src := firstObject(t, `{ type: "integer", minimum: 1, maximum: 100, minLength: 2, maxLength: 50 }`)
	f := &scanner.Field{}
	applyNumericConstraints(f, obj, src)
	if f.Minimum == nil || *f.Minimum != 1 {
		t.Errorf("minimum = %v", f.Minimum)
	}
	if f.Maximum == nil || *f.Maximum != 100 {
		t.Errorf("maximum = %v", f.Maximum)
	}
	if f.MinLength == nil || *f.MinLength != 2 {
		t.Errorf("minLength = %v", f.MinLength)
	}
	if f.MaxLength == nil || *f.MaxLength != 50 {
		t.Errorf("maxLength = %v", f.MaxLength)
	}
}
