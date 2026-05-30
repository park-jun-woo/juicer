//ff:func feature=scan type=test control=sequence topic=express
//ff:what firstLexDecl 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstLexDecl(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	decls := findAllByType(fi.Root, "lexical_declaration")
	if len(decls) == 0 {
		t.Fatal("no lexical_declaration")
	}
	return decls[0]
}
