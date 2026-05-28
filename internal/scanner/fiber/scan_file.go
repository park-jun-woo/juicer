//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what Fiber import가 있는 파일에서 함수별로 라우트를 탐색한다
package fiber

import (
	"go/ast"
	"go/token"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func scanFile(file *ast.File, filePath string, fset *token.FileSet) ([]scanner.Endpoint, map[int][]ast.Expr) {
	fiberAlias := fiberPkgName(file)
	if fiberAlias == "" {
		return nil, nil
	}

	var endpoints []scanner.Endpoint
	hmap := map[int][]ast.Expr{}
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		routers := make(map[string]*routerInfo)
		registerParams(fn, fiberAlias, routers)
		var eps []scanner.Endpoint
		localMap := map[int][]ast.Expr{}
		walkStmts(fn.Body.List, fiberAlias, filePath, fset, routers, &eps, localMap)
		// localMap 인덱스를 전역 오프셋으로 변환
		offset := len(endpoints)
		for k, v := range localMap {
			hmap[offset+k] = v
		}
		endpoints = append(endpoints, eps...)
	}
	return endpoints, hmap
}
