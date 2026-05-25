//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestExprString 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExprString(t *testing.T) {
	tests := []struct {
		name string
		expr ast.Expr
		want string
	}{
		{
			name: "ident",
			expr: &ast.Ident{Name: "foo"},
			want: "foo",
		},
		{
			name: "selector",
			expr: &ast.SelectorExpr{X: &ast.Ident{Name: "pkg"}, Sel: &ast.Ident{Name: "Func"}},
			want: "pkg.Func",
		},
		{
			name: "composite lit with type",
			expr: &ast.CompositeLit{Type: &ast.Ident{Name: "Foo"}},
			want: "Foo{}",
		},
		{
			name: "composite lit without type",
			expr: &ast.CompositeLit{},
			want: "{}",
		},
		{
			name: "star expr",
			expr: &ast.StarExpr{X: &ast.Ident{Name: "foo"}},
			want: "*foo",
		},
		{
			name: "unary expr",
			expr: &ast.UnaryExpr{X: &ast.Ident{Name: "foo"}},
			want: "foo",
		},
		{
			name: "call expr",
			expr: &ast.CallExpr{Fun: &ast.Ident{Name: "foo"}},
			want: "foo()",
		},
		{
			name: "index expr",
			expr: &ast.IndexExpr{X: &ast.Ident{Name: "slice"}, Index: &ast.Ident{Name: "i"}},
			want: "slice[i]",
		},
		{
			name: "map type",
			expr: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "int"}},
			want: "map[string]int",
		},
		{
			name: "array type",
			expr: &ast.ArrayType{Elt: &ast.Ident{Name: "int"}},
			want: "[]int",
		},
		{
			name: "interface type",
			expr: &ast.InterfaceType{Methods: &ast.FieldList{}},
			want: "interface{}",
		},
		{
			name: "basic lit",
			expr: &ast.BasicLit{Kind: token.STRING, Value: `"hello"`},
			want: `"hello"`,
		},
		{
			name: "unknown type",
			expr: &ast.Ellipsis{},
			want: "(*ast.Ellipsis)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exprString(tt.expr)
			if got != tt.want {
				t.Errorf("exprString() = %q, want %q", got, tt.want)
			}
		})
	}
}
