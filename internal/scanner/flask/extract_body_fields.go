//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what Flask 핸들러 본문에서 form/json 바디 필드를 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// extractBodyFields walks a handler's function_definition body and collects
// request.form / request.json access keys, classified by content type.
func extractBodyFields(funcDef *sitter.Node, src []byte) bodyFields {
	jsonVars := collectJSONVars(funcDef, src)

	var bf bodyFields
	bf = collectSubscriptFields(funcDef, src, jsonVars, bf)
	bf = collectCallFields(funcDef, src, bf)
	return bf
}
