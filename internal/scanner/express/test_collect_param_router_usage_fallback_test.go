//ff:func feature=scan type=test control=sequence topic=express
//ff:what 타입 어노테이션 없는 함수 파라미터에서 호출 패턴으로 라우터를 추정한다
package express

import "testing"

func TestCollectParamRouters_UsageFallback(t *testing.T) {
	src := []byte(`
export default (router) => {
  router.get("/items", listItems);
  router.delete("/items/:id", deleteItem);
};
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi)
	if !routers["router"] {
		t.Errorf("expected 'router' in routers via usage fallback, got %v", routers)
	}
	routes := extractRoutes(fi, routers)
	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}
}
