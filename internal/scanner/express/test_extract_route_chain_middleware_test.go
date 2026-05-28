//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 체인 미들웨어 추출 테스트: .get(auth, validate, handler) → middleware ["auth", "validate"]
package express

import "testing"

func TestExtractRouteChainMiddleware(t *testing.T) {
	src := []byte(`
const router = express.Router();
router.route("/:id").get(auth, validate, getUser).put(updateUser);
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi, nil)
	routes := extractRoutes(fi, routers)
	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}

	found := map[string]routeInfo{}
	for _, r := range routes {
		found[r.Method] = r
	}

	get := found["GET"]
	if get.Handler != "getUser" {
		t.Errorf("GET handler: want getUser, got %s", get.Handler)
	}
	if len(get.Middleware) != 2 {
		t.Fatalf("GET middleware: want 2, got %d: %v", len(get.Middleware), get.Middleware)
	}
	if get.Middleware[0] != "auth" {
		t.Errorf("GET middleware[0]: want auth, got %s", get.Middleware[0])
	}
	if get.Middleware[1] != "validate" {
		t.Errorf("GET middleware[1]: want validate, got %s", get.Middleware[1])
	}

	put := found["PUT"]
	if put.Handler != "updateUser" {
		t.Errorf("PUT handler: want updateUser, got %s", put.Handler)
	}
	if len(put.Middleware) != 0 {
		t.Errorf("PUT middleware: want 0, got %d: %v", len(put.Middleware), put.Middleware)
	}
}
