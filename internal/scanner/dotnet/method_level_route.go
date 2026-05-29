//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 메서드 노드의 [Route] 어트리뷰트에서 액션 레벨 경로 세그먼트를 보충한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func methodLevelRoute(m *sitter.Node, src []byte) string {
	attr := findAttribute(m, src, AttrRoute)
	if attr == nil {
		return ""
	}
	return attributeFirstStringArg(attr, src)
}
