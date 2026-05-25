//ff:func feature=sql type=parse control=sequence
//ff:what TestRefineCRUDFromAST_Empty 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestRefineCRUDFromAST_Empty(t *testing.T) {
	got := refineCRUDFromAST(&ast.BlockStmt{})
	if got != "EXEC" {
		t.Fatalf("expected EXEC, got %s", got)
	}
}
