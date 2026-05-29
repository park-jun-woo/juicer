//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what logEndpoints — 엔드포인트 목록을 테스트 로그로 출력한다
package laravel

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func logEndpoints(t *testing.T, endpoints []scanner.Endpoint) {
	t.Helper()
	for i, ep := range endpoints {
		t.Logf("  endpoint %d: %s %s (%s)", i, ep.Method, ep.Path, ep.Handler)
	}
}
