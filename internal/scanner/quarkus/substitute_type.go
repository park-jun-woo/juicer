//ff:func feature=scan type=convert control=iteration dimension=1 topic=quarkus
//ff:what 단일 타입 문자열의 타입 파라미터를 실제 타입으로 치환한다
package quarkus

import "strings"

func substituteType(t string, paramMap map[string]string) string {
	if replacement, ok := paramMap[t]; ok {
		return replacement
	}
	if strings.HasPrefix(t, "[]") {
		if replacement, ok := paramMap[t[2:]]; ok {
			return "[]" + replacement
		}
	}
	if strings.HasPrefix(t, "array:") {
		if replacement, ok := paramMap[t[6:]]; ok {
			return "array:" + replacement
		}
	}
	idx := strings.Index(t, "<")
	if idx < 0 {
		return t
	}
	outer := t[:idx]
	inner := t[idx+1 : len(t)-1]
	parts := splitGenericArgs(inner)
	for i, p := range parts {
		parts[i] = substituteType(strings.TrimSpace(p), paramMap)
	}
	return outer + "<" + strings.Join(parts, ", ") + ">"
}
