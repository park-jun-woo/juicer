//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what TestTypeString 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString(t *testing.T) {
	tests := []struct {
		name string
		expr ast.Expr
		want string
	}{
		{"ident", &ast.Ident{Name: "string"}, "string"},
		{"star", &ast.StarExpr{X: &ast.Ident{Name: "User"}}, "*User"},
		{"selector", &ast.SelectorExpr{X: &ast.Ident{Name: "sql"}, Sel: &ast.Ident{Name: "DB"}}, "sql.DB"},
		{"array", &ast.ArrayType{Elt: &ast.Ident{Name: "int"}}, "[]int"},
		{"map", &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "int"}}, "map[string]int"},
		{"interface", &ast.InterfaceType{}, "interface{}"},
		{"ellipsis", &ast.Ellipsis{Elt: &ast.Ident{Name: "string"}}, "...string"},
		{"unknown", &ast.FuncType{}, "*ast.FuncType"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := typeString(tt.expr)
			if got != tt.want {
				t.Errorf("typeString() = %q, want %q", got, tt.want)
			}
		})
	}
}
