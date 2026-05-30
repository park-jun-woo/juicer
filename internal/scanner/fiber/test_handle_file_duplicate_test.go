//ff:func feature=scan type=test control=sequence
//ff:what TestHandleFile_Duplicate 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleFile_Duplicate(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"avatar"`}}}
	handleFile(&ep, call)
	handleFile(&ep, call)
	if len(ep.Request.Files) != 1 {
		t.Fatalf("expected 1 file (dedup), got %d", len(ep.Request.Files))
	}
}
