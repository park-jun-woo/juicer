//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestBuildOneEndpoint_Closure 테스트
package laravel

import "testing"

func TestBuildOneEndpoint_Closure(t *testing.T) {
	ri := routeInfo{method: "GET", path: "/health", file: "routes/api.php", line: 5, middleware: []string{"auth"}}
	ep := buildOneEndpoint(t.TempDir(), ri, map[string]*fileInfo{})
	if ep.Method != "GET" || ep.Path != "/health" || ep.Handler != "closure" {
		t.Fatalf("got %+v", ep)
	}
	if ep.Line != 5 || len(ep.Middleware) != 1 {
		t.Fatalf("meta: %+v", ep)
	}
}
