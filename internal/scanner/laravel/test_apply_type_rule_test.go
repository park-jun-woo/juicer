//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyTypeRule 테스트
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyTypeRule(t *testing.T) {
	f := &scanner.Field{}
	isNum := false
	if !applyTypeRule(f, "integer", &isNum) || f.Type != "integer" || !isNum {
		t.Fatalf("integer: %+v num=%v", f, isNum)
	}
	f2 := &scanner.Field{}
	isNum2 := false
	if !applyTypeRule(f2, "string", &isNum2) || f2.Type != "string" || isNum2 {
		t.Fatalf("string: %+v num=%v", f2, isNum2)
	}
	f3 := &scanner.Field{}
	if applyTypeRule(f3, "unknown", &isNum2) {
		t.Fatal("unexpected handled for unknown type")
	}
}
