package scanner

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_ValidUse(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Use"}},
		Args: []ast.Expr{&ast.Ident{Name: "authMiddleware"}},
	}
	routers := map[string]*routerInfo{"r": {}}
	tryUseCall(call, routers)
	if len(routers["r"].middleware) != 1 {
		t.Fatal("expected 1 middleware")
	}
}

func TestTryUseCall_NonSelectorFun(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}
	tryUseCall(call, nil)
}

func TestTryUseCall_NonUseMethod(t *testing.T) {
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
	}
	tryUseCall(call, nil)
}

func TestTryUseCall_MissingRouter(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Use"}},
		Args: []ast.Expr{&ast.Ident{Name: "mw"}},
	}
	tryUseCall(call, map[string]*routerInfo{})
}
