//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 중복되지 않은 actionMethod만 추가한다
package django

// appendUnseenMethods appends actionMethods that haven't been seen yet.
func appendUnseenMethods(methods []actionMethod, ms []actionMethod, seen map[string]bool) []actionMethod {
	for _, m := range ms {
		key := m.action + m.method
		if !seen[key] {
			methods = append(methods, m)
			seen[key] = true
		}
	}
	return methods
}
