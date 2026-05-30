//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveEnumFromAST 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveEnumFromAST(t *testing.T) {
	src := []byte(`enum Status { OPEN, CLOSED }`)
	root, _ := parseJava(src)
	fields := resolveEnumFromAST(root, src, "Status", map[string][]scanner.Field{})
	if len(fields) != 1 || len(fields[0].Enum) != 2 {
		t.Fatalf("got %+v", fields)
	}
}
