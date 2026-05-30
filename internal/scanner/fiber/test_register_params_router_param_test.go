//ff:func feature=scan type=test control=sequence
//ff:what TestRegisterParams_RouterParam 테스트
package fiber

import "testing"

func TestRegisterParams_RouterParam(t *testing.T) {
	fn := funcDecl(t, "package m\nfunc Setup(app *fiber.App, other int) {}\n")
	routers := map[string]*routerInfo{}
	registerParams(fn, "fiber", routers)
	if _, ok := routers["app"]; !ok {
		t.Fatalf("expected app registered, got %v", routers)
	}
	if _, ok := routers["other"]; ok {
		t.Fatalf("non-router param should not be registered: %v", routers)
	}
}
