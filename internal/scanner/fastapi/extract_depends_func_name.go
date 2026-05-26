//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what Depends(...) 호출 텍스트에서 함수명을 추출한다
package fastapi

import "strings"

// extractDependsFuncName extracts the function name from a Depends(...) call string.
// e.g., "Depends(get_current_user)" -> "get_current_user"
func extractDependsFuncName(callText string) string {
	if !strings.HasPrefix(callText, "Depends(") {
		return ""
	}
	inner := callText[8:]
	if strings.HasSuffix(inner, ")") {
		inner = inner[:len(inner)-1]
	}
	return strings.TrimSpace(inner)
}
