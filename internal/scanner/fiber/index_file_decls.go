//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 파일의 함수/struct 선언을 pos 및 이름 기반 인덱스에 등록한다
package fiber

import (
	"go/ast"
	"go/types"
)

func indexFileDecls(file *ast.File, info *types.Info, idx *funcIndex) {
	for _, decl := range file.Decls {
		indexFuncDecl(decl, info, idx)
		indexStructDecl(decl, idx)
	}
}
