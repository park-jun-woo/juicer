//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestChainMethodToRoute_Round5 테스트
package express

import "testing"

func TestChainMethodToRoute_Round5(t *testing.T) {
	cm := chainMethod{method: "POST", handler: "h", middleware: []string{"auth"}, line: 3, authLevel: "user", roles: []string{"admin"}}
	r := chainMethodToRoute(cm, "/items", "router")
	if r.Method != "POST" || r.Path != "/items" || r.Router != "router" || r.Handler != "h" {
		t.Fatalf("route: %+v", r)
	}
	if len(r.Middleware) != 1 || r.AuthLevel != "user" || len(r.Roles) != 1 {
		t.Fatalf("route meta: %+v", r)
	}
}
