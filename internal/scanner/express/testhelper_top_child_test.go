//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what topChild 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func topChild(t *testing.T, fi *fileInfo, typ string) *sitter.Node {
	t.Helper()
	root := fi.Root
	for i := 0; i < int(root.ChildCount()); i++ {
		if root.Child(i).Type() == typ {
			return root.Child(i)
		}
	}
	t.Fatalf("no top-level %s", typ)
	return nil
}
