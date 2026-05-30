//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractRouteFromCall: 체인 / 단일라우트 / nil 분기
package express

import "testing"

func TestExtractRouteFromCall_Chain(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/:id').get(getH);`))
	call := outermostCall(fi)
	routes := extractRouteFromCall(call, fi.Src, map[string]bool{"router": true}, map[uintptr]bool{})
	if len(routes) != 1 || routes[0].Path != "/:id" {
		t.Fatalf("got %+v", routes)
	}
}

func TestExtractRouteFromCall_Single(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', h);`))
	routes := extractRouteFromCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}, map[uintptr]bool{})
	if len(routes) != 1 || routes[0].Method != "GET" || routes[0].Path != "/x" {
		t.Fatalf("got %+v", routes)
	}
}

func TestExtractRouteFromCall_Nil(t *testing.T) {
	fi := mustParse(t, []byte(`foo('/x');`))
	if routes := extractRouteFromCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}, map[uintptr]bool{}); routes != nil {
		t.Fatalf("expected nil, got %+v", routes)
	}
}
