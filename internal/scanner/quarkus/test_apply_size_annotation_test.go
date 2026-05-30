//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestApplySizeAnnotation 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplySizeAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Size(min = 1, max = 10) private String s; }`)
	f := &scanner.Field{}
	applySizeAnnotation(field, src, f)
	if f.MinLength == nil || *f.MinLength != 1 {
		t.Fatalf("minLen: %v", f.MinLength)
	}
	if f.MaxLength == nil || *f.MaxLength != 10 {
		t.Fatalf("maxLen: %v", f.MaxLength)
	}
}
