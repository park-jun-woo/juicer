//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what Annotated[T, Depends(func)] 텍스트에서 Depends 함수명을 추출한다
package fastapi

import "strings"

// extractDependsFromAnnotated extracts the Depends function name from an
// Annotated[T, Depends(func)] text. Returns the func name, or "" if Depends
// is not found. For empty Depends(), returns "Depends" as a sentinel.
func extractDependsFromAnnotated(text string) string {
	idx := strings.Index(text, "Depends(")
	if idx < 0 {
		return ""
	}
	inner := text[idx+8:]
	closeIdx := strings.Index(inner, ")")
	if closeIdx < 0 {
		return ""
	}
	fn := strings.TrimSpace(inner[:closeIdx])
	if fn == "" {
		return "Depends"
	}
	return fn
}
