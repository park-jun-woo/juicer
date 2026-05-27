//ff:func feature=scan type=test control=sequence
//ff:what TestHandleFile_DuplicateCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleFile_DuplicateCov(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{Files: []scanner.Param{{Name: "avatar", Type: "file"}}}}
	call := &ast.CallExpr{
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"avatar"`}},
	}
	handleFile(ep, call)
	if len(ep.Request.Files) != 1 {
		t.Fatal("expected 1 file (deduped)")
	}
}
