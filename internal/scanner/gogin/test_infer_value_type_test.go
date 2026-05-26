//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestInferValueType 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestInferValueType(t *testing.T) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	tests := []struct {
		name string
		expr ast.Expr
		want string
	}{
		{"string lit", &ast.BasicLit{Kind: token.STRING, Value: `"hello"`}, "string"},
		{"int lit", &ast.BasicLit{Kind: token.INT, Value: "42"}, "integer"},
		{"float lit", &ast.BasicLit{Kind: token.FLOAT, Value: "3.14"}, "number"},
		{"true", &ast.Ident{Name: "true"}, "boolean"},
		{"false", &ast.Ident{Name: "false"}, "boolean"},
		{"nil", &ast.Ident{Name: "nil"}, "null"},
		{"unknown ident", &ast.Ident{Name: "x"}, "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inferValueType(tt.expr, info)
			if got != tt.want {
				t.Errorf("inferValueType() = %q, want %q", got, tt.want)
			}
		})
	}
}
