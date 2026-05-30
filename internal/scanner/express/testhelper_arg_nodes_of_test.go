//ff:func feature=scan type=test control=sequence topic=express
//ff:what argNodesOf 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func argNodesOf(t *testing.T, fi *fileInfo) []*sitter.Node {
	t.Helper()
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	if args == nil {
		t.Fatal("no arguments")
	}
	return collectArgNodes(args)
}
