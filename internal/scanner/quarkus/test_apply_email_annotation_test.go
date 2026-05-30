//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestApplyEmailAnnotation 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyEmailAnnotation(t *testing.T) {
	field, src := firstFieldDecl(t, `class C { @Email private String email; }`)
	f := &scanner.Field{}
	applyEmailAnnotation(field, src, f)
	if f.Type != "string:email" {
		t.Fatalf("got %q", f.Type)
	}
}
