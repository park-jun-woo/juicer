//ff:func feature=scan type=test control=sequence
//ff:what TestExtractBinaryPath_AddOp 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractBinaryPath_AddOp(t *testing.T) {
	e := &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"/api"`},
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"/v1"`},
	}
	path, ok := extractBinaryPath(e)
	if !ok {
		t.Fatal("expected ok")
	}
	if path != "/api/v1" {
		t.Fatalf("expected /api/v1, got %s", path)
	}

	// non-ADD op
	e2 := &ast.BinaryExpr{Op: token.MUL, X: &ast.Ident{Name: "a"}, Y: &ast.Ident{Name: "b"}}
	_, ok = extractBinaryPath(e2)
	if ok {
		t.Fatal("expected false for non-ADD")
	}

	// ADD but no string parts
	e3 := &ast.BinaryExpr{Op: token.ADD, X: &ast.Ident{Name: "a"}, Y: &ast.Ident{Name: "b"}}
	_, ok = extractBinaryPath(e3)
	if ok {
		t.Fatal("expected false for no string parts")
	}
}

