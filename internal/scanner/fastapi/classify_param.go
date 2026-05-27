//ff:func feature=scan type=extract control=selection topic=fastapi
//ff:what 단일 함수 파라미터를 path/query/body/file/depends 로 분류한다
package fastapi

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// classifyParam classifies a single function parameter.
// aliasMap maps type alias names (e.g. "SessionDep") to their Depends function
// names, built by resolveTypeAliases.
func classifyParam(param *sitter.Node, src []byte, ri *routeInfo, pathNames map[string]bool, aliasMap map[string]string) {
	name, typeName, defaultVal, defaultCall, isNone := parseParamNode(param, src)
	if name == "" || name == "self" || name == "cls" {
		return
	}

	switch {
	case defaultCall != "" && specialDefaults[defaultCall] != "":
		classifyByDefault(defaultCall, name, typeName, defaultVal, param, src, ri)
	case isAnnotatedDepends(typeName, aliasMap):
		classifyAnnotatedDepends(name, typeName, aliasMap, ri)
	case uploadFileTypes[typeName]:
		ri.files = append(ri.files, scanner.Param{Name: name, Type: "file"})
	case pathNames[name]:
		ri.params = append(ri.params, scanner.Param{Name: name, Type: mapTypeToOpenAPI(typeName)})
	case isPydanticModelType(typeName):
		ri.bodyType = typeName
		ri.bodyVarName = name
	case defaultVal != "":
		ri.query = append(ri.query, scanner.Param{Name: name, Type: mapTypeToOpenAPI(typeName), Default: defaultVal})
	case isNone:
		ri.query = append(ri.query, scanner.Param{Name: name, Type: mapTypeToOpenAPI(typeName), DefaultIsNull: true})
	case typeName != "":
		ri.query = append(ri.query, scanner.Param{Name: name, Type: mapTypeToOpenAPI(typeName)})
	}
}
