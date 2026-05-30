//ff:func feature=scan type=test control=sequence
//ff:what TestHandleFile_Basic 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
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
