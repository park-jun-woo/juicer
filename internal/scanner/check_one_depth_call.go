//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 핸들러 body의 함수 호출에서 gin.Context를 인자로 넘기면 해당 함수 body를 1-depth 추적하고, caller 인자에서 상태코드와 응답 타입을 전파한다
package scanner

import (
	"go/ast"
	"go/types"
)

func checkOneDepthCall(ep *Endpoint, call *ast.CallExpr, ctxName string, info *types.Info, idx *funcIndex) {
	if info == nil {
		return
	}

	// 호출 인자 중 ctxName이 있는지 확인
	hasCtx := false
	for _, arg := range call.Args {
		if id, ok := arg.(*ast.Ident); ok && id.Name == ctxName {
			hasCtx = true
			break
		}
	}
	if !hasCtx {
		return
	}

	// 호출 대상 함수 해석
	targetPos := resolveCallTarget(call, info)
	if !targetPos.IsValid() {
		return
	}

	fnDecl, fnInfo := lookupFunc(targetPos, idx)
	if fnDecl == nil || fnDecl.Body == nil {
		return
	}

	targetCtxName := ginCtxParamName(fnDecl.Type)
	if targetCtxName == "" {
		return
	}

	// 1단계: callee body 스캔 (기존 동작)
	prevLen := len(ep.Responses)
	callSource := exprName(call.Fun)
	scanBody(ep, fnDecl.Body, targetCtxName, fnInfo, idx, callSource)

	// 새로 추가된 응답이 없으면 병합 불필요
	if len(ep.Responses) == prevLen {
		return
	}

	// 2단계: caller 인자에서 상태코드와 응답 데이터 타입 해석
	callerStatus, callerTypeName, callerFields, callerConfidence := resolveCallerArgs(fnDecl, call, info, fnInfo)

	// 3단계: 새로 추가된 응답에 caller 인자 정보 병합
	for i := prevLen; i < len(ep.Responses); i++ {
		r := &ep.Responses[i]

		if r.Status == "(unknown)" && callerStatus != "" {
			r.Status = callerStatus
		}

		if callerTypeName != "" || len(callerFields) > 0 {
			r.TypeName = callerTypeName
			r.Fields = callerFields
			r.Confidence = callerConfidence
		}
	}
}
