//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRouteFromCall_Chain 테스트
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
