//ff:func feature=scan type=test control=sequence
//ff:what TestExtractBinaryPath_NoStringPartsCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractBinaryPath_NoStringPartsCov(t *testing.T) {
	e := &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.Ident{Name: "prefix"},
		Y:  &ast.Ident{Name: "suffix"},
	}
	_, ok := extractBinaryPath(e)
	if ok {
		t.Fatal("expected not ok for non-string parts")
	}
}
