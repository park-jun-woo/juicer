//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyFormatRule 테스트
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyFormatRule(t *testing.T) {
	f := &scanner.Field{}
	if !applyFormatRule(f, "email") {
		t.Fatal("email not handled")
	}
	if f.Type != "string" || f.Validate != "email" {
		t.Fatalf("got %+v", f)
	}
	f2 := &scanner.Field{}
	if applyFormatRule(f2, "notaformat") {
		t.Fatal("unexpected handled")
	}
}
