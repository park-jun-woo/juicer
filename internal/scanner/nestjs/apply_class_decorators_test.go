//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestApplyClassDecorators_Middleware 테스트
package nestjs

import "testing"

func TestApplyClassDecorators_Middleware(t *testing.T) {
	ci := controllerInfo{
		classMiddleware: []string{"JwtAuthGuard"},
		endpoints: []endpointInfo{
			{handler: "findAll", middleware: nil},
			{handler: "create", middleware: []string{"RolesGuard"}},
		},
	}
	applyClassDecorators(&ci)
	if len(ci.endpoints[0].middleware) != 1 || ci.endpoints[0].middleware[0] != "JwtAuthGuard" {
		t.Fatalf("findAll: expected [JwtAuthGuard], got %v", ci.endpoints[0].middleware)
	}
	if len(ci.endpoints[1].middleware) != 2 {
		t.Fatalf("create: expected 2 middleware, got %v", ci.endpoints[1].middleware)
	}
	if ci.endpoints[1].middleware[0] != "JwtAuthGuard" {
		t.Fatalf("create: expected first=JwtAuthGuard, got %q", ci.endpoints[1].middleware[0])
	}
	if ci.endpoints[1].middleware[1] != "RolesGuard" {
		t.Fatalf("create: expected second=RolesGuard, got %q", ci.endpoints[1].middleware[1])
	}
}
