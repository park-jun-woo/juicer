//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 본문의 return 식 후보 중 성공(2xx) 응답을 우선 선택한다
package dotnet

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func selectBodyResponse(body *sitter.Node, src []byte) bodyResponse {
	var first bodyResponse
	for _, ret := range findAllByType(body, "return_statement") {
		cand := matchReturnStatement(ret, src)
		if !cand.found {
			continue
		}
		if strings.HasPrefix(cand.status, "2") {
			return cand
		}
		if !first.found {
			first = cand
		}
	}
	return first
}
