//ff:func feature=scan type=convert control=sequence topic=quarkus
//ff:what Response, Uni, Multi, List 등의 래퍼를 벗겨내어 실제 반환 타입을 추출한다
package quarkus

import "strings"

func unwrapReturnType(raw string) (string, bool) {
	if raw == "void" {
		return "", false
	}
	if raw == "Response" {
		return "", false
	}
	if strings.HasPrefix(raw, "Uni<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[4 : len(raw)-1])
		return unwrapReturnType(inner)
	}
	if strings.HasPrefix(raw, "Multi<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[6 : len(raw)-1])
		return inner, true
	}
	if strings.HasPrefix(raw, "CompletionStage<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[16 : len(raw)-1])
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
