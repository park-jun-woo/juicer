//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractPreAuthorizeRoles 테스트
package spring

import "testing"

func TestExtractPreAuthorizeRoles(t *testing.T) {
	root, src := parseS(t, `class C { @PreAuthorize("hasRole('ADMIN')") void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	roles := extractPreAuthorizeRoles(m, src)
	if len(roles) != 1 || roles[0] != "ADMIN" {
		t.Fatalf("got %v", roles)
	}
}
