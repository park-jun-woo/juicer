//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what JSON Schema AST 노드를 scanner.Param 슬라이스로 변환한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func jsonSchemaToParams(schemaNode *sitter.Node, src []byte) []scanner.Param {
	if schemaNode == nil || schemaNode.Type() != "object" {
		return nil
	}
	propsNode := findPairValue(schemaNode, src, "properties")
	if propsNode == nil || propsNode.Type() != "object" {
		return nil
	}
	var params []scanner.Param
	for _, pair := range childrenOfType(propsNode, "pair") {
		p := convertPropertyToParam(pair, src)
		if p != nil {
			params = append(params, *p)
		}
	}
	return params
}
