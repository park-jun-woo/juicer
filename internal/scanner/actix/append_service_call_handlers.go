//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 노드가 .service() 호출이면 그 인자 식별자들을 핸들러 목록에 추가한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func appendServiceCallHandlers(parent *sitter.Node, src []byte, handlers []string) []string {
	if parent.Type() != "call_expression" {
		return handlers
	}
	fe := findChildByType(parent, "field_expression")
	if fe == nil {
		return handlers
	}
	fid := findChildByType(fe, "field_identifier")
	if fid == nil || nodeText(fid, src) != "service" {
		return handlers
	}
	args := findChildByType(parent, "arguments")
	if args == nil {
		return handlers
	}
	return appendIdentifierArgs(args, src, handlers)
}
