//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what E2E POST 엔드포인트 body/response 검증 헬퍼
package fastify

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func verifyPostEndpoint(t *testing.T, result *scanner.ScanResult) {
	t.Helper()
	for _, ep := range result.Endpoints {
		if ep.Method != "POST" || ep.Path != "/api/users" {
			continue
		}
		if ep.Request == nil || ep.Request.Body == nil {
			t.Error("POST /api/users: expected request body from schema")
			return
		}
		if len(ep.Request.Body.Fields) != 2 {
			t.Errorf("POST /api/users: expected 2 body fields, got %d", len(ep.Request.Body.Fields))
		}
		if len(ep.Responses) == 0 {
			t.Error("POST /api/users: expected response from schema")
		}
		return
	}
	t.Error("POST /api/users endpoint not found")
}
