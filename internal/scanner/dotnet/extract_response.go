//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 메서드 반환 타입에서 응답 정보를 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractReturnInfo(m *sitter.Node, src []byte, ep *endpointInfo) {
	for i := 0; i < int(m.ChildCount()); i++ {
		if matchReturnType(m.Child(i), src, ep) {
			break
		}
	}
	if ep.returnType == "" {
		applyBodyResponse(m, src, ep)
	}
}
