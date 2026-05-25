//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 엔드포인트의 타입 정보 풍부도를 점수로 반환한다
package scanner

func richness(ep Endpoint) int {
	score := 0
	if ep.Request != nil && ep.Request.Body != nil {
		if ep.Request.Body.TypeName != "" {
			score += 3
		}
		score += len(ep.Request.Body.Fields)
	}
	for _, r := range ep.Responses {
		if r.TypeName != "" {
			score += 2
		}
		score += len(r.Fields)
	}
	return score
}

