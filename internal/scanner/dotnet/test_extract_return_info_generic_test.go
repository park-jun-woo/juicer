//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractReturnInfo_Generic 테스트
package dotnet

import "testing"

func TestExtractReturnInfo_Generic(t *testing.T) {
	root, src := parseCS(t, `class C { public ActionResult<UserDto> Get() { return Ok(); } }`)
	m := findAllByType(root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractReturnInfo(m, src, ep)
	if ep.returnType != "UserDto" {
		t.Fatalf("got %q", ep.returnType)
	}
}
