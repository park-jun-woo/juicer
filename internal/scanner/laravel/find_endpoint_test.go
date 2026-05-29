//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what findEndpoint — method+path가 일치하는 첫 엔드포인트를 반환한다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func findEndpoint(endpoints []scanner.Endpoint, method, path string) *scanner.Endpoint {
	for i := range endpoints {
		if endpoints[i].Method == method && endpoints[i].Path == path {
			return &endpoints[i]
		}
	}
	return nil
}
