//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what import 경로가 Echo 패키지(v4/v5)인지 검사한다
package echo

func isEchoPkgPath(path string) bool {
	for _, p := range echoPkgPaths {
		if path == p {
			return true
		}
	}
	return false
}
