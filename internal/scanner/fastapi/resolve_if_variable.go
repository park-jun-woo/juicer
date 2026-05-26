//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 접두사 값이 변수 참조이면 같은 파일 내 할당문으로 해석한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// resolveIfVariable checks if val is a variable reference (identifier text
// rather than a path literal) and resolves it against same-file assignments.
// Returns the original val if it is empty, already a path literal, or if
// resolution fails.
func resolveIfVariable(root *sitter.Node, val string, src []byte) string {
	if val == "" {
		return val
	}
	if val[0] == '/' {
		return val
	}
	if !isIdentifier(val) {
		return val
	}
	resolved := resolveVariableValue(root, val, src)
	if resolved != "" {
		return resolved
	}
	return val
}
