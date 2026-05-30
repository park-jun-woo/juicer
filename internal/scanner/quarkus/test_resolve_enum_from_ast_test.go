//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestResolveEnumFromAST 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveEnumFromAST(t *testing.T) {
	src := []byte(`enum Status { OPEN, CLOSED }`)
	root, _ := parseJava(src)
	cache := map[string][]scanner.Field{}
	fields := resolveEnumFromAST(root, src, "Status", cache)
	if len(fields) != 1 || len(fields[0].Enum) != 2 {
		t.Fatalf("got %+v", fields)
	}
}
