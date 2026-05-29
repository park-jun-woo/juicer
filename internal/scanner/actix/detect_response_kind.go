//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what HttpResponse 호출 체인에 .json()이 연결되어 있으면 "json"을 반환한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func detectResponseKind(scopedID *sitter.Node, src []byte) string {
	parent := scopedID.Parent()
	if parent == nil {
		return ""
	}
	grandParent := parent.Parent()
	if isJSONFieldExpr(grandParent, src) {
		return "json"
	}
	if grandParent != nil && isJSONFieldExpr(grandParent.Parent(), src) {
		return "json"
	}
	return ""
}
