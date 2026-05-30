//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestBuildEndpoint_PathParams 테스트
package flask

import "testing"

func TestBuildEndpoint_PathParams(t *testing.T) {
	ri := routeInfo{
		method:  "GET",
		path:    "/users/{id}",
		handler: "get_user",
		file:    "app.py",
		line:    5,
		params:  []urlParam{{name: "id", converter: "int"}},
	}
	ep := buildEndpoint(ri)
	if ep.Method != "GET" || ep.Path != "/users/{id}" {
		t.Fatalf("ep = %+v", ep)
	}
	if ep.Request == nil || len(ep.Request.PathParams) != 1 || ep.Request.PathParams[0].Name != "id" {
		t.Fatalf("path params = %+v", ep.Request)
	}
}
