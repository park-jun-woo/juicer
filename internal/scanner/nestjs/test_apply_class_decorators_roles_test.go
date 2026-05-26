//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestApplyClassDecorators_Roles 테스트
package nestjs

import "testing"

func TestApplyClassDecorators_Roles(t *testing.T) {
	ci := controllerInfo{
		classRoles: []string{"Role.admin"},
		endpoints: []endpointInfo{
			{handler: "findAll", roles: nil},
			{handler: "create", roles: []string{"Role.premium"}},
		},
	}
	applyClassDecorators(&ci)
	if len(ci.endpoints[0].roles) != 1 || ci.endpoints[0].roles[0] != "Role.admin" {
		t.Fatalf("findAll: expected [Role.admin], got %v", ci.endpoints[0].roles)
	}
	if len(ci.endpoints[1].roles) != 2 {
		t.Fatalf("create: expected 2 roles, got %v", ci.endpoints[1].roles)
	}
	if ci.endpoints[1].roles[0] != "Role.admin" {
		t.Fatalf("create: expected first=Role.admin, got %q", ci.endpoints[1].roles[0])
	}
	if ci.endpoints[1].roles[1] != "Role.premium" {
		t.Fatalf("create: expected second=Role.premium, got %q", ci.endpoints[1].roles[1])
	}
}
