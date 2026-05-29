//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what web::scope 호출 노드 하나에서 prefix/핸들러를 해석해 스코프 목록에 추가한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func captureScope(n *sitter.Node, src []byte, scopes *[]scopeInfo) {
	if n.Type() != "call_expression" {
		return
	}
	scopedID := findChildByType(n, "scoped_identifier")
	if scopedID == nil || nodeText(scopedID, src) != "web::scope" {
		return
	}
	prefix := extractFirstStringArg(n, src)
	if prefix == "" {
		return
	}
	*scopes = append(*scopes, scopeInfo{
		prefix:   prefix,
		handlers: collectServiceHandlers(n, src),
	})
}
