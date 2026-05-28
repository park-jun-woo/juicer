//ff:func feature=scan type=extract control=selection topic=hono
//ff:what 단일 call_expression에서 c.json()/c.text()/c.body() 응답을 추출한다
package hono

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func extractOneResponse(call *sitter.Node, src []byte) *scanner.Response {
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return nil
	}
	obj := findChildByType(mem, "identifier")
	if obj == nil {
		return nil
	}
	if nodeText(obj, src) != "c" {
		return nil
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return nil
	}
	methodName := nodeText(prop, src)
	switch methodName {
	case "json":
		return parseJsonResponse(call, src)
	case "text":
		return &scanner.Response{Status: "200", Kind: "text"}
	case "body":
		return parseBodyResponse(call, src)
	}
	return nil
}
