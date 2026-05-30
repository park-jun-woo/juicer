//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractAuthorizeRoles 테스트
package dotnet

import "testing"

func TestExtractAuthorizeRoles(t *testing.T) {
	root, src := parseCS(t, `class C { [Authorize(Roles = "Admin,User")] void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	roles := extractAuthorizeRoles(m, src)
	if len(roles) != 2 || roles[0] != "Admin" {
		t.Fatalf("got %v", roles)
	}
}
