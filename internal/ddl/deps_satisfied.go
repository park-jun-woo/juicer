//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 노드의 모든 의존성이 이미 출력되었는지 검사
package ddl

// depsSatisfied reports whether every dependency of a node is already emitted.
func depsSatisfied(set, emitted map[string]bool) bool {
	for dep := range set {
		if !emitted[dep] {
			return false
		}
	}
	return true
}
