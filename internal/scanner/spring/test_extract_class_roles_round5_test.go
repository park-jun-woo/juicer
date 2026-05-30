//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractClassRoles_Round5 테스트
package spring

import "testing"

func TestExtractClassRoles_Round5(t *testing.T) {
	root, src := sParse(t, `@Secured("ROLE_ADMIN") class C {}`)
	cls := sFirst(t, root, "class_declaration")
	roles := extractClassRoles(cls, src)
	if len(roles) != 1 {
		t.Fatalf("roles: %v", roles)
	}
}
