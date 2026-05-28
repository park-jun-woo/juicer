//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 함수 파라미터 목록에서 첫 번째 파라미터 이름을 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractFirstParamName(params *sitter.Node, src []byte) string {
	for i := 0; i < int(params.ChildCount()); i++ {
		name := paramNodeName(params.Child(i), src)
		if name != "" {
			return name
		}
	}
	return ""
}
