//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRouteFromCall_Single 테스트
package express

import "testing"

func TestExtractRouteFromCall_Single(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', h);`))
	routes := extractRouteFromCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}, map[uintptr]bool{})
	if len(routes) != 1 || routes[0].Method != "GET" || routes[0].Path != "/x" {
		t.Fatalf("got %+v", routes)
	}
}
