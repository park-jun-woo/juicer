//ff:func feature=scan type=extract control=iteration dimension=3
//ff:what 로드된 패키지에서 모든 함수 선언을 위치 기반으로 인덱싱한다
package echo

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/packages"
)

func buildFuncIndex(pkgs []*packages.Package) *funcIndex {
	idx := &funcIndex{
		byPos:      make(map[token.Pos]*ast.FuncDecl),
		info:       make(map[token.Pos]*types.Info),
		byName:     make(map[string]*ast.FuncDecl),
		astStructs: make(map[string]*ast.StructType),
	}
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			indexFileDecls(file, pkg.TypesInfo, idx)
		}
	}
	return idx
}
