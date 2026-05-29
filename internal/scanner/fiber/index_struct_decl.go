//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what GenDecl의 type spec에서 struct 타입을 이름 인덱스에 등록한다
package fiber

import (
	"go/ast"
)

func indexStructDecl(decl ast.Decl, idx *funcIndex) {
	gen, ok := decl.(*ast.GenDecl)
	if !ok {
		return
	}
	for _, spec := range gen.Specs {
		registerStructSpec(spec, idx)
	}
}
