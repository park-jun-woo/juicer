//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what map[string]any{...} 리터럴에서 키 목록과 값 타입을 추출한다
package echo

import (
	"go/ast"
	"go/types"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractMapFields(comp *ast.CompositeLit, info *types.Info) []scanner.Field {
	var fields []scanner.Field

	for _, elt := range comp.Elts {
		if field := buildMapField(elt, info); field != nil {
			fields = append(fields, *field)
		}
	}

	return fields
}
