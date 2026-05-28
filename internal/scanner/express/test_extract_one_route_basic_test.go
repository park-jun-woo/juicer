//ff:func feature=scan type=test control=sequence topic=express
//ff:what 기본 라우트 추출 테스트: app.get("/users", handler) → GET /users
package express

import "testing"

func TestExtractOneRoute_Basic(t *testing.T) {
	src := []byte(`
const app = express();
app.get("/users", listUsers);
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi)
	routes := extractRoutes(fi, routers)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	r := routes[0]
	if r.Method != "GET" {
		t.Errorf("method: want GET, got %s", r.Method)
	}
	if r.Path != "/users" {
		t.Errorf("path: want /users, got %s", r.Path)
	}
	if r.Handler != "listUsers" {
		t.Errorf("handler: want listUsers, got %s", r.Handler)
	}
}
