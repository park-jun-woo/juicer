//ff:func feature=scan type=test control=sequence
//ff:what extractBinaryPath — 문자열 연결 경로 추출 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"testing"
)

func binExpr(t *testing.T, expr string) *ast.BinaryExpr {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	be, ok := e.(*ast.BinaryExpr)
	if !ok {
		t.Fatalf("%q is not a binary expr", expr)
	}
	return be
}

func TestExtractBinaryPath_Concat(t *testing.T) {
	got, ok := extractBinaryPath(binExpr(t, `"/api" + "/users"`))
	if !ok || got != "/api/users" {
		t.Fatalf("got %q ok=%v", got, ok)
	}
}

func TestExtractBinaryPath_NotAdd(t *testing.T) {
	_, ok := extractBinaryPath(binExpr(t, "a - b"))
	if ok {
		t.Fatal("non-ADD should return false")
	}
}

func TestExtractBinaryPath_NoStringParts(t *testing.T) {
	// ADD of two idents -> no string parts -> false
	_, ok := extractBinaryPath(binExpr(t, "a + b"))
	if ok {
		t.Fatal("ident concat should return false")
	}
}
