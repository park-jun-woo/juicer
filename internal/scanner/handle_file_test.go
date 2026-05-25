package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleFile_Basic(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"avatar"`}},
	}
	handleFile(ep, call)
	if len(ep.Request.Files) != 1 {
		t.Fatal("expected 1 file")
	}
}

func TestHandleFile_NoArgs(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{}
	handleFile(ep, call)
	if ep.Request != nil {
		t.Fatal("expected nil request")
	}
}

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
