//ff:func feature=scan type=test control=sequence
//ff:what TestHandleFile_Basic 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleFile_Basic(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"avatar"`},
		},
	}

	handleFile(&ep, call)

	if ep.Request == nil {
		t.Fatal("expected non-nil Request")
	}
	if len(ep.Request.Files) != 1 || ep.Request.Files[0].Name != "avatar" {
		t.Fatalf("expected file 'avatar', got %v", ep.Request.Files)
	}
}

func TestHandleFile_NoArgs(t *testing.T) {
	ep := scanner.Endpoint{}
	handleFile(&ep, &ast.CallExpr{})
	if ep.Request != nil {
		t.Fatalf("expected no request for no args, got %v", ep.Request)
	}
}

func TestHandleFile_EmptyName(t *testing.T) {
	ep := scanner.Endpoint{}
	// non-string arg -> stringLitValue "" -> return
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "nameVar"}}}
	handleFile(&ep, call)
	if ep.Request != nil {
		t.Fatalf("expected no request for empty name, got %v", ep.Request)
	}
}

func TestHandleFile_Duplicate(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"avatar"`}}}
	handleFile(&ep, call)
	handleFile(&ep, call) // duplicate -> not appended again
	if len(ep.Request.Files) != 1 {
		t.Fatalf("expected 1 file (dedup), got %d", len(ep.Request.Files))
	}
}
