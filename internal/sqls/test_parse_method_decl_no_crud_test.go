//ff:func feature=sql type=test control=sequence
//ff:what TestParseMethodDecl_NoCRUD 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestParseMethodDecl_NoCRUD(t *testing.T) {
	fn := &ast.FuncDecl{
		Recv: &ast.FieldList{List: []*ast.Field{
			{Type: &ast.StarExpr{X: &ast.Ident{Name: "UserRepo"}}},
		}},
		Name: &ast.Ident{Name: "GetUser"},
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{},
	}
	if parseMethodDecl(fn) != nil {
		t.Fatal("expected nil for no DB calls")
	}
}
