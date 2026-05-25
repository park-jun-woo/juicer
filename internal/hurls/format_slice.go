//ff:func feature=hurl type=render control=sequence
//ff:what 문자열 슬라이스를 "[a, b, c]" 형식으로 포맷
package hurls

import (
	"strings"
)

// formatSlice formats a string slice as "[a, b, c]".
func formatSlice(s []string) string {
	return "[" + strings.Join(s, ", ") + "]"
}
