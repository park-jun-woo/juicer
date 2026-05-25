//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what method+path 기준으로 중복 엔드포인트를 제거한다 (타입 정보가 풍부한 쪽 우선)
package scanner

func deduplicateEndpoints(endpoints []Endpoint) []Endpoint {
	type key struct{ method, path string }
	best := map[key]Endpoint{}
	order := []key{}

	for _, ep := range endpoints {
		k := key{ep.Method, ep.Path}
		if existing, ok := best[k]; ok {
			if richness(ep) > richness(existing) {
				best[k] = ep
			}
		} else {
			best[k] = ep
			order = append(order, k)
		}
	}

	result := make([]Endpoint, 0, len(order))
	for _, k := range order {
		result = append(result, best[k])
	}
	return result
}

