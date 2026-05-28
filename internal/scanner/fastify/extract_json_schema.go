//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 라우트 옵션 객체에서 schema.body/querystring/params/response를 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractJSONSchema(optsObj *sitter.Node, src []byte) *schemaInfo {
	if optsObj == nil {
		return nil
	}
	schemaNode := findPairValue(optsObj, src, "schema")
	if schemaNode == nil || schemaNode.Type() != "object" {
		return nil
	}
	si := &schemaInfo{Response: make(map[string]*sitter.Node)}
	for _, pair := range childrenOfType(schemaNode, "pair") {
		assignSchemaSection(si, pair, src)
	}
	return si
}
