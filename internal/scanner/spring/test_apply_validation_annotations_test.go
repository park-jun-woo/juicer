//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestApplyValidationAnnotations 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyValidationAnnotations(t *testing.T) {
	field, src := firstField(t, `class C { @NotBlank @Min(1) private int x; }`)
	f := &scanner.Field{}
	applyValidationAnnotations(field, src, f)
	if f.Validate != "required" || f.Minimum == nil {
		t.Fatalf("got %+v", f)
	}
}
