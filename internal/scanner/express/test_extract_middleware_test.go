//ff:func feature=scan type=test control=sequence topic=express
//ff:what 미들웨어 추출 테스트: app.get("/path", auth, handler) → middleware ["auth"]
package express

import "testing"

func TestExtractMiddleware(t *testing.T) {
	src := []byte(`
const app = express();
app.get("/admin", authMiddleware, validateRole, adminHandler);
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi)
	routes := extractRoutes(fi, routers)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	r := routes[0]
	if len(r.Middleware) != 2 {
		t.Fatalf("expected 2 middleware, got %d: %v", len(r.Middleware), r.Middleware)
	}
	if r.Middleware[0] != "authMiddleware" {
		t.Errorf("middleware[0]: want authMiddleware, got %s", r.Middleware[0])
	}
	if r.Middleware[1] != "validateRole" {
		t.Errorf("middleware[1]: want validateRole, got %s", r.Middleware[1])
	}
	if r.Handler != "adminHandler" {
		t.Errorf("handler: want adminHandler, got %s", r.Handler)
	}
}
