//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestProcessAssign_GroupAndSkips 테스트
package fiber

import "testing"

func TestProcessAssign_GroupAndSkips(t *testing.T) {
	src := `package m
func f() {
	app := fiber.New()
	api := app.Group("/api")
	x := 5            // rhs not a call -> skip
	y := compute()    // call but fun not selector -> skip
	_ = app.Group("/z") // lhs not ident -> skip
}
`
	routers := map[string]*routerInfo{}
	for _, a := range assignStmts(t, src) {
		processAssign(a, "fiber", routers)
	}
	if _, ok := routers["app"]; !ok {
		t.Fatal("app not registered")
	}
	if ri, ok := routers["api"]; !ok || ri.prefix != "/api" {
		t.Fatalf("api group not registered with prefix: %+v", routers["api"])
	}
}
