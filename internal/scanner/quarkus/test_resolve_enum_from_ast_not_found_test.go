//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestResolveEnumFromAST_NotFound 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveEnumFromAST_NotFound(t *testing.T) {
	src := []byte(`enum Status { OPEN }`)
	root, _ := parseJava(src)
	if got := resolveEnumFromAST(root, src, "Other", map[string][]scanner.Field{}); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}
