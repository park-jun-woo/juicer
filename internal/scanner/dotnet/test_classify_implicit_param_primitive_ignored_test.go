//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestClassifyImplicitParam_PrimitiveIgnored 테스트
package dotnet

import "testing"

func TestClassifyImplicitParam_PrimitiveIgnored(t *testing.T) {
	ep := &endpointInfo{method: "POST", path: "/users"}
	classifyImplicitParam("string", "name", ep)
	if ep.bodyType != "" {
		t.Fatalf("primitive should not be body: %q", ep.bodyType)
	}
}
