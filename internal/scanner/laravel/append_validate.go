//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what 기존 validate 문자열에 규칙을 콤마로 이어붙인다
package laravel

// appendValidate appends a validation rule to the existing validate string.
func appendValidate(existing, rule string) string {
	if existing == "" {
		return rule
	}
	return existing + "," + rule
}
