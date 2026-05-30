//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApiResourceRoute_Collection 테스트
package laravel

import "testing"

func TestApiResourceRoute_Collection(t *testing.T) {
	a := apiResourceAction{method: "GET", suffix: "", action: "index", hasParam: false}
	r := apiResourceRoute(a, "/users", "user", "UserController", "routes/api.php", 10, []string{"auth"})
	if r.method != "GET" || r.path != "/users" || r.action != "index" {
		t.Fatalf("got %+v", r)
	}
	if r.controller != "UserController" || r.file != "routes/api.php" || r.line != 10 {
		t.Fatalf("meta: %+v", r)
	}
	if len(r.middleware) != 1 || r.middleware[0] != "auth" {
		t.Fatalf("mw: %+v", r.middleware)
	}
}
