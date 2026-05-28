//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 메서드 파라미터에서 @PathParam, @QueryParam, body 등을 분류한다
package quarkus

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
