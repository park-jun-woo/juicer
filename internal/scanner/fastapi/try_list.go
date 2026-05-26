//ff:func feature=scan type=convert control=sequence topic=fastapi
//ff:what list[X], List[X] 패턴에서 배열 타입을 추출한다
package fastapi

import "strings"

// tryList checks for list[X] and List[X] patterns.
func tryList(py string) (openAPIType, bool) {
	lower := strings.ToLower(py)
	if (strings.HasPrefix(lower, "list[") || strings.HasPrefix(py, "List[")) &&
		strings.HasSuffix(py, "]") {
		inner := strings.TrimSpace(py[5 : len(py)-1])
		return openAPIType{Type: "array", Items: inner}, true
	}
	return openAPIType{}, false
}
