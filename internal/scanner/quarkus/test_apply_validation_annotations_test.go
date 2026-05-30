//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestApplyValidationAnnotations 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyValidationAnnotations(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @NotNull @Min(1) private int x; }`)
	f := &scanner.Field{}
	applyValidationAnnotations(field, src, f)
	if f.Validate != "required" {
		t.Fatalf("validate: %q", f.Validate)
	}
	if f.Minimum == nil || *f.Minimum != 1 {
		t.Fatalf("min: %v", f.Minimum)
	}
}
