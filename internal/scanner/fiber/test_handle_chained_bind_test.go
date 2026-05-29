//ff:func feature=scan type=test control=sequence
//ff:what TestHandleChainedBind_BindBody 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleChainedBind_BindBody(t *testing.T) {
	ep := scanner.Endpoint{}

	// c.Bind()
	bindCall := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "c"},
			Sel: &ast.Ident{Name: "Bind"},
		},
	}

	// c.Bind().Body(&book)
	outerCall := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   bindCall,
			Sel: &ast.Ident{Name: "Body"},
		},
		Args: []ast.Expr{
			&ast.UnaryExpr{
				Op: token.AND,
				X:  &ast.Ident{Name: "book"},
			},
		},
	}

	handleChainedBind(&ep, outerCall, "Bind", nil)

	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatalf("expected request body to be set")
	}
	if ep.Request.Body.VarName != "book" {
		t.Fatalf("expected varName 'book', got %s", ep.Request.Body.VarName)
	}
	if ep.Request.Body.Method != "Bind" {
		t.Fatalf("expected method 'Bind', got %s", ep.Request.Body.Method)
	}

	// 두 번째 바인딩은 무시되어야 한다 (첫 번째 바인딩만 기록)
	second := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   bindCall,
			Sel: &ast.Ident{Name: "Body"},
		},
		Args: []ast.Expr{
			&ast.Ident{Name: "other"},
		},
	}
	handleChainedBind(&ep, second, "Bind", nil)
	if ep.Request.Body.VarName != "book" {
		t.Fatalf("expected first binding to be kept, got %s", ep.Request.Body.VarName)
	}
}
