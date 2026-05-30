//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractReturnInfo_Array 테스트
package dotnet

import "testing"

func TestExtractReturnInfo_Array(t *testing.T) {
	root, src := parseCS(t, `class C { public ActionResult<List<UserDto>> List() { return Ok(); } }`)
	m := findAllByType(root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractReturnInfo(m, src, ep)
	if ep.returnType != "UserDto" || !ep.returnIsArray {
		t.Fatalf("got %+v", ep)
	}
}
