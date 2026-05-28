//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 체인 패턴 테스트: router.route("/:id").get(h).put(h) → GET /:id, PUT /:id
package express

import "testing"

func TestExtractRouteChain(t *testing.T) {
	src := []byte(`
const router = express.Router();
router.route("/:id").get(getUser).put(updateUser);
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi)
	routes := extractRoutes(fi, routers)
	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}

	found := map[string]string{}
	for _, r := range routes {
		found[r.Method] = r.Handler
	}
	if found["GET"] != "getUser" {
		t.Errorf("GET handler: want getUser, got %s", found["GET"])
	}
	if found["PUT"] != "updateUser" {
		t.Errorf("PUT handler: want updateUser, got %s", found["PUT"])
	}
	for _, r := range routes {
		if r.Path != "/:id" {
			t.Errorf("path: want /:id, got %s", r.Path)
		}
	}
}
