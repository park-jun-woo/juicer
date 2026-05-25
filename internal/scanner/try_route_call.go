//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 호출이 HTTP 메서드 등록(GET/POST/…)이면 Endpoint를 반환한다
package scanner

import (
	"go/ast"
	"go/token"
)

func tryRouteCall(call *ast.CallExpr, routers map[string]*routerInfo, filePath string, fset *token.FileSet) (Endpoint, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return Endpoint{}, false
	}
	if !ginMethods[sel.Sel.Name] {
		return Endpoint{}, false
	}
	recv := identName(sel.X)
	router, ok := routers[recv]
	if !ok {
		return Endpoint{}, false
	}
	if len(call.Args) < 2 {
		return Endpoint{}, false
	}
	pathStr, ok := extractPathString(call.Args[0])
	if !ok {
		return Endpoint{}, false
	}

	path := joinPath(router.prefix, pathStr)
	// 핸들러 인자: 마지막 인자가 핸들러, 경로와 핸들러 사이의 인자는 미들웨어
	handlerArg := call.Args[len(call.Args)-1]
	handler := exprName(handlerArg)
	pos := fset.Position(call.Pos())

	// 핸들러 AST 표현 보존 (handler.go에서 함수 body 접근용)
	var handlerExprs []ast.Expr
	for _, arg := range call.Args[1:] {
		handlerExprs = append(handlerExprs, arg)
	}

	ep := Endpoint{
		Method:       sel.Sel.Name,
		Path:         path,
		Handler:      filePath + ":" + handler,
		File:         filePath,
		Line:         pos.Line,
		handlerExprs: handlerExprs,
	}
	if len(router.middleware) > 0 {
		ep.Middleware = append([]string{}, router.middleware...)
	}
	if params := pathParams(path); len(params) > 0 {
		ep.Request = &Request{PathParams: params}
	}
	return ep, true
}

