//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what forEach 본문의 router.use() 호출에서 부모 라우터 변수명을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func forEachParentRouter(call *sitter.Node, src []byte, routers map[string]bool) string {
	for _, inner := range findAllByType(call, "call_expression") {
		if isRouterUseCall(inner, src, routers) {
			return routerVarOfCall(inner, src)
		}
	}
	return ""
}
