//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what argument_list에서 정수 리터럴을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractIntLiteralFromArgList(argList *sitter.Node, src []byte) string {
	for i := 0; i < int(argList.ChildCount()); i++ {
		child := argList.Child(i)
		if child.Type() == "decimal_integer_literal" {
			return nodeText(child, src)
		}
	}
	return ""
}
