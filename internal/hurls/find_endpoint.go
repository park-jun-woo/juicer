//ff:func feature=hurl type=parse control=iteration dimension=1
//ff:what scan 결과에서 엔드포인트 ID에 매칭되는 Endpoint 탐색
package hurls

import (
	"github.com/park-jun-woo/juicer/scanner"
)

// findEndpoint returns the scanner.Endpoint matching the given id ("METHOD /path").
func findEndpoint(endpoints []scanner.Endpoint, id string) *scanner.Endpoint {
	for i, ep := range endpoints {
		if ep.Method+" "+ep.Path == id {
			return &endpoints[i]
		}
	}
	return nil
}
