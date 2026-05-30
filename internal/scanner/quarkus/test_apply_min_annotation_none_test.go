//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestApplyMinAnnotation_None 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyMinAnnotation_None(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { private int x; }`)
	f := &scanner.Field{}
	applyMinAnnotation(field, src, f)
	if f.Minimum != nil {
		t.Fatal("expected nil")
	}
}
