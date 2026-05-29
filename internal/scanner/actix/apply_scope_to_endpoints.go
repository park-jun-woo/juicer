//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 한 스코프의 핸들러들에 대해 해당 엔드포인트 경로에 prefix를 적용한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyScopeToEndpoints(scope scopeInfo, endpoints []scanner.Endpoint) {
	for _, handlerName := range scope.handlers {
		prefixEndpointsForHandler(endpoints, handlerName, scope.prefix)
	}
}
