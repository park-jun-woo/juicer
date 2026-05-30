//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractRolesFromExpr 테스트
package spring

import "testing"

func TestExtractRolesFromExpr(t *testing.T) {
	roles := extractRolesFromExpr(`'ROLE_ADMIN', 'ROLE_USER'`)
	if len(roles) != 2 || roles[0] != "ADMIN" {
		t.Fatalf("got %v", roles)
	}
}
