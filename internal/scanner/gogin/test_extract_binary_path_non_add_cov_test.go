//ff:func feature=scan type=test control=sequence
//ff:what TestExtractBinaryPath_NonAddCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractBinaryPath_NonAddCov(t *testing.T) {
	e := &ast.BinaryExpr{
		Op: token.MUL,
		X:  &ast.BasicLit{Kind: token.INT, Value: "1"},
		Y:  &ast.BasicLit{Kind: token.INT, Value: "2"},
	}
	_, ok := extractBinaryPath(e)
	if ok {
		t.Fatal("expected not ok for non-ADD op")
	}
}
