//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what register 첫 인자가 autoload 식별자인지 판정한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func isAutoloadArg(node *sitter.Node, src []byte, names map[string]bool) bool {
	if node.Type() != "identifier" {
		return false
	}
	return names[nodeText(node, src)]
}
