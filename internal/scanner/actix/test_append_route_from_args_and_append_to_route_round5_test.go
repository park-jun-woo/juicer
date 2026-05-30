//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestAppendRouteFromArgs_And_AppendToRoute_Round5 테스트
package actix

import "testing"

func TestAppendRouteFromArgs_And_AppendToRoute_Round5(t *testing.T) {
	fi := aFi(t, `fn f() { web::get().to(handler); }`)
	args := aFirst(t, fi.root, "arguments")
	var routes []builderRoute
	appendRouteFromArgs(args, fi.src, "/x", &routes)
	var routes2 []builderRoute
	appendToRoute(args, fi.src, "/x", &routes2)
}
