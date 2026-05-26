//ff:func feature=scan type=convert control=iteration dimension=1 topic=fastapi
//ff:what Annotated[X, ...] 패턴에서 첫 번째 타입 인자를 추출한다
package fastapi

import "strings"

// tryAnnotated checks for Annotated[X, ...] and extracts the first type argument.
func tryAnnotated(py string) (openAPIType, bool) {
	if !strings.HasPrefix(py, "Annotated[") || !strings.HasSuffix(py, "]") {
		return openAPIType{}, false
	}
	inner := py[10 : len(py)-1]

	// Find the first comma at bracket depth 0.
	depth := 0
	for i, ch := range inner {
		depth += bracketDelta(ch)
		if ch == ',' && depth == 0 {
			first := strings.TrimSpace(inner[:i])
			return pyTypeToOpenAPI(first), true
		}
	}

	// No comma found — single argument like Annotated[int].
	first := strings.TrimSpace(inner)
	return pyTypeToOpenAPI(first), true
}
