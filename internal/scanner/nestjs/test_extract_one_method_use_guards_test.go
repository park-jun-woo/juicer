//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractOneMethod_UseGuards 테스트
package nestjs

import "testing"

func TestExtractOneMethod_UseGuards(t *testing.T) {
	src := []byte(`
class C {
  @UseGuards(JwtAuthGuard, RolesGuard)
  @Get('profile')
  getProfile() {}
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
	if len(ep.middleware) != 2 {
		t.Fatalf("expected 2 middleware, got %d: %v", len(ep.middleware), ep.middleware)
	}
	if ep.middleware[0] != "JwtAuthGuard" {
		t.Fatalf("expected JwtAuthGuard, got %q", ep.middleware[0])
	}
	if ep.middleware[1] != "RolesGuard" {
		t.Fatalf("expected RolesGuard, got %q", ep.middleware[1])
	}
}
