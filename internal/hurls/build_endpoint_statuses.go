//ff:func feature=hurl type=parse control=iteration dimension=1
//ff:what 정렬된 엔드포인트에서 EndpointStatus 슬라이스 생성
package hurls

import (
	"github.com/park-jun-woo/juicer/scanner"
)

// buildEndpointStatuses converts sorted scanner endpoints to EndpointStatus slice.
func buildEndpointStatuses(sorted []scanner.Endpoint) []EndpointStatus {
	statuses := make([]EndpointStatus, len(sorted))
	for i, ep := range sorted {
		statuses[i] = EndpointStatus{
			ID:     ep.Method + " " + ep.Path,
			Status: "TODO",
		}
	}
	return statuses
}
