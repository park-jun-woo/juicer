//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what call/attribute 노드에서 form.get 키와 JSON 본문 존재를 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// collectCallFields scans call nodes for request.form.get(...) keys and marks
// hasJSON whenever a request.json / request.get_json() reference exists.
func collectCallFields(funcDef *sitter.Node, src []byte, bf bodyFields) bodyFields {
	for _, call := range findAllByType(funcDef, "call") {
		bf.formFields = appendUnique(bf.formFields, formGetKey(call, src))
		if isJSONSource(call, src) {
			bf.hasJSON = true
		}
	}
	for _, attr := range findAllByType(funcDef, "attribute") {
		if nodeText(attr, src) == "request.json" {
			bf.hasJSON = true
		}
	}
	return bf
}
