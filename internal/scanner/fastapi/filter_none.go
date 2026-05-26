//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 문자열 슬라이스에서 None 항목을 제거한다
package fastapi

import "strings"

// filterNone removes "None" entries from a string slice.
func filterNone(parts []string) []string {
	var result []string
	for _, p := range parts {
		if strings.TrimSpace(p) != "None" {
			result = append(result, strings.TrimSpace(p))
		}
	}
	return result
}
