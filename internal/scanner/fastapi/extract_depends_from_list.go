//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 리스트 노드에서 Depends(func) 호출의 함수명을 추출한다
package fastapi

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// extractDependsFromList extracts Depends function names from a list node
// containing Depends(...) call expressions.
func extractDependsFromList(listNode *sitter.Node, src []byte) []string {
	var deps []string
	calls := findAllByType(listNode, "call")
	for _, call := range calls {
		callText := nodeText(call, src)
		if !strings.HasPrefix(callText, "Depends(") {
			continue
		}
		fn := extractDependsFuncName(callText)
		if fn != "" {
			deps = append(deps, fn)
		}
	}
	return deps
}
