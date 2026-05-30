//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRouteChain_Valid 테스트
package express

import "testing"

func TestExtractRouteChain_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/:id').get(getH).put(putH);`))
	call := outermostCall(fi)
	routes := extractRouteChain(call, fi.Src, map[string]bool{"router": true})
	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %+v", routes)
	}
	if routes[0].Path != "/:id" || routes[1].Path != "/:id" {
		t.Fatalf("unexpected paths %+v", routes)
	}
}
