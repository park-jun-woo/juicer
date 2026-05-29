//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 의존성이 모두 출력된 미출력 노드 중 알파벳 우선 항목 반환
package ddl

// pickReady returns the alphabetically-first not-yet-emitted node whose
// dependencies are all emitted, or "" when none qualifies.
func pickReady(names []string, deps map[string]map[string]bool, emitted map[string]bool) string {
	for _, name := range names {
		if emitted[name] {
			continue
		}
		if depsSatisfied(deps[name], emitted) {
			return name
		}
	}
	return ""
}
