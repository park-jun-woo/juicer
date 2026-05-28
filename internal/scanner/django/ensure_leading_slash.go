//ff:func feature=scan type=convert control=sequence topic=django
//ff:what 경로가 /로 시작하도록 보장한다
package django

import "strings"

// ensureLeadingSlash ensures the path starts with "/".
func ensureLeadingSlash(path string) string {
	if path == "" {
		return "/"
	}
	if !strings.HasPrefix(path, "/") {
		return "/" + path
	}
	return path
}
