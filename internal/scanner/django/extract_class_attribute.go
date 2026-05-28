//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 클래스 body에서 단순 속성 값을 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// extractClassAttribute extracts a simple identifier attribute from a class body.
// e.g., serializer_class = UserSerializer -> "UserSerializer"
func extractClassAttribute(body *sitter.Node, attrName string, src []byte) string {
	for _, node := range childrenOfType(body, "expression_statement") {
		val := extractAssignedValue(node, attrName, src)
		if val != "" {
			return val
		}
	}
	return ""
}
