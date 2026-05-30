//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyMinRule 테스트
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyMinRule(t *testing.T) {
	f := &scanner.Field{}
	applyMinRule(f, "2", true)
	if f.Minimum == nil || *f.Minimum != 2 {
		t.Fatalf("number min: %+v", f)
	}
	f2 := &scanner.Field{}
	applyMinRule(f2, "3", false)
	if f2.MinLength == nil || *f2.MinLength != 3 {
		t.Fatalf("string minlen: %+v", f2)
	}
	f3 := &scanner.Field{}
	applyMinRule(f3, "x", false)
	if f3.MinLength != nil {
		t.Fatal("expected no change on parse error")
	}
}
