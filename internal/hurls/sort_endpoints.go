//ff:func feature=hurl type=parse control=sequence
//ff:what 엔드포인트를 의존성 순서로 정렬 (공개 > GET > POST+DELETE > PUT/PATCH)
package hurls

import (
	"sort"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// sortEndpoints sorts endpoints by dependency order:
// 1. public (no middleware, or auth/login paths)
// 2. GET
// 3. POST + DELETE
// 4. PUT / PATCH
func sortEndpoints(eps []scanner.Endpoint) []scanner.Endpoint {
	sorted := make([]scanner.Endpoint, len(eps))
	copy(sorted, eps)

	sort.SliceStable(sorted, func(i, j int) bool {
		return endpointOrder(sorted[i]) < endpointOrder(sorted[j])
	})
	return sorted
}
