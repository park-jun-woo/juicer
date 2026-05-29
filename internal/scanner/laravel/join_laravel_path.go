//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what prefix와 path 세그먼트를 정리된 URL 경로로 결합한다
package laravel

import "strings"

// joinLaravelPath joins prefix and path segments into a clean URL path.
func joinLaravelPath(prefix, path string) string {
	prefix = strings.TrimRight(prefix, "/")
	path = strings.TrimLeft(path, "/")
	if prefix == "" && path == "" {
		return "/"
	}
	if prefix == "" {
		return "/" + path
	}
	if path == "" {
		return "/" + prefix
	}
	return "/" + prefix + "/" + path
}
