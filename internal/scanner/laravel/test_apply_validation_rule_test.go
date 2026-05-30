//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyValidationRule 테스트
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyValidationRule(t *testing.T) {

	f := &scanner.Field{}
	num := false
	applyValidationRule(f, "integer", &num)
	if f.Type != "integer" || !num {
		t.Fatalf("type: %+v num=%v", f, num)
	}

	f2 := &scanner.Field{}
	num2 := false
	applyValidationRule(f2, "email", &num2)
	if f2.Type != "string" {
		t.Fatalf("format: %+v", f2)
	}

	f3 := &scanner.Field{}
	applyValidationRule(f3, "nullable", &num2)
	if !f3.Nullable {
		t.Fatalf("flag: %+v", f3)
	}

	f4 := &scanner.Field{}
	applyValidationRule(f4, "max:5", &num2)
	if f4.MaxLength == nil {
		t.Fatalf("constraint: %+v", f4)
	}
}
