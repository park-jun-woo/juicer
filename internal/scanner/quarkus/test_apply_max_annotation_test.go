//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestApplyMaxAnnotation 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyMaxAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Max(100) private int x; }`)
	f := &scanner.Field{}
	applyMaxAnnotation(field, src, f)
	if f.Maximum == nil || *f.Maximum != 100 {
		t.Fatalf("got %v", f.Maximum)
	}
}
