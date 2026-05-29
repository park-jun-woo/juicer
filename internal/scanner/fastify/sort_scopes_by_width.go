//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what wrapper 스코프를 너비(바깥쪽 먼저) 내림차순으로 정렬한다
package fastify

import "sort"

func sortScopesByWidth(scopes []wrapperScope) {
	sort.SliceStable(scopes, func(i, j int) bool {
		wi := scopes[i].End - scopes[i].Start
		wj := scopes[j].End - scopes[j].Start
		return wi > wj
	})
}
