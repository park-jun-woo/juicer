//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRouteChain_NotChain 테스트
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractRouteChain_NotChain(t *testing.T) {
	fi := mustParse(t, []byte(`foo();`))
	var first *sitter.Node = firstCallExpr(t, fi)
	if routes := extractRouteChain(first, fi.Src, map[string]bool{"router": true}); routes != nil {
		t.Fatalf("expected nil, got %+v", routes)
	}
}
