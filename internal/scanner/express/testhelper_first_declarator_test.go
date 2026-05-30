//ff:func feature=scan type=test control=sequence topic=express
//ff:what firstDeclarator 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstDeclarator(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	ds := findAllByType(fi.Root, "variable_declarator")
	if len(ds) == 0 {
		t.Fatal("no variable_declarator")
	}
	return ds[0]
}
