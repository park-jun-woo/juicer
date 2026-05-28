//ff:func feature=scan type=convert control=sequence topic=django
//ff:what 라우터 접두사와 라우트 경로를 결합한다
package django

import "strings"

// combinePath joins a prefix and a route path with a leading slash.
func combinePath(prefix, path string) string {
	prefix = strings.TrimRight(prefix, "/")
	path = strings.TrimLeft(path, "/")
	if path == "" {
		if prefix == "" {
			return "/"
		}
		return prefix
	}
	if prefix == "" {
		return "/" + path
	}
	return prefix + "/" + path
}
