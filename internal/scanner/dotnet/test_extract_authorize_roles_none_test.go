//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractAuthorizeRoles_None 테스트
package dotnet

import "testing"

func TestExtractAuthorizeRoles_None(t *testing.T) {
	root, src := parseCS(t, `class C { void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	if roles := extractAuthorizeRoles(m, src); roles != nil {
		t.Fatalf("got %v", roles)
	}
}
