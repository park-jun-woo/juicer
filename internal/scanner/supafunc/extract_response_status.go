//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what new Response(..., { status: N }) 패턴에서 HTTP 상태 코드를 추출한다
package supafunc

import sitter "github.com/smacker/go-tree-sitter"

func extractResponseStatus(body *sitter.Node, src []byte) []string {
	var statuses []string
	seen := map[string]bool{}
	newExprs := findAllByType(body, "new_expression")
	for _, ne := range newExprs {
		id := findChildByType(ne, "identifier")
		if id == nil || nodeText(id, src) != "Response" {
			continue
		}
		status := extractStatusFromResponse(ne, src)
		if status != "" && !seen[status] {
			seen[status] = true
			statuses = append(statuses, status)
		}
	}
	return statuses
}
