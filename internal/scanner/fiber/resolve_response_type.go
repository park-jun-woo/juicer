//ff:func feature=scan type=extract control=sequence
//ff:what 응답 인자의 타입을 추적한다 (named struct)
package fiber

import (
	"go/ast"
	"go/types"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveResponseType(expr ast.Expr, info *types.Info) (typeName string, fields []scanner.Field, confidence string) {
	if info == nil {
		return "", nil, ""
	}

	typeName, fields = resolveExprType(expr, info)
	if len(fields) > 0 {
		return typeName, fields, "full"
	}
	return typeName, nil, ""
}
