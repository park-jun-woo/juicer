//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractBinaryPath_NonAdd 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractBinaryPath_NonAdd(t *testing.T) {
	e := &ast.BinaryExpr{Op: token.MUL, X: &ast.Ident{Name: "a"}, Y: &ast.Ident{Name: "b"}}
	_, ok := extractBinaryPath(e)
	if ok {
		t.Fatal("expected not ok")
	}
}
