//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what 문자열 슬라이스에 중복 없이 값을 추가한다
package flask

// appendUnique appends s to list only if it is non-empty and not already present.
func appendUnique(list []string, s string) []string {
	if s == "" {
		return list
	}
	for _, e := range list {
		if e == s {
			return list
		}
	}
	return append(list, s)
}
