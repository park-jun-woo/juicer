//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestBuildEndpoint 테스트
package dotnet

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestBuildEndpoint(t *testing.T) {
	ci := controllerInfo{prefix: "api/users", className: "UsersController", roles: []string{"Admin"}}
	ep := endpointInfo{method: "GET", path: "{id}", handler: "Get", params: []scanner.Param{{Name: "id", Type: "integer"}}, returnType: "UserDto"}
	got := buildEndpoint(ci, ep)
	if got.Method != "GET" || got.Path == "" {
		t.Fatalf("meta: %+v", got)
	}
	if got.Request == nil || len(got.Responses) != 1 {
		t.Fatalf("req/resp: %+v", got)
	}
}
