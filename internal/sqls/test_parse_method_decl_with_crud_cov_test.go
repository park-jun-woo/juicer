//ff:func feature=sql type=test control=sequence
//ff:what TestParseMethodDecl_WithCRUDCov 테스트
package sqls

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestParseMethodDecl_WithCRUDCov(t *testing.T) {
	fn := &ast.FuncDecl{
		Recv: &ast.FieldList{List: []*ast.Field{
			{Type: &ast.StarExpr{X: &ast.Ident{Name: "UserRepo"}}},
		}},
		Name: &ast.Ident{Name: "GetUser"},
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{List: []ast.Stmt{
			&ast.ExprStmt{X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "db"}, Sel: &ast.Ident{Name: "QueryContext"}},
				Args: []ast.Expr{
					&ast.Ident{Name: "ctx"},
					&ast.BasicLit{Kind: token.STRING, Value: "`SELECT id, name FROM users WHERE id = $1`"},
				},
			}},
		}},
	}
	sk := parseMethodDecl(fn)
	if sk == nil {
		t.Fatal("expected non-nil skeleton")
	}
	if sk.CRUD != "SELECT" {
		t.Fatalf("expected SELECT, got %s", sk.CRUD)
	}
}
