//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 호출이 HTTP 메서드 등록(GET/POST/…)이면 Endpoint와 handler 표현 목록을 반환한다
package echo

import (
	"go/ast"
	"go/token"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func tryRouteCall(call *ast.CallExpr, routers map[string]*routerInfo, filePath string, fset *token.FileSet) (scanner.Endpoint, []ast.Expr, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return scanner.Endpoint{}, nil, false
	}
	if !echoMethods[sel.Sel.Name] {
		return scanner.Endpoint{}, nil, false
	}
	recv := identName(sel.X)
	router, ok := routers[recv]
	if !ok {
		return scanner.Endpoint{}, nil, false
	}
	if len(call.Args) < 2 {
		return scanner.Endpoint{}, nil, false
	}
	pathStr, ok := extractPathString(call.Args[0])
	if !ok {
		return scanner.Endpoint{}, nil, false
	}

	path := scanner.JoinPath(router.prefix, pathStr)
	// 핸들러 인자: 마지막 인자가 핸들러, 경로와 핸들러 사이의 인자는 미들웨어
	handlerArg := call.Args[len(call.Args)-1]
	handler := exprName(handlerArg)
	pos := fset.Position(call.Pos())

	// 핸들러 AST 표현 보존 (analyzeHandlers에서 함수 body 접근용)
	var handlerExprs []ast.Expr
	for _, arg := range call.Args[1:] {
		handlerExprs = append(handlerExprs, arg)
	}

	ep := scanner.Endpoint{
		Method:  sel.Sel.Name,
		Path:    path,
		Handler: filePath + ":" + handler,
		File:    filePath,
		Line:    pos.Line,
	}
	if len(router.middleware) > 0 {
		ep.Middleware = append([]string{}, router.middleware...)
	}
	if params := scanner.PathParams(path); len(params) > 0 {
		ep.Request = &scanner.Request{PathParams: params}
	}
	return ep, handlerExprs, true
}
