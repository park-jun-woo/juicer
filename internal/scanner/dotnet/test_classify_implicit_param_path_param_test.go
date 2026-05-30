//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestClassifyImplicitParam_PathParam 테스트
package dotnet

import "testing"

func TestClassifyImplicitParam_PathParam(t *testing.T) {
	ep := &endpointInfo{method: "GET", path: "/users/{id}"}
	classifyImplicitParam("int", "id", ep)
	if len(ep.params) != 1 || ep.params[0].Name != "id" {
		t.Fatalf("got %+v", ep.params)
	}
}
