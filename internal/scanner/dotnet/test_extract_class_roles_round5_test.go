//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractClassRoles_Round5 테스트
package dotnet

import "testing"

func TestExtractClassRoles_Round5(t *testing.T) {
	root, src := parseCS(t, `[Authorize(Roles = "Admin,User")] class C {}`)
	cls := firstOfType(t, root, "class_declaration")
	roles := extractClassRoles(cls, src)
	if len(roles) != 2 {
		t.Fatalf("expected 2 roles, got %v", roles)
	}
}
