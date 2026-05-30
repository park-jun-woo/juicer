//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestResolveParentFields_NoSuper 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveParentFields_NoSuper(t *testing.T) {
	src := []byte(`class C { String x; }`)
	root, _ := parseJava(src)
	cls := findAllByType(root, "class_declaration")[0]
	if got := resolveParentFields(cls, src, "/abs/C.java", "/abs", map[string]string{}, map[string][]scanner.Field{}); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}
