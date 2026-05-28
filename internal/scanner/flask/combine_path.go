//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what 라우터 접두사와 라우트 경로를 결합한다
package flask

import "strings"

// combinePath joins a prefix and a route path.
func combinePath(prefix, path string) string {
	prefix = strings.TrimRight(prefix, "/")
	if path == "" || path == "/" {
		if prefix == "" {
			return "/"
		}
		return prefix
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return prefix + path
}
