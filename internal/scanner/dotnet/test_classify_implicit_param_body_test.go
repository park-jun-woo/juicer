//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestClassifyImplicitParam_Body 테스트
package dotnet

import "testing"

func TestClassifyImplicitParam_Body(t *testing.T) {
	ep := &endpointInfo{method: "POST", path: "/users"}
	classifyImplicitParam("CreateUserDto", "dto", ep)
	if ep.bodyType != "CreateUserDto" {
		t.Fatalf("got %q", ep.bodyType)
	}
}
