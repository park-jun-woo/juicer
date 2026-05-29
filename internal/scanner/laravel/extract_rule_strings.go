//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what rules 항목 값(파이프 문자열 또는 배열)에서 규칙 문자열 목록을 추출한다
package laravel

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// extractRuleStrings extracts rule strings from the value side of a rules entry.
// Supports both pipe-delimited strings ('required|string|max:255')
// and array syntax (['required', 'string', 'max:255']).
func extractRuleStrings(elem *sitter.Node, src []byte) []string {
	arr := findChildByType(elem, "array_creation_expression")
	if arr != nil {
		return extractStringArray(arr, src)
	}

	strNodes := childrenOfType(elem, "string")
	if len(strNodes) < 2 {
		return nil
	}
	valueStr := extractStringContent(strNodes[1], src)
	if valueStr == "" {
		return nil
	}
	return strings.Split(valueStr, "|")
}
