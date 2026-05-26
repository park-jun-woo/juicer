//ff:func feature=scan type=test control=sequence
//ff:what TestHandleFile_Basic 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandleFile_Basic(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"avatar"`}},
	}
	handleFile(ep, call)
	if len(ep.Request.Files) != 1 {
		t.Fatal("expected 1 file")
	}
}

