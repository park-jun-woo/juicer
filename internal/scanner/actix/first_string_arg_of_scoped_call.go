//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 지정한 scoped 식별자 호출의 첫 문자열 인자를 찾는다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func firstStringArgOfScopedCall(node *sitter.Node, src []byte, scopedName string) string {
	var result string
	walkNodes(node, func(n *sitter.Node) {
		captureScopedCallArg(n, src, scopedName, &result)
	})
	return result
}
