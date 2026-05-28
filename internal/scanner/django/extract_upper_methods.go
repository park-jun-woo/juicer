//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 키워드 인자에서 메서드 문자열 리스트를 추출하고 대문자로 변환한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// extractUpperMethods extracts and uppercases method strings from a list kwarg.
func extractUpperMethods(child *sitter.Node, src []byte) []string {
	listNode := findChildByType(child, "list")
	if listNode == nil {
		return nil
	}
	methods := extractStringList(listNode, src)
	for idx, m := range methods {
		methods[idx] = strings.ToUpper(m)
	}
	return methods
}
