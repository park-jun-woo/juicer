//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 핸들러명이 일치하고 아직 prefix가 없는 엔드포인트 경로에 prefix를 붙인다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func prefixEndpointsForHandler(endpoints []scanner.Endpoint, handlerName, prefix string) {
	for i := range endpoints {
		if endpoints[i].Handler == handlerName && !hasPrefix(endpoints[i].Path, prefix) {
			endpoints[i].Path = joinPath(prefix, endpoints[i].Path)
		}
	}
}
