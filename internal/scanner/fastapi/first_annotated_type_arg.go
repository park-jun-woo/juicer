//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what Annotated[T, ...] 텍스트에서 첫 번째 타입 인자를 추출한다
package fastapi

import "strings"

// firstAnnotatedTypeArg extracts the first type argument from "Annotated[T, ...]".
// Returns "" if the text is not an Annotated pattern.
func firstAnnotatedTypeArg(text string) string {
	const prefix = "Annotated["
	if !strings.HasPrefix(text, prefix) {
		return ""
	}
	inner := text[len(prefix):]
	depth := 0
	for i, ch := range inner {
		depth += bracketDelta(ch)
		if depth < 0 {
			return strings.TrimSpace(inner[:i])
		}
		if ch == ',' && depth == 0 {
			return strings.TrimSpace(inner[:i])
		}
	}
	return ""
}
