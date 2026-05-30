//ff:func feature=scan type=test control=sequence topic=express
//ff:what exFirst 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func exFirst(t *testing.T, fi *fileInfo, typ string) *sitter.Node {
	t.Helper()
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == typ {
			found = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(fi.Root)
	if found == nil {
		t.Fatalf("no %s node", typ)
	}
	return found
}
