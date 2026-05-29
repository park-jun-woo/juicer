//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 'middleware' 옵션 값(문자열 또는 배열)에서 미들웨어 이름들을 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// middlewareValues extracts middleware names from a 'middleware' option value,
// accepting either a single string or an array_creation_expression of strings.
func middlewareValues(value *sitter.Node, fi fileInfo) []string {
	if value.Type() == "array_creation_expression" {
		return extractStringArray(value, fi.src)
	}
	if s := extractStringContent(value, fi.src); s != "" {
		return []string{s}
	}
	return nil
}
