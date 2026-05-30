//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestApplyMinAnnotation 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyMinAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Min(5) private int x; }`)
	f := &scanner.Field{}
	applyMinAnnotation(field, src, f)
	if f.Minimum == nil || *f.Minimum != 5 {
		t.Fatalf("got %v", f.Minimum)
	}
}
