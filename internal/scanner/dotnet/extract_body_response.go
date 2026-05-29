//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 메서드 본문의 return 식에서 응답 타입·상태 코드를 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractBodyResponse(m *sitter.Node, src []byte) bodyResponse {
	body := findChildByType(m, "block")
	if body == nil {
		return bodyResponse{}
	}
	return selectBodyResponse(body, src)
}
