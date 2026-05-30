//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyFlagRule 테스트
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyFlagRule(t *testing.T) {
	f := &scanner.Field{}
	if !applyFlagRule(f, "nullable") || !f.Nullable {
		t.Fatal("nullable not applied")
	}
	f2 := &scanner.Field{}
	if !applyFlagRule(f2, "required") || f2.Validate != "required" {
		t.Fatalf("required not applied: %+v", f2)
	}
	f3 := &scanner.Field{}
	if applyFlagRule(f3, "email") {
		t.Fatal("unexpected handled for email")
	}
}
