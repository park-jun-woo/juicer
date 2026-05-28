//ff:func feature=scan type=extract control=selection topic=express
//ff:what arrow_function 또는 function/function_expression 노드에서 본문을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractFunctionBody(node *sitter.Node) *sitter.Node {
	switch node.Type() {
	case "arrow_function":
		body := node.ChildByFieldName("body")
		if body != nil {
			return body
		}
	case "function", "function_expression":
		return findChildByType(node, "statement_block")
	}
	return nil
}
