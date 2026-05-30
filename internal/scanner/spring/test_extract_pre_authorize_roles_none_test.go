//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractPreAuthorizeRoles_None 테스트
package spring

import "testing"

func TestExtractPreAuthorizeRoles_None(t *testing.T) {
	root, src := parseS(t, `class C { void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	if roles := extractPreAuthorizeRoles(m, src); roles != nil {
		t.Fatalf("got %v", roles)
	}
}
