//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractRolesFromNode 테스트
package spring

import "testing"

func TestExtractRolesFromNode(t *testing.T) {
	root, src := parseS(t, `class C { @PreAuthorize("hasRole('USER')") void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	roles := extractRolesFromNode(m, src)
	if len(roles) == 0 || roles[0] != "USER" {
		t.Fatalf("got %v", roles)
	}
}
