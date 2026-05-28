//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 파라미터 노드에서 변수명을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractParamName(param *sitter.Node, src []byte) string {
	for i := 0; i < int(param.ChildCount()); i++ {
		child := param.Child(i)
		if child.Type() == "identifier" {
			return nodeText(child, src)
		}
	}
	return ""
}
