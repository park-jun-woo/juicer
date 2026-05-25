//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what gin.H{...} 리터럴에서 키 목록과 값 타입을 추출한다
package scanner

import (
	"go/ast"
	"go/types"
)

func extractGinHFields(comp *ast.CompositeLit, info *types.Info) []Field {
	var fields []Field

	for _, elt := range comp.Elts {
		if field := buildGinHField(elt, info); field != nil {
			fields = append(fields, *field)
		}
	}

	return fields
}
