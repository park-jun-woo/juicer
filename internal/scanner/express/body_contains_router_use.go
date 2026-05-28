//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 함수 본문에 router.use() 호출이 포함되어 있는지 확인한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func bodyContainsRouterUse(fn *sitter.Node, src []byte, routers map[string]bool) bool {
	for _, inner := range findAllByType(fn, "call_expression") {
		if isRouterUseCall(inner, src, routers) {
			return true
		}
	}
	return false
}
