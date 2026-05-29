//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what assignment 노드의 RHS가 JSON source면 LHS 변수명을 반환한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// assignmentJSONVar returns the target variable name when an assignment's
// right-hand side is a JSON source (request.json / request.get_json()).
// Returns "" otherwise.
func assignmentJSONVar(asgn *sitter.Node, src []byte) string {
	lhs := asgn.Child(0)
	if lhs == nil || lhs.Type() != "identifier" {
		return ""
	}
	rhs := asgn.NamedChild(int(asgn.NamedChildCount()) - 1)
	if !isJSONSource(rhs, src) {
		return ""
	}
	return nodeText(lhs, src)
}
