//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 데코레이터 인자에서 methods=['GET', 'POST'] 목록을 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// extractMethodsArg extracts HTTP methods from the methods= keyword argument.
// e.g., methods=["GET", "POST"] -> ["GET", "POST"]
// Returns nil if no methods argument is found.
func extractMethodsArg(args *sitter.Node, src []byte) []string {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "keyword_argument" {
			continue
		}
		keyNode := findChildByType(child, "identifier")
		if keyNode == nil || nodeText(keyNode, src) != "methods" {
			continue
		}
		// Find the list node: methods=["GET", "POST"]
		listNode := findChildByType(child, "list")
		if listNode == nil {
			continue
		}
		return extractStringList(listNode, src)
	}
	return nil
}
