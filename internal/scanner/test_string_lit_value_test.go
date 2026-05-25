//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestStringLitValue 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestStringLitValue(t *testing.T) {
	tests := []struct {
		name string
		expr ast.Expr
		want string
	}{
		{
			name: "string literal",
			expr: &ast.BasicLit{Kind: token.STRING, Value: `"hello"`},
			want: "hello",
		},
		{
			name: "non-string literal",
			expr: &ast.BasicLit{Kind: token.INT, Value: "42"},
			want: "",
		},
		{
			name: "non-literal",
			expr: &ast.Ident{Name: "x"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringLitValue(tt.expr)
			if got != tt.want {
				t.Errorf("stringLitValue() = %q, want %q", got, tt.want)
			}
		})
	}
}
