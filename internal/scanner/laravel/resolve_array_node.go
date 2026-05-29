//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 노드 자체 또는 그 하위에서 array_creation_expression 노드를 찾는다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func resolveArrayNode(node *sitter.Node) *sitter.Node {
	if arr := findChildByType(node, "array_creation_expression"); arr != nil {
		return arr
	}
	if node.Type() == "array_creation_expression" {
		return node
	}
	return nil
}
