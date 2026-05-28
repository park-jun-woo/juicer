//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what 기본 라우트 추출 테스트
package hono

import "testing"

func TestCollectRoutes_Basic(t *testing.T) {
	src := []byte(`
import { Hono } from "hono"
const app = new Hono()
app.get("/users", listUsers)
app.post("/users", createUser)
app.get("/users/:id", getUser)
`)
	fi := mustParse(t, src)
	vars := collectHonoVars(fi)
	routes := collectRoutes(fi, vars)
	if len(routes) != 3 {
		t.Fatalf("expected 3 routes, got %d", len(routes))
	}
	expected := []struct {
		method string
		path   string
	}{
		{"GET", "/users"},
		{"POST", "/users"},
		{"GET", "/users/:id"},
	}
	for i, e := range expected {
		if routes[i].Method != e.method {
			t.Errorf("route %d: expected method %s, got %s", i, e.method, routes[i].Method)
		}
		if routes[i].Path != e.path {
			t.Errorf("route %d: expected path %s, got %s", i, e.path, routes[i].Path)
		}
	}
}
