//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 점으로 구분된 리소스 이름을 URL 경로 세그먼트와 파라미터명으로 변환한다
package laravel

import "strings"

// buildResourcePath converts a dotted resource name to URL path segments.
// "users" -> ("users", "user")
// "users.posts" -> ("users/{user}/posts", "post")
func buildResourcePath(name string) (string, string) {
	name = strings.Trim(name, "/")
	parts := strings.Split(name, ".")
	if len(parts) == 1 {
		return parts[0], lastSegmentSingular(parts[0])
	}
	var path string
	for i, part := range parts {
		if i > 0 {
			path += "/" + part
		} else {
			path = part
		}
		if i < len(parts)-1 {
			path += "/{" + singularize(part) + "}"
		}
	}
	return path, lastSegmentSingular(parts[len(parts)-1])
}
