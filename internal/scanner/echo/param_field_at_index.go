//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 파라미터 목록에서 argIdx번째 위치의 필드와 이름을 반환한다
package echo

import (
	"go/ast"
)

func paramFieldAtIndex(params *ast.FieldList, argIdx int) (*ast.Field, string) {
	idx := 0
	for _, field := range params.List {
		names := field.Names
		if len(names) == 0 {
			names = []*ast.Ident{{Name: "_"}}
		}
		for _, n := range names {
			if idx == argIdx {
				return field, n.Name
			}
			idx++
		}
	}
	return nil, ""
}
