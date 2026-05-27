//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 데코레이터 노드에서 call_expression → arguments → object → enum 배열을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractEnumFromDecoratorNode extracts the enum array from a decorator node's call expression.
func extractEnumFromDecoratorNode(dn *sitter.Node, src []byte) []string {
	call := findChildByType(dn, "call_expression")
	if call == nil {
		return nil
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	obj := findChildByType(args, "object")
	if obj == nil {
		return nil
	}
	return extractEnumArray(obj, src)
}
