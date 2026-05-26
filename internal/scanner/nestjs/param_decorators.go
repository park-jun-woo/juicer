//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 파라미터 노드에서 데코레이터를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// paramDecorators extracts decorators from a parameter node.
// In tree-sitter TS, parameter decorators are children of the parameter node.
func paramDecorators(param *sitter.Node, src []byte) []decoratorInfo {
	return collectDecoratorChildren(param, src)
}
