//ff:func feature=scan type=extract control=sequence
//ff:what 인자가 router.Group("prefix") 호출이면 결합된 prefix와 부모 라우터 정보를 반환한다
package gogin

import (
	"go/ast"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func extractGroupArgPrefix(arg ast.Expr, ctx *groupArgCtx) (string, *routerInfo, bool) {
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
