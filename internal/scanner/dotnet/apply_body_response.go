//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 본문 return 식 분석 결과를 endpointInfo의 타입·상태 코드에 반영한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func applyBodyResponse(m *sitter.Node, src []byte, ep *endpointInfo) {
	res := extractBodyResponse(m, src)
	if !res.found {
		return
	}
	if ep.statusCode == "" {
		ep.statusCode = res.status
	}
	if res.typeName != "" {
		ep.returnType = res.typeName
		ep.returnIsArray = res.isArray
	}
}
