//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 노드에서 class_constant_access_expression을 직계 우선, 없으면 하위에서 찾는다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func findClassConstantAccess(node *sitter.Node) *sitter.Node {
	if direct := findChildByType(node, "class_constant_access_expression"); direct != nil {
		return direct
	}
	nodes := findAllByType(node, "class_constant_access_expression")
	if len(nodes) > 0 {
		return nodes[0]
	}
	return nil
}
