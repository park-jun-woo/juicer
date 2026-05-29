//ff:func feature=scan type=extract control=sequence
//ff:what TypeSpec이 struct이면 타입명→StructType 인덱스에 등록한다
package echo

import (
	"go/ast"
)

func registerStructSpec(spec ast.Spec, idx *funcIndex) {
	ts, ok := spec.(*ast.TypeSpec)
	if !ok {
		return
	}
	st, ok := ts.Type.(*ast.StructType)
	if !ok {
		return
	}
	if _, exists := idx.astStructs[ts.Name.Name]; !exists {
		idx.astStructs[ts.Name.Name] = st
	}
}
