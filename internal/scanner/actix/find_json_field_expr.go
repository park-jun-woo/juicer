//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what HttpResponse scoped_identifier 위에서 .json field_expression 노드를 찾는다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func findJSONFieldExpr(scopedID *sitter.Node, src []byte) *sitter.Node {
	parent := scopedID.Parent()
	if parent == nil {
		return nil
	}
	grandParent := parent.Parent()
	if isJSONFieldExpr(grandParent, src) {
		return grandParent
	}
	if grandParent != nil && isJSONFieldExpr(grandParent.Parent(), src) {
		return grandParent.Parent()
	}
	return nil
}
