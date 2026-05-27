//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 중복 operationId 그룹에 경로 prefix를 추가하여 결과 맵에 기록한다
package scanner

import "strings"

func processDuplicateGroup(id string, indices []int, endpoints []Endpoint, result map[int]string, seen map[string]bool) {
	for _, idx := range indices {
		prefix := firstPathSegment(endpoints[idx].Path)
		prefixed := prefix + strings.ToUpper(id[:1]) + id[1:]
		prefixed = resolveSecondaryDuplicate(prefixed, seen)
		result[idx] = prefixed
		seen[prefixed] = true
	}
}
