//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 파라미터 노드에서 기본값을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractParamDefault(param *sitter.Node, src []byte) string {
	for i := 0; i < int(param.ChildCount()); i++ {
		child := param.Child(i)
		if child.Type() != "=" {
			continue
		}
		if i+1 < int(param.ChildCount()) {
			return nodeText(param.Child(i+1), src)
		}
	}
	return ""
}
