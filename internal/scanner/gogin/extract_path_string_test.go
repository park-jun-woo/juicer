//ff:func feature=scan type=test control=sequence
//ff:what TestExtractPathString_BasicLit 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractPathString_BasicLit(t *testing.T) {
	expr := &ast.BasicLit{Kind: token.STRING, Value: `"/api/v1"`}
	path, ok := extractPathString(expr)
	if !ok || path != "/api/v1" {
		t.Fatalf("expected /api/v1, got %s ok=%v", path, ok)
	}

	// binary expr
	bin := &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"/a"`},
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"/b"`},
	}
	path2, ok := extractPathString(bin)
	if !ok || path2 != "/a/b" {
		t.Fatalf("binary: got %s ok=%v", path2, ok)
	}

	// non-string expr
	_, ok = extractPathString(&ast.Ident{Name: "x"})
	if ok {
		t.Fatal("expected false for ident")
	}
}

