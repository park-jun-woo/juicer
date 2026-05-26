//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractPathString_Unknown 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExtractPathString_Unknown(t *testing.T) {
	expr := &ast.Ident{Name: "x"}
	_, ok := extractPathString(expr)
	if ok {
		t.Fatal("expected not ok")
	}
}
