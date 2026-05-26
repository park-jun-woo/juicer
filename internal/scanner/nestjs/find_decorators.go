//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 노드에 붙은 데코레이터 목록을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// findDecorators returns decorators attached to a node.
// For exported classes, decorators are children of the parent export_statement.
// For methods, decorators are consecutive siblings preceding the method in class_body.
func findDecorators(node *sitter.Node, src []byte) []decoratorInfo {
	parent := node.Parent()
	if parent == nil {
		return nil
	}
	if parent.Type() == "export_statement" {
		return collectDecoratorChildren(parent, src)
	}
	return collectPrecedingSiblingDecorators(parent, node, src)
}
