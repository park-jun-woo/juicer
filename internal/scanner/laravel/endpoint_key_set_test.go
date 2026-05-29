//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what endpointKeySet — 엔드포인트를 "method path" 키 집합으로 만든다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func endpointKeySet(endpoints []scanner.Endpoint) map[string]bool {
	found := map[string]bool{}
	for _, ep := range endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	return found
}
