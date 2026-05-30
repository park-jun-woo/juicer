//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what applyNumericConstraints 테스트
package fastify

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
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

func TestApplyNumericConstraints_None(t *testing.T) {
	obj, src := firstObject(t, `{ type: "integer" }`)
	f := &scanner.Field{}
	applyNumericConstraints(f, obj, src)
	if f.Minimum != nil || f.Maximum != nil || f.MinLength != nil || f.MaxLength != nil {
		t.Fatalf("expected all nil, got %v", f)
	}
}

func TestApplyNumericConstraints_NonNumeric(t *testing.T) {
	// minimum present but not a parseable int -> Atoi error, stays nil
	obj, src := firstObject(t, `{ type: "integer", minimum: "abc" }`)
	f := &scanner.Field{}
	applyNumericConstraints(f, obj, src)
	if f.Minimum != nil {
		t.Fatalf("expected nil for non-numeric, got %v", *f.Minimum)
	}
}
