//ff:func feature=scan type=extract control=selection
//ff:what 값 표현에서 타입을 추론한다 (리터럴은 정확, 그 외 best-effort)
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
)

func inferValueType(expr ast.Expr, info *types.Info) string {
	switch e := expr.(type) {
	case *ast.BasicLit:
		switch e.Kind {
		case token.STRING:
			return "string"
		case token.INT:
			return "integer"
		case token.FLOAT:
			return "number"
		}
	case *ast.Ident:
		if e.Name == "true" || e.Name == "false" {
			return "boolean"
		}
		if e.Name == "nil" {
			return "null"
		}
	case *ast.CompositeLit:
		if isGinH(e, info) {
			return "object"
		}
		// 슬라이스 리터럴
		if _, ok := e.Type.(*ast.ArrayType); ok {
			return "array"
		}
	case *ast.SliceExpr:
		return "array"
	}

	// go/types 폴백
	if info != nil {
		if tv, ok := info.Types[expr]; ok {
			return formatType(tv.Type)
		}
	}

	return "unknown"
}

