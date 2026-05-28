//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what 기본 라우트 추출 테스트: fastify.get("/users", handler)
package fastify

import "testing"

func TestExtractRoutes_Basic(t *testing.T) {
	src := []byte(`
import Fastify from "fastify";
const app = Fastify();
app.get("/users", listUsers);
app.post("/users", createUser);
app.get("/users/:id", getUser);
`)
	fi := mustParse(t, src)
	instances := collectInstances(fi)
	routes := extractRoutes(fi, instances)
	if len(routes) != 3 {
		t.Fatalf("expected 3 routes, got %d", len(routes))
	}
	expected := []struct {
		method  string
		path    string
		handler string
	}{
		{"GET", "/users", "listUsers"},
		{"POST", "/users", "createUser"},
		{"GET", "/users/:id", "getUser"},
	}
	for i, e := range expected {
		if routes[i].Method != e.method {
			t.Errorf("route[%d].Method: want %s, got %s", i, e.method, routes[i].Method)
		}
		if routes[i].Path != e.path {
			t.Errorf("route[%d].Path: want %s, got %s", i, e.path, routes[i].Path)
		}
		if routes[i].Handler != e.handler {
			t.Errorf("route[%d].Handler: want %s, got %s", i, e.handler, routes[i].Handler)
		}
	}
}
