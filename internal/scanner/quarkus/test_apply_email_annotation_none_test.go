//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestApplyEmailAnnotation_None 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyEmailAnnotation_None(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { private String name; }`)
	f := &scanner.Field{}
	applyEmailAnnotation(field, src, f)
	if f.Type != "" {
		t.Fatalf("expected unchanged, got %q", f.Type)
	}
}
