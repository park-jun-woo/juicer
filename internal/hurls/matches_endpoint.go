//ff:func feature=hurl type=parse control=iteration dimension=1
//ff:what hurl 파일 내용이 주어진 method+path를 참조하는지 확인
package hurls

import (
	"strings"
)

// matchesEndpoint checks if hurl file content references the given method+path.
func matchesEndpoint(content, method, path string) bool {
	patterns := []string{
		method + " {{host}}" + path,
		method + " https://{{host}}" + path,
		method + " http://{{host}}" + path,
	}
	for _, pat := range patterns {
		if strings.Contains(content, pat) {
			return true
		}
	}
	return false
}
