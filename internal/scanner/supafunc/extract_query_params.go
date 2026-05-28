//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what searchParams.get("x") 호출에서 query parameter 이름을 추출한다
package supafunc

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func extractQueryParams(body *sitter.Node, src []byte) []string {
	var params []string
	seen := map[string]bool{}
	calls := findAllByType(body, "call_expression")
	for _, call := range calls {
		text := nodeText(call, src)
		if !strings.Contains(text, "searchParams.get") {
			continue
		}
		name := extractParamFromCall(call, src)
		if name != "" && !seen[name] {
			seen[name] = true
			params = append(params, name)
		}
	}
	return params
}
