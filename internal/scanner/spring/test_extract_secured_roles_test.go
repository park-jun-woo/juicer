//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractSecuredRoles 테스트
package spring

import "testing"

func TestExtractSecuredRoles(t *testing.T) {
	root, src := parseS(t, `class C { @Secured({"ROLE_ADMIN"}) void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	roles := extractSecuredRoles(m, src)
	if len(roles) != 1 || roles[0] != "ADMIN" {
		t.Fatalf("got %v", roles)
	}
}
