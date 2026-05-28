//ff:func feature=scan type=test control=sequence topic=express
//ff:what 타입 어노테이션 파라미터 라우터 인식 테스트: (router: express.Router) → router 감지
package express

import "testing"

func TestCollectParamRouters_Typed(t *testing.T) {
	src := []byte(`
export default (router: express.Router) => {
  router.get("/", handler);
  router.post("/", createHandler);
};
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
	if routes[0].Method != "GET" || routes[0].Path != "/" {
		t.Errorf("route[0]: want GET /, got %s %s", routes[0].Method, routes[0].Path)
	}
	if routes[1].Method != "POST" || routes[1].Path != "/" {
		t.Errorf("route[1]: want POST /, got %s %s", routes[1].Method, routes[1].Path)
	}
}
