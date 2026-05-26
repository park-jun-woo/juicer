//ff:func feature=scan type=test control=sequence
//ff:what TestCollectStringParts_BinaryExprCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestCollectStringParts_BinaryExprCov(t *testing.T) {
	var parts []string
	expr := &ast.BinaryExpr{
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"/api"`},
		Op: token.ADD,
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"/users"`},
	}
	collectStringParts(expr, &parts)
	if len(parts) != 2 {
		t.Fatalf("expected 2, got %d", len(parts))
	}
}
