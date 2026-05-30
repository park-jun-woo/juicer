//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestBuildEndpoint 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestBuildEndpoint(t *testing.T) {
	ci := controllerInfo{prefix: "/users", roles: []string{"ADMIN"}}
	ep := endpointInfo{method: "GET", path: "/{id}", handler: "get", params: []scanner.Param{{Name: "id", Type: "string"}}, returnType: "UserDto"}
	got := buildEndpoint(ci, ep)
	if got.Path != "/users/{id}" || got.Method != "GET" {
		t.Fatalf("meta: %+v", got)
	}
	if got.Request == nil || len(got.Responses) != 1 {
		t.Fatalf("req/resp: %+v", got)
	}
}
