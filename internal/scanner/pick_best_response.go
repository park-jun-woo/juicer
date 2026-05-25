//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 같은 status의 응답 중 타입 정보가 가장 풍부한 것을 선택한다
package scanner

func pickBestResponse(resps []Response) Response {
	if len(resps) == 1 {
		return resps[0]
	}
	best := resps[0]
	bestScore := len(best.Fields)
	if best.TypeName != "" {
		bestScore += 2
	}
	for _, r := range resps[1:] {
		score := len(r.Fields)
		if r.TypeName != "" {
			score += 2
		}
		if score > bestScore {
			best = r
			bestScore = score
		}
	}
	return best
}

