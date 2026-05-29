//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what scoped 식별자 노드가 대상이면 부모 호출의 첫 문자열 인자를 결과에 기록한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func captureScopedCallArg(n *sitter.Node, src []byte, scopedName string, result *string) {
	if *result != "" {
		return
	}
	if n.Type() != "scoped_identifier" || nodeText(n, src) != scopedName {
		return
	}
	parent := n.Parent()
	if parent != nil && parent.Type() == "call_expression" {
		*result = extractFirstStringArg(parent, src)
	}
}
