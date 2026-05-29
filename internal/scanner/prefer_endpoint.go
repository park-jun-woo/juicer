//ff:func feature=scan type=extract control=sequence
//ff:what 같은 method+path의 두 후보 중 우선할 쪽을 결정한다 (richness 우선, 동점 시 File→Line 안정적 tie-break)
package scanner

// preferEndpoint은 candidate가 current보다 우선되어야 하면 true를 반환한다.
// 1순위: richness(타입 정보 풍부도)가 더 큰 쪽.
// 동점: 입력 슬라이스 순서에 의존하지 않도록 (File, Line)이 더 작은 쪽을 선택해
// 결정성을 보장한다.
func preferEndpoint(candidate, current Endpoint) bool {
	rc, re := richness(candidate), richness(current)
	if rc != re {
		return rc > re
	}
	if candidate.File != current.File {
		return candidate.File < current.File
	}
	return candidate.Line < current.Line
}
