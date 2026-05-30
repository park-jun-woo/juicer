//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractRolesAllowed 테스트
package spring

import "testing"

func TestExtractRolesAllowed(t *testing.T) {
	root, src := parseS(t, `@RolesAllowed({"admin"}) class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	roles := extractRolesAllowed(cls, src)
	if len(roles) != 1 || roles[0] != "admin" {
		t.Fatalf("got %v", roles)
	}
}
