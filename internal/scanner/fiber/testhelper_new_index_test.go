//ff:func feature=scan type=test topic=fiber control=sequence
//ff:what fiber 테스트 헬퍼 — 빈 funcIndex 생성
package fiber

import (
	"go/ast"
	"go/token"
)

func newFiberIndex() *funcIndex {
	return &funcIndex{
		byName:     map[string]*ast.FuncDecl{},
		byPos:      map[token.Pos]*ast.FuncDecl{},
		astStructs: map[string]*ast.StructType{},
	}
}
