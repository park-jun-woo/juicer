//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractRoleStrings_Round5 테스트
package spring

import "testing"

func TestExtractRoleStrings_Round5(t *testing.T) {
	root, src := sParse(t, `@Secured({"ROLE_ADMIN", "ROLE_USER"}) class C {}`)
	ann := sFirst(t, root, "annotation")
	args := annotationArgs(ann, src)
	roles := extractRoleStrings(args, src)
	if len(roles) != 2 {
		t.Fatalf("roles: %v", roles)
	}
}
