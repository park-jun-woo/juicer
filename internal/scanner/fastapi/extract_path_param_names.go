//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 라우트 경로에서 경로 파라미터 이름을 추출한다
package fastapi

import "strings"

// extractPathParamNames extracts path parameter names from a route path.
// e.g., "/users/{user_id}/posts/{post_id}" -> {"user_id": true, "post_id": true}
func extractPathParamNames(path string) map[string]bool {
	names := make(map[string]bool)
	for {
		start := strings.Index(path, "{")
		if start < 0 {
			break
		}
		end := strings.Index(path[start:], "}")
		if end < 0 {
			break
		}
		name := path[start+1 : start+end]
		if name != "" {
			names[name] = true
		}
		path = path[start+end+1:]
	}
	return names
}
