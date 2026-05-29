//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 패키지 경로가 Echo 패키지(v4/v5) suffix를 가지는지 검사한다
package echo

import "strings"

func hasEchoPkgSuffix(path string) bool {
	for _, p := range echoPkgPaths {
		if strings.HasSuffix(path, p) {
			return true
		}
	}
	return false
}
