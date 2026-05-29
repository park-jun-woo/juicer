//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what call_expression 자식에서 web::<method> 빌더를 HTTP 메서드로 변환한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func webMethodFromCall(child *sitter.Node, src []byte) string {
	if child.Type() != "call_expression" {
		return ""
	}
	scopedID := findChildByType(child, "scoped_identifier")
	if scopedID == nil {
		return ""
	}
	parts := splitScoped(nodeText(scopedID, src))
	if len(parts) != 2 || parts[0] != "web" {
		return ""
	}
	if m, ok := webMethodBuilders[parts[1]]; ok {
		return m
	}
	return ""
}
