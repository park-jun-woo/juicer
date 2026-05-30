//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what classifyAllParams 테스트 헬퍼
package fastapi

func classifyAllParams(src []byte, ri *routeInfo, pathNames map[string]bool, aliasMap map[string]string) {
	root, err := parsePython(src)
	if err != nil {
		return
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if isParamNode(child) {
			classifyParam(child, src, ri, pathNames, aliasMap)
		}
	}
}
