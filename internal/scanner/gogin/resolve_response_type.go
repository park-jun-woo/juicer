//ff:func feature=scan type=extract control=sequence
//ff:what 응답 인자의 타입을 추적한다 (named struct 또는 gin.H 리터럴)
package gogin

import (
	"go/ast"
	"go/types"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveResponseType(expr ast.Expr, info *types.Info) (typeName string, fields []scanner.Field, confidence string) {
	if info == nil {
		return "", nil, ""
	}

	// gin.H{...} 리터럴 감지
	if comp, ok := expr.(*ast.CompositeLit); ok {
		if isGinH(comp, info) {
			fields := extractGinHFields(comp, info)
			return "gin.H", fields, "partial"
		}
	}

	typeName, fields = resolveExprType(expr, info)
	if len(fields) > 0 {
		return typeName, fields, "full"
	}
	return typeName, nil, ""
}

