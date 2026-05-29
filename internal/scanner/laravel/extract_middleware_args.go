//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 인자 목록에서 미들웨어 이름들(단일 문자열 또는 배열)을 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractMiddlewareArgs extracts middleware names from argument list.
func extractMiddlewareArgs(args *sitter.Node, fi fileInfo) []string {
	argNodes := childrenOfType(args, "argument")
	if len(argNodes) == 0 {
		return nil
	}
	arr := findAllByType(argNodes[0], "array_creation_expression")
	if len(arr) > 0 {
		return extractStringArray(arr[0], fi.src)
	}
	s := extractStringContent(argNodes[0], fi.src)
	if s != "" {
		return []string{s}
	}
	return nil
}
