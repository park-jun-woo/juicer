//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what epSet — 엔드포인트 목록을 "method path" 키 집합으로 변환한다
package express

func epSet(eps []endpointLike) map[string]bool {
	m := map[string]bool{}
	for _, e := range eps {
		m[e.method+" "+e.path] = true
	}
	return m
}
