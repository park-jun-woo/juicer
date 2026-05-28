//ff:func feature=scan type=test control=sequence topic=express
//ff:what lexical_declaration 라우터 패턴이 여전히 동작하는지 회귀 테스트
package express

import "testing"

func TestCollectParamRouters_DeclRegression(t *testing.T) {
	src := []byte(`
const router = express.Router();
router.get("/users", listUsers);
router.post("/users", createUser);
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi, nil)
	if !routers["router"] {
		t.Errorf("expected 'router' in routers, got %v", routers)
	}
	routes := extractRoutes(fi, routers)
	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}
}
