//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 엔드포인트 목록에서 (file, line) → 인덱스 매핑을 구축한다
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func buildEndpointIndex(endpoints []scanner.Endpoint) map[struct{ file string; line int }]int {
	m := make(map[struct{ file string; line int }]int, len(endpoints))
	for i, ep := range endpoints {
		m[struct{ file string; line int }{ep.File, ep.Line}] = i
	}
	return m
}
