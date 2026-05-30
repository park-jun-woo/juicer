//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestApplyMinAnnotation_ValueArg 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyMinAnnotation_ValueArg(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Min(value = 3) private int x; }`)
	f := &scanner.Field{}
	applyMinAnnotation(field, src, f)
	if f.Minimum == nil || *f.Minimum != 3 {
		t.Fatalf("got %v", f.Minimum)
	}
}
