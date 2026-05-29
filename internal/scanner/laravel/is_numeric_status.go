//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 문자열이 비어있지 않고 모두 숫자로 된 상태 코드인지 검사한다
package laravel

func isNumericStatus(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
