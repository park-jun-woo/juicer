//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestApplyMinMaxAnnotation 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyMinMaxAnnotation(t *testing.T) {
	field, src := firstField(t, `class C { @Min(1) @Max(10) private int x; }`)
	f := &scanner.Field{}
	applyMinAnnotation(field, src, f)
	applyMaxAnnotation(field, src, f)
	if f.Minimum == nil || *f.Minimum != 1 || f.Maximum == nil || *f.Maximum != 10 {
		t.Fatalf("got min=%v max=%v", f.Minimum, f.Maximum)
	}
}
