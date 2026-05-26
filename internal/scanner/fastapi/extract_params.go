//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 함수 시그니처에서 파라미터를 분류한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractParams analyzes the function signature and classifies parameters
// into path, query, body, file, and depends categories.
// aliasMap maps type alias names to their Depends function names.
func extractParams(funcDef *sitter.Node, src []byte, ri *routeInfo, aliasMap map[string]string) {
	params := findChildByType(funcDef, "parameters")
	if params == nil {
		return
	}
	pathNames := extractPathParamNames(ri.path)
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if isParamNode(child) {
			classifyParam(child, src, ri, pathNames, aliasMap)
		}
	}
}
