//ff:func feature=scan type=test control=sequence
//ff:what TestCollectStringParts_BasicLit 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestCollectStringParts_BasicLit(t *testing.T) {
	var parts []string
	lit := &ast.BasicLit{Kind: token.STRING, Value: `"/api"`}
	collectStringParts(lit, &parts)
	if len(parts) != 1 || parts[0] != "/api" {
		t.Fatalf("got %v", parts)
	}

	// binary expr (concatenation)
	var parts2 []string
	bin := &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"/a"`},
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"/b"`},
	}
	collectStringParts(bin, &parts2)
	if len(parts2) != 2 {
		t.Fatalf("expected 2 parts, got %v", parts2)
	}
}
