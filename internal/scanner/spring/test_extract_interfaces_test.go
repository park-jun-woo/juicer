//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractInterfaces 테스트
package spring

import "testing"

func TestExtractInterfaces(t *testing.T) {
	root, src := parseS(t, `class C implements UserApi, Other {}`)
	cls := findAllByType(root, "class_declaration")[0]
	ifaces := extractInterfaces(cls, src)
	if len(ifaces) != 2 || ifaces[0] != "UserApi" {
		t.Fatalf("got %v", ifaces)
	}
}
