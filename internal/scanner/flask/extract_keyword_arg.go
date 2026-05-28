//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what argument_list에서 이름으로 키워드 인자 값을 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// extractKeywordArg finds a keyword argument by name and returns its string value.
func extractKeywordArg(args *sitter.Node, name string, src []byte) string {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "keyword_argument" {
			continue
		}
		keyNode := findChildByType(child, "identifier")
		if keyNode == nil || nodeText(keyNode, src) != name {
			continue
		}
		valNode := findChildByType(child, "string")
		if valNode != nil {
			return unquotePython(nodeText(valNode, src))
		}
	}
	return ""
}
