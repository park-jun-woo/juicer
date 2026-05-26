//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what AST 내 단순 할당문에서 변수의 문자열 리터럴 값을 해석한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// resolveVariableValue looks up a simple assignment like `X = "literal"` in the
// same AST and returns the string literal value. Returns "" if not found or if
// the value is not a simple string literal.
func resolveVariableValue(root *sitter.Node, varName string, src []byte) string {
	assignments := findAllByType(root, "assignment")
	for _, assign := range assignments {
		left := findChildByType(assign, "identifier")
		if left == nil || nodeText(left, src) != varName {
			continue
		}
		strNode := findChildByType(assign, "string")
		if strNode != nil {
			return unquotePython(nodeText(strNode, src))
		}
	}
	return ""
}
