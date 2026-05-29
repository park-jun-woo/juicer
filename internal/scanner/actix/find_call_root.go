//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 호출 노드 하위에서 web::scope/web::resource 루트 식별자를 찾는다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func findCallRoot(node *sitter.Node, src []byte) string {
	var result string
	walkNodes(node, func(n *sitter.Node) {
		captureCallRoot(n, src, &result)
	})
	return result
}
