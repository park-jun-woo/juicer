//ff:func feature=scan type=parse control=iteration dimension=1 topic=joi
//ff:what `{ body: Joi..., query: Joi..., params: Joi... }` object를 RequestSchema로 변환한다
package joi

import sitter "github.com/smacker/go-tree-sitter"

// ParseRequestObject — Joi 검증 const의 object 리터럴 → RequestSchema.
// body/query/params 키의 value(각 Joi.object().keys(...))를 필드 슬라이스로 파싱한다.
func ParseRequestObject(objNode *sitter.Node, src []byte) RequestSchema {
	var rs RequestSchema
	if objNode == nil || objNode.Type() != "object" {
		return rs
	}
	for _, pair := range childrenOfType(objNode, "pair") {
		keyNode := pair.ChildByFieldName("key")
		valueNode := pair.ChildByFieldName("value")
		if keyNode == nil || valueNode == nil {
			continue
		}
		fields := ParseSchema(valueNode, src)
		switch unquoteJoi(nodeText(keyNode, src)) {
		case "body":
			rs.Body = fields
		case "query":
			rs.Query = fields
		case "params", "param":
			rs.Params = fields
		}
	}
	return rs
}
