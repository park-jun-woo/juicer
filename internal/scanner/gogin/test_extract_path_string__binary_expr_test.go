//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractPathString_BinaryExpr 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractPathString_BinaryExpr(t *testing.T) {
	expr := &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"/api"`},
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"/v1"`},
	}
	path, ok := extractPathString(expr)
	if !ok || path != "/api/v1" {
		t.Fatalf("expected /api/v1, got %s ok=%v", path, ok)
	}
}
