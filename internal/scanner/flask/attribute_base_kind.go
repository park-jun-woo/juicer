//ff:func feature=scan type=extract control=selection topic=flask
//ff:what attribute 베이스 텍스트를 form/json 으로 분류한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// attributeBaseKind classifies an attribute base node (request.form / request.json).
func attributeBaseKind(base *sitter.Node, src []byte) string {
	switch nodeText(base, src) {
	case "request.form":
		return "form"
	case "request.json":
		return "json"
	default:
		return ""
	}
}
