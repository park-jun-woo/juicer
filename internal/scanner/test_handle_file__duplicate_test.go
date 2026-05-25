//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleFile_Duplicate 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleFile_Duplicate(t *testing.T) {
	ep := &Endpoint{Request: &Request{Files: []Param{{Name: "avatar", Type: "file"}}}}
	call := &ast.CallExpr{
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"avatar"`}},
	}
	handleFile(ep, call)
	if len(ep.Request.Files) != 1 {
		t.Fatal("should not duplicate")
	}
}
