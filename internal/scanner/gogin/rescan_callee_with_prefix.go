//ff:func feature=scan type=extract control=sequence
//ff:what 호출 대상 함수를 resolve하고, prefix를 전파하며 body를 재스캔하여 기존 endpoints를 업데이트한다
package gogin

import (
	"go/ast"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func rescanCalleeWithPrefix(call *ast.CallExpr, argIdx int, prefix string, parent *routerInfo, ctx *groupArgCtx) {
	targetPos := resolveCallTarget(call, ctx.info)
	if !targetPos.IsValid() {
		return
	}
	fnDecl, fnInfo := lookupFunc(targetPos, ctx.idx)
	if fnDecl == nil || fnDecl.Body == nil {
		return
	}

	paramName := ginRouterParamAtIndex(fnDecl, fnInfo, argIdx)
	if paramName == "" {
		return
	}

	targetFile := resolveTargetFilePath(targetPos, ctx)
	targetGinAlias := resolveTargetGinAlias(targetPos, ctx)
	if targetGinAlias == "" {
		targetGinAlias = ctx.ginAlias
	}

	targetRouters := map[string]*routerInfo{
		paramName: {prefix: prefix, middleware: append([]string{}, parent.middleware...)},
	}
	var eps []scanner.Endpoint
	localMap := map[int][]ast.Expr{}
	walkStmts(fnDecl.Body.List, targetGinAlias, targetFile, ctx.fset, targetRouters, &eps, localMap)

	applyRescanResults(eps, ctx)
}
