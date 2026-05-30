//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractRolesFromNode_Round5 테스트
package quarkus

import "testing"

func TestExtractRolesFromNode_Round5(t *testing.T) {
	root, src := qParse(t, `@RolesAllowed({"admin", "user"}) class C {}`)
	cls := qFirst(t, root, "class_declaration")
	roles := extractRolesFromNode(cls, src)
	if len(roles) != 2 {
		t.Fatalf("expected 2 roles, got %v", roles)
	}
}
