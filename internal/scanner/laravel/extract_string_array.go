//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 배열 리터럴에서 문자열 값들을 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractStringArray extracts string values from an array literal.
func extractStringArray(arr *sitter.Node, src []byte) []string {
	var result []string
	for _, elem := range childrenOfType(arr, "array_element_initializer") {
		s := extractStringContent(elem, src)
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
