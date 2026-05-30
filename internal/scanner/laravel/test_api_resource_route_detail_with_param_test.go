//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApiResourceRoute_DetailWithParam 테스트
package laravel

import "testing"

func TestApiResourceRoute_DetailWithParam(t *testing.T) {
	a := apiResourceAction{method: "GET", suffix: "/{%s}", action: "show", hasParam: true}
	r := apiResourceRoute(a, "/users", "user", "UserController", "routes/api.php", 1, nil)
	if r.path != "/users/{user}" {
		t.Fatalf("path got %q", r.path)
	}
}
