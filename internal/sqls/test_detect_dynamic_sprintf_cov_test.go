//ff:func feature=sql type=test control=sequence
//ff:what TestDetectDynamic_SprintfCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestDetectDynamic_SprintfCov(t *testing.T) {
	body := &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "fmt"}, Sel: &ast.Ident{Name: "Sprintf"}},
		}},
	}}
	if !detectDynamic(body) {
		t.Fatal("expected true for fmt.Sprintf")
	}
}
