//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestBuildRouteFromArgs_Round5 테스트
package express

import "testing"

func TestBuildRouteFromArgs_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', handler);`))
	call := firstCallExpr(t, fi)
	args := findChildByType(call, "arguments")
	r := buildRouteFromArgs(args, fi.Src, "GET", 1)
	if r == nil || r.Method != "GET" || r.Path != "/x" {
		t.Fatalf("route: %+v", r)
	}
}
