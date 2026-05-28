//ff:func feature=scan type=test control=sequence
//ff:what TestTryUseCall_Pre 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_Pre(t *testing.T) {
	routers := map[string]*routerInfo{
		"e": {},
	}
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "e"},
			Sel: &ast.Ident{Name: "Pre"},
		},
		Args: []ast.Expr{&ast.Ident{Name: "trailingSlash"}},
	}
	tryUseCall(call, routers)
	if len(routers["e"].middleware) != 1 || routers["e"].middleware[0] != "trailingSlash" {
		t.Fatalf("expected [trailingSlash], got %v", routers["e"].middleware)
	}
}
