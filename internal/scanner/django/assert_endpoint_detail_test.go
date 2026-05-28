//ff:func feature=scan type=test control=sequence topic=django
//ff:what E2E 테스트의 개별 엔드포인트 상세 검증 헬퍼
package django

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func verifyEndpointDetail(t *testing.T, ep scanner.Endpoint) {
	t.Helper()
	if ep.Method == "POST" && ep.Path == "/users" {
		if ep.Request == nil || ep.Request.Body == nil {
			t.Error("POST /users should have request body from serializer")
		} else if len(ep.Request.Body.Fields) == 0 {
			t.Error("POST /users request body should have fields")
		}
	}
	if ep.Method == "GET" && ep.Path == "/users/{pk}" {
		if ep.Request == nil || len(ep.Request.PathParams) == 0 {
			t.Error("GET /users/{pk} should have pk path param")
		} else if ep.Request.PathParams[0].Name != "pk" {
			t.Errorf("expected pk path param, got %s", ep.Request.PathParams[0].Name)
		}
	}
}
