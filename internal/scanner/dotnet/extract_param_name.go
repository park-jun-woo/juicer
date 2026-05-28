//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 파라미터 노드에서 이름을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractParamName(param *sitter.Node, src []byte) string {
	ids := childrenOfType(param, "identifier")
	if len(ids) == 0 {
		return ""
	}
	return nodeText(ids[len(ids)-1], src)
}
