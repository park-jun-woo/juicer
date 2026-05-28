//ff:func feature=scan type=extract control=sequence topic=django
//ff:what @action(detail=True, methods=["post"]) 데코레이터를 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parseActionDecorator parses @action(detail=True, methods=["post"]) decorator.
func parseActionDecorator(dec *sitter.Node, src []byte) *actionInfo {
	callNode := findChildByType(dec, "call")
	if callNode == nil {
		return nil
	}
	funcNode := findChildByType(callNode, "identifier")
	if funcNode == nil || nodeText(funcNode, src) != "action" {
		return nil
	}
	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return nil
	}
	ai := &actionInfo{methods: []string{"GET"}}
	applyActionKeywords(ai, args, src)
	return ai
}
