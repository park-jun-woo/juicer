//ff:func feature=scan type=extract control=sequence
//ff:what 단일 caller 인자와 callee 파라미터 타입에서 상태코드와 응답 타입을 해석한다
package scanner

import (
	"go/ast"
	"go/types"
)

// resolveCallerArg resolves status code or response type from a single caller argument.
func resolveCallerArg(paramType types.Type, callerArg ast.Expr, callerInfo *types.Info) callerArgResult {
	if isGinContextTypeInfo(paramType) {
		return callerArgResult{skip: true}
	}

	underlying := paramType.Underlying()

	if basic, ok := underlying.(*types.Basic); ok && isIntKind(basic.Kind()) {
		s := resolveStatusCode(callerArg, callerInfo)
		if s != "(unknown)" {
			return callerArgResult{status: s}
		}
		return callerArgResult{}
	}

	if ifc, ok := underlying.(*types.Interface); ok && ifc.NumMethods() == 0 {
		tn, f, conf := resolveResponseType(callerArg, callerInfo)
		if tn != "" || len(f) > 0 {
			return callerArgResult{typeName: tn, fields: f, confidence: conf}
		}
	}

	return callerArgResult{}
}
