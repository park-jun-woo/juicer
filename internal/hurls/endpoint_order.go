//ff:func feature=hurl type=parse control=selection
//ff:what 엔드포인트의 의존성 기반 정렬 키 반환
package hurls

import (
	"github.com/park-jun-woo/juicer/scanner"
)

// endpointOrder returns a sort key for dependency-based ordering.
func endpointOrder(ep scanner.Endpoint) int {
	if isPublicEndpoint(ep) {
		return 0
	}
	switch ep.Method {
	case "GET":
		return 1
	case "POST", "DELETE":
		return 2
	default:
		return 3
	}
}
