//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what subscript 노드들에서 form/json 바디 필드 키를 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// collectSubscriptFields scans every subscript node, routing string keys into
// form or json fields per the subscript base classification.
func collectSubscriptFields(funcDef *sitter.Node, src []byte, jsonVars map[string]bool, bf bodyFields) bodyFields {
	for _, sub := range findAllByType(funcDef, "subscript") {
		kind := subscriptKind(sub, src, jsonVars)
		if kind == "" {
			continue
		}
		key := subscriptStringKey(sub, src)
		bf = applyBodyKey(bf, kind, key)
	}
	return bf
}
