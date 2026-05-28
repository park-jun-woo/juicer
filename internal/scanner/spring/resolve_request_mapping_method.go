//ff:func feature=scan type=convert control=sequence topic=spring
//ff:what @RequestMapping의 method 속성을 HTTP 메서드 문자열로 변환한다
package spring

import "strings"

func resolveRequestMappingMethod(s string) string {
	s = strings.TrimSpace(s)
	if m, ok := requestMappingMethods[s]; ok {
		return m
	}
	s = strings.TrimPrefix(s, "{")
	s = strings.TrimSuffix(s, "}")
	s = strings.TrimSpace(s)
	if m, ok := requestMappingMethods[s]; ok {
		return m
	}
	return ""
}
