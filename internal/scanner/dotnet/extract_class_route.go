//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 클래스의 [Route] 어트리뷰트에서 prefix 경로를 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractClassRoute(cls *sitter.Node, src []byte, className string) string {
	attr := findAttribute(cls, src, AttrRoute)
	if attr == nil {
		return ""
	}
	route := attributeFirstStringArg(attr, src)
	if route == "" {
		return ""
	}
	return expandRouteTokens(route, className, "")
}
