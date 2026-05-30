//ff:func feature=scan type=test control=sequence topic=express
//ff:what firstImportStmt 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func firstImportStmt(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	stmts := findAllByType(fi.Root, "import_statement")
	if len(stmts) == 0 {
		t.Fatal("no import_statement")
	}
	return stmts[0]
}
