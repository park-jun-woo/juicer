//ff:func feature=scan type=extract control=sequence
//ff:what FuncDecl이면 pos 인덱스(타입정보 있을 때)와 이름 인덱스에 등록한다
package echo

import (
	"go/ast"
	"go/types"
)

func indexFuncDecl(decl ast.Decl, info *types.Info, idx *funcIndex) {
	fn, ok := decl.(*ast.FuncDecl)
	if !ok || fn.Body == nil {
		return
	}
	if info != nil {
		idx.byPos[fn.Name.Pos()] = fn
		idx.info[fn.Name.Pos()] = info
	}
	if _, exists := idx.byName[fn.Name.Name]; !exists {
		idx.byName[fn.Name.Name] = fn
	}
}
