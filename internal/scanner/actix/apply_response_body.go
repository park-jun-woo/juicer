//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what .json(...) 인자에서 응답 본문 타입명·필드를 해석해 Response에 반영한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyResponseBody(resp *scanner.Response, scopedID *sitter.Node, ctx *responseCtx) {
	argNode := findJSONArg(scopedID, ctx.src)
	if argNode == nil {
		return
	}
	typeName := resolveResponseTypeName(argNode, ctx)
	if typeName == "" {
		return
	}
	resp.TypeName = typeName
	fields := resolveStructFields(typeName, ctx.sIdx, ctx.cache)
	if len(fields) > 0 {
		resp.Fields = fields
		resp.Confidence = "full"
	}
}
