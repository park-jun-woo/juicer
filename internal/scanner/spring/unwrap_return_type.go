//ff:func feature=scan type=convert control=sequence topic=spring
//ff:what ResponseEntity, List 등의 래퍼를 벗겨내어 실제 반환 타입을 추출한다
package spring

import "strings"

func unwrapReturnType(raw string) (string, bool) {
	if raw == "void" {
		return "", false
	}
	if strings.HasPrefix(raw, "ResponseEntity<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[len("ResponseEntity<") : len(raw)-1])
		if inner == "?" || inner == "Void" || inner == "void" {
			return "", false
		}
		return unwrapReturnType(inner)
	}
	if strings.HasPrefix(raw, "List<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[5 : len(raw)-1])
		return inner, true
	}
	if strings.HasPrefix(raw, "Set<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[4 : len(raw)-1])
		return inner, true
	}
	if strings.HasPrefix(raw, "Collection<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[11 : len(raw)-1])
		return inner, true
	}
	if strings.HasSuffix(raw, "[]") {
		return raw[:len(raw)-2], true
	}
	return raw, false
}
