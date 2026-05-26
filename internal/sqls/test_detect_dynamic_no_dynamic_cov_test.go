//ff:func feature=sql type=test control=sequence
//ff:what TestDetectDynamic_NoDynamicCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestDetectDynamic_NoDynamicCov(t *testing.T) {
	body := &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "x"}},
	}}
	if detectDynamic(body) {
		t.Fatal("expected false")
	}
}
