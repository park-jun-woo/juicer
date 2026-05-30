//ff:func feature=scan type=test control=iteration dimension=1
//ff:what indexDeclsAndFindCall 테스트 헬퍼: 파일의 FuncDecl을 색인하고 registerRoutes 호출 노드를 찾음
package fiber

import (
	"go/ast"
	"go/token"
	"go/types"
)

// indexDeclsAndFindCall builds a funcIndex over file's FuncDecls and returns
// the first call to registerRoutes found within them.
func indexDeclsAndFindCall(file *ast.File, info *types.Info) (*funcIndex, *ast.CallExpr) {
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, info: map[token.Pos]*types.Info{}}
	var call *ast.CallExpr
	for _, d := range file.Decls {
		call = indexOneDecl(idx, info, d, call)
	}
	return idx, call
}
