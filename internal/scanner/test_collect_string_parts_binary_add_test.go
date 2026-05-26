//ff:func feature=scan type=test control=sequence
//ff:what TestCollectStringParts_BinaryAdd 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestCollectStringParts_BinaryAdd(t *testing.T) {
	var parts []string
	expr := &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"/api"`},
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"/v1"`},
	}
	collectStringParts(expr, &parts)
	if len(parts) != 2 {
		t.Fatalf("expected 2 parts, got %d", len(parts))
	}
}
