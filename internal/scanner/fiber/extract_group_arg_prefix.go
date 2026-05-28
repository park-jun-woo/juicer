//ff:func feature=scan type=extract control=sequence
//ff:what 인자가 라우터 변수 또는 router.Group("prefix") 호출이면 결합된 prefix와 라우터 정보를 반환한다
package fiber

import (
	"go/ast"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractGroupArgPrefix(arg ast.Expr, ctx *groupArgCtx) (string, *routerInfo, bool) {
	// Variable argument: authGroup etc — look up directly in routers map
	if ident, ok := arg.(*ast.Ident); ok {
		if ri, ok := ctx.routers[ident.Name]; ok {
			return ri.prefix, ri, true
		}
	}

	groupCall, ok := arg.(*ast.CallExpr)
	if !ok {
		return "", nil, false
	}
	sel, ok := groupCall.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Group" {
		return "", nil, false
	}
	recv := identName(sel.X)
	parent, ok := ctx.routers[recv]
	if !ok {
		return "", nil, false
	}

	prefix := parent.prefix
	if len(groupCall.Args) > 0 {
		if s, ok := extractPathString(groupCall.Args[0]); ok {
			prefix = scanner.JoinPath(prefix, s)
		}
	}
	return prefix, parent, true
}
