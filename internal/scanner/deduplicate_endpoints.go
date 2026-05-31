//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what method+path+출처파일 기준으로 중복 엔드포인트를 제거한다 (타입 정보가 풍부한 쪽 우선)
package scanner

// DeduplicateEndpoints — 동일 (method, path, file) 엔드포인트를 하나로 합친다.
//
// dedup 키에 출처 파일(File)을 포함하는 이유 (BUG-009 / Phase138-A, 라우트 소실 안전망):
// 마운트 prefix가 합성되지 않아 서로 다른 라우터의 라우트가 같은 bare path(예: `/posts`)로
// 방출될 때, (method, path)만으로 dedup하면 출처가 다른 라우트가 폐기되어 소실된다.
// 출처 파일을 키에 더하면 그런 cross-file 충돌에서 양쪽이 모두 보존된다.
// 같은 파일(또는 File 미설정) 내 동일 (method, path)는 종전대로 richer 우선으로 합쳐진다.
func DeduplicateEndpoints(endpoints []Endpoint) []Endpoint {
	type key struct{ method, path, file string }
	best := map[key]Endpoint{}
	order := []key{}

	for _, ep := range endpoints {
		k := key{ep.Method, ep.Path, ep.File}
		if existing, ok := best[k]; ok {
			if preferEndpoint(ep, existing) {
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
