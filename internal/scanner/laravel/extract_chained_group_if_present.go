//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what scoped 호출 없이 체인된 member_call이 있으면 체인 그룹으로 처리한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func extractChainedGroupIfPresent(mc *sitter.Node, fi fileInfo, outerPrefix string, outerMiddleware []string) []routeInfo {
	innerMC := findChildByType(mc, "member_call_expression")
	if innerMC == nil {
		return nil
	}
	return extractChainedGroup(mc, innerMC, fi, outerPrefix, outerMiddleware)
}
