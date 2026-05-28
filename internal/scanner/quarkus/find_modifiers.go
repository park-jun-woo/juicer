//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 노드에서 modifiers 자식을 찾는다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func findModifiers(node *sitter.Node) *sitter.Node {
	return findChildByType(node, "modifiers")
}
