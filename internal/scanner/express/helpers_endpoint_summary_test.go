//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 테스트 헬퍼: 엔드포인트 목록을 method+path 문자열로 요약한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

func endpointSummary(eps []scanner.Endpoint) []string {
	var out []string
	for _, ep := range eps {
		out = append(out, ep.Method+" "+ep.Path)
	}
	return out
}
