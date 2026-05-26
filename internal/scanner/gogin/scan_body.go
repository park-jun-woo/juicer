//ff:func feature=scan type=extract control=sequence
//ff:what 함수 body에서 gin.Context 메서드 호출을 탐색하여 요청/응답 정보를 수집한다
package gogin

import (
	"go/ast"
	"go/types"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func scanBody(ep *scanner.Endpoint, body *ast.BlockStmt, ctxName string, info *types.Info, idx *funcIndex, source string) {
	if body == nil {
		return
	}
	ast.Inspect(body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			// gin.Context가 아닌 함수 호출이면 1-depth 추적 확인
			if source == "handler" {
				checkOneDepthCall(ep, call, ctxName, info, idx)
			}
			return true
		}

		// ctxName.Method() 패턴 확인
		recvIdent, ok := sel.X.(*ast.Ident)
		if !ok || recvIdent.Name != ctxName {
			// gin.Context가 아닌 메서드 호출이면 1-depth 추적 확인
			if source == "handler" {
				checkOneDepthCall(ep, call, ctxName, info, idx)
			}
			return true
		}

		methodName := sel.Sel.Name

		// 요청 감지
		if bindMethods[methodName] {
			handleBind(ep, call, methodName, info)
		} else if queryMethods[methodName] {
			handleQuery(ep, call, methodName)
		} else if paramMethods[methodName] {
			// path param은 Phase 001에서 경로 패턴으로도 추출 — 중복 방지
			handlePathParam(ep, call)
		} else if formMethods[methodName] {
			handleForm(ep, call)
		} else if fileMethods[methodName] {
			handleFile(ep, call)
		} else if rawBodyMethods[methodName] {
			scanner.EnsureRequest(ep)
			ep.Request.RawBody = true
		}

		// 응답 감지
		if kind, ok := responseMethods[methodName]; ok {
			handleResponse(ep, call, kind, info, source)
		}

		return true
	})
}

