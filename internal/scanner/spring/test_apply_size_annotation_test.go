//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestApplySizeAnnotation 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplySizeAnnotation(t *testing.T) {
	field, src := firstField(t, `class C { @Size(min = 2, max = 8) private String s; }`)
	f := &scanner.Field{}
	applySizeAnnotation(field, src, f)
	if f.MinLength == nil || *f.MinLength != 2 || f.MaxLength == nil || *f.MaxLength != 8 {
		t.Fatalf("got %v %v", f.MinLength, f.MaxLength)
	}
}
