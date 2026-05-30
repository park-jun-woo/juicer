//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestApplyJsonProperty_None 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyJsonProperty_None(t *testing.T) {
	root, _ := parseJava([]byte(`class D { private String name; }`))
	src := []byte(`class D { private String name; }`)
	field := findAllByType(root, "field_declaration")[0]
	f := &scanner.Field{}
	applyJsonProperty(field, src, f)
	if f.JSON != "" {
		t.Fatalf("expected empty, got %q", f.JSON)
	}
}
