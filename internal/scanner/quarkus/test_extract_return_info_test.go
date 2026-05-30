//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractReturnInfo 테스트
package quarkus

import "testing"

func TestExtractReturnInfo(t *testing.T) {
	fi := qFileInfo(t, `class R { public List<UserDto> list() { return null; } }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractReturnInfo(m, fi.src, ep)
	if ep.returnType != "UserDto" || !ep.returnIsArray {
		t.Fatalf("got %+v", ep)
	}
}
