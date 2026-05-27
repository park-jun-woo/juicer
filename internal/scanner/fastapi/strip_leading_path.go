//ff:func feature=scan type=convert control=sequence topic=fastapi
//ff:what joinPath된 경로에서 선행 prefix 경로를 제거한다
package fastapi

import "strings"

// stripLeadingPath removes the leading path prefix from s.
// stripLeadingPath("/api/v1/items", "/api") = "/v1/items"
// stripLeadingPath("/api", "/api") = ""
// stripLeadingPath("/api/v1", "") = "/api/v1"
func stripLeadingPath(s, prefix string) string {
	if prefix == "" {
		return s
	}
	trimmed := strings.TrimPrefix(s, prefix)
	if trimmed == s {
		return s // prefix not found
	}
	trimmed = strings.TrimPrefix(trimmed, "/")
	if trimmed == "" {
		return ""
	}
	return "/" + trimmed
}
