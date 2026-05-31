//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 클래스 인덱스를 따라 부모를 전이적으로 walk하여 대상 집합 포함 여부를 판별한다
package django

// hasAncestorIn reports whether any name in parents — or, transitively, any of
// their ancestors resolved through idx — belongs to targets. A worklist with a
// visited set walks the inheritance graph while guarding against cycles. With a
// nil idx this checks only the direct parents.
func hasAncestorIn(parents []string, targets map[string]bool, idx classIndex) bool {
	visited := map[string]bool{}
	queue := append([]string(nil), parents...)
	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		if targets[name] {
			return true
		}
		if visited[name] {
			continue
		}
		visited[name] = true
		queue = append(queue, idx[name]...)
	}
	return false
}
