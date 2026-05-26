//ff:func feature=scan type=extract control=iteration dimension=3
//ff:what 로드된 패키지에서 모든 함수 선언을 위치 기반으로 인덱싱한다
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/packages"
)

func buildFuncIndex(pkgs []*packages.Package) *funcIndex {
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	for _, pkg := range pkgs {
		if pkg.TypesInfo == nil {
			continue
		}
		for _, file := range pkg.Syntax {
			for _, decl := range file.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if !ok || fn.Body == nil {
					continue
				}
				idx.byPos[fn.Name.Pos()] = fn
				idx.info[fn.Name.Pos()] = pkg.TypesInfo
			}
		}
	}
	return idx
}

