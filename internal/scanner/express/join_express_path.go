//ff:func feature=scan type=extract control=sequence topic=express
//ff:what prefix와 경로를 결합한다 (슬래시 중복 방지)
package express

import "strings"

func joinExpressPath(prefix, path string) string {
	if prefix == "" {
		return path
	}
	if path == "/" || path == "" {
		return prefix
	}
	return strings.TrimRight(prefix, "/") + "/" + strings.TrimLeft(path, "/")
}
