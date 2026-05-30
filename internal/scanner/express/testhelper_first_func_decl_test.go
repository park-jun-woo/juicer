//ff:func feature=scan type=test control=sequence topic=express
//ff:what firstFuncDecl 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstFuncDecl(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	fns := findAllByType(fi.Root, "function_declaration")
	if len(fns) == 0 {
		t.Fatal("no function_declaration")
	}
	return fns[0]
}
