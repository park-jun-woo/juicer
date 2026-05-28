//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 메서드 파라미터에서 [FromBody], [FromQuery], [FromRoute] 등을 분류한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractMethodParams(m *sitter.Node, src []byte, ep *endpointInfo) {
	params := findChildByType(m, "parameter_list")
	if params == nil {
		return
	}
	for _, param := range childrenOfType(params, "parameter") {
		classifyParam(param, src, ep)
	}
}
