//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 프로퍼티 노드에서 데코레이터를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// propertyDecorators extracts decorators from a property node.
func propertyDecorators(prop *sitter.Node, src []byte) []decoratorInfo {
	parent := prop.Parent()
	if parent == nil {
		return nil
	}
	return collectPrecedingSiblingDecorators(parent, prop, src)
}
