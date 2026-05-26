//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 부모 노드의 모든 데코레이터 자식을 수집한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// collectDecoratorChildren collects all decorator children of a parent node.
func collectDecoratorChildren(parent *sitter.Node, src []byte) []decoratorInfo {
	decNodes := childrenOfType(parent, "decorator")
	var result []decoratorInfo
	for _, dn := range decNodes {
		d := parseDecorator(dn, src)
		if d.name != "" {
			result = append(result, d)
		}
	}
	return result
}
