//ff:func feature=scan type=extract control=selection topic=fastapi
//ff:what 기본값 함수 호출명에 따라 파라미터를 분류한다
package fastapi

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// classifyByDefault classifies a parameter by its default call function name.
// param is the tree-sitter parameter node (may be nil in tests).
func classifyByDefault(defaultCall, name, typeName, defaultVal string, param *sitter.Node, src []byte, ri *routeInfo) {
	switch specialDefaults[defaultCall] {
	case "query":
		ri.query = append(ri.query, scanner.Param{Name: name, Type: mapTypeToOpenAPI(typeName)})
	case "body":
		ri.bodyType = typeName
		ri.bodyVarName = name
		if param != nil {
			ri.bodyAlias, ri.bodyEmbed = extractBodyKwargs(param, src)
		}
	case "file":
		ri.files = append(ri.files, scanner.Param{Name: name, Type: "file"})
	case "depends":
		handleDepends(name, defaultVal, ri)
	}
}
