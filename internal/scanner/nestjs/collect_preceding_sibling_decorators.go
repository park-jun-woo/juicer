//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 지정 노드 앞에 연속된 데코레이터 형제 노드를 수집한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// collectPrecedingSiblingDecorators scans backwards from node in parent to find consecutive decorators.
func collectPrecedingSiblingDecorators(parent, node *sitter.Node, src []byte) []decoratorInfo {
	nodeIdx := findNodeIndex(parent, node)
	var result []decoratorInfo
	for i := nodeIdx - 1; i >= 0; i-- {
		child := parent.Child(i)
		if child.Type() != "decorator" {
			break
		}
		d := parseDecorator(child, src)
		if d.name != "" {
			result = append(result, d)
		}
	}
	return result
}
