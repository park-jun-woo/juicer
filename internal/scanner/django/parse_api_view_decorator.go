//ff:func feature=scan type=extract control=sequence topic=django
//ff:what @api_view(["GET", "POST"]) 데코레이터를 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parseAPIViewDecoratorNode parses @api_view(["GET", "POST"]) decorator.
func parseAPIViewDecoratorNode(dec *sitter.Node, src []byte) []string {
	callNode := findChildByType(dec, "call")
	if callNode == nil {
		return nil
	}
	funcNode := findChildByType(callNode, "identifier")
	if funcNode == nil || nodeText(funcNode, src) != "api_view" {
		return nil
	}
	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return []string{"GET"}
	}
	listNode := findChildByType(args, "list")
	if listNode == nil {
		return []string{"GET"}
	}
	return extractStringList(listNode, src)
}
