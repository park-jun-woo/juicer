//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 메서드 파라미터에서 @PathVariable, @RequestParam, @RequestBody 등을 분류한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractMethodParams(m *sitter.Node, src []byte, ep *endpointInfo, imports map[string]string, referrerPath, projectRoot string) {
	params := findChildByType(m, "formal_parameters")
	if params == nil {
		return
	}
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if child.Type() == "formal_parameter" || child.Type() == "spread_parameter" {
			classifyParam(child, src, ep, imports, referrerPath, projectRoot)
		}
	}
}
