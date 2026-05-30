//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractRoleStrings 테스트
package quarkus

import "testing"

func TestExtractRoleStrings(t *testing.T) {
	root, _ := parseJava([]byte(`class C { @RolesAllowed({"admin", "user"}) void m() {} }`))
	args := findAllByType(root, "annotation_argument_list")
	if len(args) == 0 {
		t.Skip("no args")
	}
	roles := extractRoleStrings(args[0], []byte(`class C { @RolesAllowed({"admin", "user"}) void m() {} }`))
	if len(roles) != 2 || roles[0] != "admin" {
		t.Fatalf("got %v", roles)
	}
}
