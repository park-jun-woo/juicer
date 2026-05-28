//ff:func feature=scan type=extract control=selection
//ff:what 핸들러 AST 표현에서 함수 body를 찾아 요청/응답을 추출한다
package echo

import (
	"go/ast"
	"go/types"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func analyzeExpr(ep *scanner.Endpoint, expr ast.Expr, info *types.Info, idx *funcIndex) {
	switch e := expr.(type) {
	case *ast.FuncLit:
		// 인라인 함수 — body 직접 분석
		ctxName := echoCtxParamNameInfo(e.Type, info)
		if ctxName == "" {
			return
		}
		scanBody(ep, e.Body, ctxName, info, idx, "handler")

	case *ast.SelectorExpr:
		// h.Method 형태 — 타입 정보에서 메서드 선언 해석
		sel, ok := info.Selections[e]
		if !ok {
			return
		}
		fn := sel.Obj()
		if fn == nil {
			return
		}
		fnDecl, fnInfo := lookupFunc(fn.Pos(), idx)
		if fnDecl == nil {
			return
		}
		ctxName := echoCtxParamNameInfo(fnDecl.Type, fnInfo)
		if ctxName == "" {
			return
		}
		scanBody(ep, fnDecl.Body, ctxName, fnInfo, idx, "handler")

	case *ast.Ident:
		// 패키지 레벨 함수 참조
		obj := info.Uses[e]
		if obj == nil {
			return
		}
		fnDecl, fnInfo := lookupFunc(obj.Pos(), idx)
		if fnDecl == nil {
			return
		}
		ctxName := echoCtxParamNameInfo(fnDecl.Type, fnInfo)
		if ctxName == "" {
			return
		}
		scanBody(ep, fnDecl.Body, ctxName, fnInfo, idx, "handler")

	case *ast.CallExpr:
		// handler() 같은 호출 결과가 핸들러인 경우 — 내부 함수 추적 시도
		analyzeExpr(ep, e.Fun, info, idx)
	}
}
