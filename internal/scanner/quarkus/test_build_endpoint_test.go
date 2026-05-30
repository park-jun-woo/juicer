//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestBuildEndpoint 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestBuildEndpoint(t *testing.T) {
	ri := resourceInfo{prefix: "/users", roles: []string{"admin"}}
	ep := endpointInfo{method: "GET", path: "/{id}", handler: "get", params: []scanner.Param{{Name: "id", Type: "string"}}, returnType: "UserDto"}
	got := buildEndpoint(ri, ep)
	if got.Path != "/users/{id}" || got.Method != "GET" {
		t.Fatalf("meta: %+v", got)
	}
	if len(got.Roles) != 1 || got.Roles[0] != "admin" {
		t.Fatalf("roles: %v", got.Roles)
	}
	if got.Request == nil || len(got.Responses) != 1 {
		t.Fatalf("req/resp: %+v", got)
	}
}
