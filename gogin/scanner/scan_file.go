//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what Gin import가 있는 파일에서 함수별로 라우트를 탐색한다
package scanner

import (
	"go/ast"
	"go/token"
)

func scanFile(file *ast.File, filePath string, fset *token.FileSet) []Endpoint {
	ginAlias := ginPkgName(file)
	if ginAlias == "" {
		return nil
	}

	var endpoints []Endpoint
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		routers := make(map[string]*routerInfo)
		registerParams(fn, ginAlias, routers)
		var eps []Endpoint
		walkStmts(fn.Body.List, ginAlias, filePath, fset, routers, &eps)
		endpoints = append(endpoints, eps...)
	}
	return endpoints
}

