//ff:func feature=scan type=test control=iteration dimension=1 topic=supafunc
//ff:what 테스트 헬퍼: method와 path로 endpoint를 찾는다
package supafunc

import "github.com/park-jun-woo/codistill/internal/scanner"

func findEndpoint(eps []scanner.Endpoint, method, path string) *scanner.Endpoint {
	for i := range eps {
		if eps[i].Method == method && eps[i].Path == path {
			return &eps[i]
		}
	}
	return nil
}
