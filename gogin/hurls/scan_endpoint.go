//ff:func feature=hurl type=parse control=sequence
//ff:what scanner.Scan 실행 후 ID에 매칭되는 엔드포인트 반환
package hurls

import (
	"github.com/park-jun-woo/juicer/scanner"
)

// scanEndpoint performs a scan and returns the matching endpoint data.
func scanEndpoint(repoDir, id string) *scanner.Endpoint {
	result, err := scanner.Scan(repoDir)
	if err != nil {
		return nil
	}
	return findEndpoint(result.Endpoints, id)
}
