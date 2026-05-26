//ff:func feature=scan type=convert control=sequence topic=fastapi
//ff:what Optional[X], Union[X, None], X | None 패턴에서 nullable 타입을 추출한다
package fastapi

import "strings"

// tryNullable checks for Optional[X], Union[X, None], and X | None patterns.
func tryNullable(py string) (openAPIType, bool) {
	if strings.HasPrefix(py, "Optional[") && strings.HasSuffix(py, "]") {
		inner := strings.TrimSpace(py[9 : len(py)-1])
		result := pyTypeToOpenAPI(inner)
		result.Nullable = true
		return result, true
	}
	if strings.HasPrefix(py, "Union[") && strings.HasSuffix(py, "]") {
		return tryUnionNullable(py)
	}
	if strings.Contains(py, "|") {
		return tryPipeNullable(py)
	}
	return openAPIType{}, false
}
