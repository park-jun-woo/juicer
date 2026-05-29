//ff:func feature=scan type=extract control=sequence
//ff:what 함수 body에서 fiber.Ctx 메서드 호출을 탐색하여 요청/응답 정보를 수집한다
package fiber

import (
	"go/ast"
	"go/types"

	"github.com/park-jun-woo/codistill/internal/scanner"
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
			// fiber.Ctx가 아닌 함수 호출이면 1-depth 추적 확인
			if source == "handler" {
				checkOneDepthCall(ep, call, ctxName, info, idx)
			}
			return true
		}

		// ctxName.Method() 패턴 확인
		recvIdent, ok := sel.X.(*ast.Ident)
		if !ok || recvIdent.Name != ctxName {
			// c.Status(N).JSON(body) 패턴: 체이닝된 호출 감지
			if innerCall, ok := sel.X.(*ast.CallExpr); ok {
				if innerSel, ok := innerCall.Fun.(*ast.SelectorExpr); ok {
					if innerRecv, ok := innerSel.X.(*ast.Ident); ok && innerRecv.Name == ctxName {
						if innerSel.Sel.Name == "Status" {
							// c.Status(N).JSON(body) 패턴
							handleChainedResponse(ep, innerCall, call, sel.Sel.Name, info, source)
							return true
						}
						if innerSel.Sel.Name == "Bind" && sel.Sel.Name == "Body" {
							// c.Bind().Body(req) 체이닝 바인딩 패턴
							handleChainedBind(ep, call, "Bind", info)
							return true
						}
					}
				}
			}
			// fiber.Ctx가 아닌 메서드 호출이면 1-depth 추적 확인
			if source == "handler" {
				checkOneDepthCall(ep, call, ctxName, info, idx)
			}
			return true
		}

		methodName := sel.Sel.Name

		// 요청 감지
		if bindMethods[methodName] {
			handleBodyParser(ep, call, methodName, info)
		} else if queryMethods[methodName] {
			handleQuery(ep, call, methodName)
		} else if paramMethods[methodName] {
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
