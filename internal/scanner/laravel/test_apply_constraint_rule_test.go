//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyConstraintRule 테스트
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyConstraintRule(t *testing.T) {
	f := &scanner.Field{}
	applyConstraintRule(f, "max:7", false)
	if f.MaxLength == nil || *f.MaxLength != 7 {
		t.Fatalf("max: %+v", f)
	}
	f2 := &scanner.Field{}
	applyConstraintRule(f2, "min:1", true)
	if f2.Minimum == nil || *f2.Minimum != 1 {
		t.Fatalf("min: %+v", f2)
	}
	f3 := &scanner.Field{}
	applyConstraintRule(f3, "in:a,b,c", false)
	if len(f3.Enum) != 3 || f3.Enum[0] != "a" {
		t.Fatalf("in: %+v", f3.Enum)
	}
	f4 := &scanner.Field{}
	applyConstraintRule(f4, "unrelated:x", false)
	if f4.MaxLength != nil || f4.Minimum != nil || len(f4.Enum) != 0 {
		t.Fatalf("expected no change: %+v", f4)
	}
}
