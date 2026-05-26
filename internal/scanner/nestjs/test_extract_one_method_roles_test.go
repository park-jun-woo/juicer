//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractOneMethod_Roles 테스트
package nestjs

import "testing"

func TestExtractOneMethod_Roles(t *testing.T) {
	src := []byte(`
class C {
  @UseGuards(JwtAuthGuard, RolesGuard)
  @Roles(Role.premium)
  @Get('premium')
  getPremium() {}
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	methods := findAllByType(root, "method_definition")
	if len(methods) == 0 {
		t.Fatal("no methods found")
	}
	ep, ok := extractOneMethod(methods[0], src, "test.ts")
	if !ok {
		t.Fatal("expected ok")
	}
	if len(ep.roles) != 1 || ep.roles[0] != "Role.premium" {
		t.Fatalf("expected roles=[Role.premium], got %v", ep.roles)
	}
	if len(ep.middleware) != 2 {
		t.Fatalf("expected 2 middleware, got %v", ep.middleware)
	}
}
