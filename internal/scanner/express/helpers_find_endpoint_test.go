//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 테스트 헬퍼: 엔드포인트 목록에서 method+path로 검색한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

func findEndpoint(endpoints []scanner.Endpoint, method, path string) *scanner.Endpoint {
	for i := range endpoints {
		if endpoints[i].Method == method && endpoints[i].Path == path {
			return &endpoints[i]
		}
	}
	return nil
}
