//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyMaxRule 테스트
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyMaxRule(t *testing.T) {
	f := &scanner.Field{}
	applyMaxRule(f, "10", true)
	if f.Maximum == nil || *f.Maximum != 10 {
		t.Fatalf("number max: %+v", f)
	}
	f2 := &scanner.Field{}
	applyMaxRule(f2, "5", false)
	if f2.MaxLength == nil || *f2.MaxLength != 5 {
		t.Fatalf("string maxlen: %+v", f2)
	}
	f3 := &scanner.Field{}
	applyMaxRule(f3, "notanum", true)
	if f3.Maximum != nil {
		t.Fatal("expected no change on parse error")
	}
}
