//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what JSON Schema AST 노드를 scanner.Field 슬라이스로 변환한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func jsonSchemaToFields(schemaNode *sitter.Node, src []byte) []scanner.Field {
	if schemaNode == nil || schemaNode.Type() == "identifier" || schemaNode.Type() != "object" {
		return nil
	}
	propsNode := findPairValue(schemaNode, src, "properties")
	if propsNode == nil || propsNode.Type() != "object" {
		return nil
	}
	requiredSet := extractRequiredSet(schemaNode, src)
	var fields []scanner.Field
	for _, pair := range childrenOfType(propsNode, "pair") {
		f := convertPropertyPair(pair, src, requiredSet)
		if f != nil {
			fields = append(fields, *f)
		}
	}
	return fields
}
