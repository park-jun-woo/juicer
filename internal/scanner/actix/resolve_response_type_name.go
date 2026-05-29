//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what .json 인자 노드에서 응답 본문 struct 타입명을 해석한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func resolveResponseTypeName(argNode *sitter.Node, ctx *responseCtx) string {
	if argNode.Type() == "struct_expression" {
		return structExprTypeName(argNode, ctx.src)
	}
	if argNode.Type() == "identifier" {
		return resolveLetBindingType(nodeText(argNode, ctx.src), ctx.block, ctx.src)
	}
	return ""
}
