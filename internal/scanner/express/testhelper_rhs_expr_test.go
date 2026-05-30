//ff:func feature=scan type=test control=sequence topic=express
//ff:what rhsExpr 테스트 헬퍼
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

// rhsExpr returns the right-hand side expression node of `x = <expr>;`.
func rhsExpr(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	assign := exFirst(t, fi, "assignment_expression")
	return assign.Child(int(assign.ChildCount()) - 1)
}
