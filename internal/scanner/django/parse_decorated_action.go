//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what decorated_definition에서 @action 메서드를 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parseDecoratedAction extracts an actionInfo from a decorated_definition if it has @action.
func parseDecoratedAction(decDef *sitter.Node, src []byte) *actionInfo {
	funcDef := findChildByType(decDef, "function_definition")
	if funcDef == nil {
		return nil
	}
	nameNode := findChildByType(funcDef, "identifier")
	if nameNode == nil {
		return nil
	}

	for _, dec := range childrenOfType(decDef, "decorator") {
		ai := parseActionDecorator(dec, src)
		if ai != nil {
			ai.name = nodeText(nameNode, src)
			ai.line = int(nameNode.StartPoint().Row) + 1
			return ai
		}
	}
	return nil
}
