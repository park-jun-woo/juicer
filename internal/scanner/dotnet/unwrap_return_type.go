//ff:func feature=scan type=convert control=sequence topic=dotnet
//ff:what ActionResult, Task, List 등의 래퍼를 벗겨내어 실제 반환 타입을 추출한다
package dotnet

import "strings"

func unwrapReturnType(raw string) (string, bool) {
	if strings.HasPrefix(raw, "ActionResult<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[len("ActionResult<") : len(raw)-1])
		return unwrapReturnType(inner)
	}
	if strings.HasPrefix(raw, "Task<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[5 : len(raw)-1])
		return unwrapReturnType(inner)
	}
	if strings.HasPrefix(raw, "List<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[5 : len(raw)-1])
		return inner, true
	}
	if strings.HasPrefix(raw, "IEnumerable<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[12 : len(raw)-1])
		return inner, true
	}
	if strings.HasPrefix(raw, "IList<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[6 : len(raw)-1])
		return inner, true
	}
	if strings.HasPrefix(raw, "ICollection<") && strings.HasSuffix(raw, ">") {
		inner := strings.TrimSpace(raw[12 : len(raw)-1])
		return inner, true
	}
	if strings.HasSuffix(raw, "[]") {
		return raw[:len(raw)-2], true
	}
	return raw, false
}
