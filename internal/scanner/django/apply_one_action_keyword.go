//ff:func feature=scan type=extract control=selection topic=django
//ff:what @action 데코레이터의 단일 키워드 인자를 적용한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// applyOneActionKeyword applies a single keyword argument to actionInfo.
func applyOneActionKeyword(ai *actionInfo, key string, child *sitter.Node, src []byte) {
	switch key {
	case "detail":
		ai.detail = hasTrue(child, src)
	case "methods":
		ai.methods = extractUpperMethods(child, src)
	case "url_path":
		valNode := findChildByType(child, "string")
		if valNode != nil {
			ai.urlPath = unquotePython(nodeText(valNode, src))
		}
	}
}
