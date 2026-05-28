//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what Flask 2.0+ 숏컷 메서드 데코레이터 (@app.get, @app.post 등)를 추출한다
package flask

import (
	"testing"
)

func TestShortcutMethods(t *testing.T) {
	src := []byte(`from flask import Flask

app = Flask(__name__)

@app.get('/items')
def list_items():
    return []

@app.post('/items')
def create_item():
    return {}, 201

@app.put('/items/<int:id>')
def update_item(id):
    return {}

@app.delete('/items/<int:id>')
def delete_item(id):
    return '', 204
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}

	bpPrefixes := make(blueprintPrefix)
	routes := extractRoutes(root, src, bpPrefixes, "app.py")

	if len(routes) != 4 {
		t.Fatalf("expected 4 routes, got %d", len(routes))
	}

	expected := []struct {
		method  string
		path    string
		handler string
	}{
		{"GET", "/items", "list_items"},
		{"POST", "/items", "create_item"},
		{"PUT", "/items/{id}", "update_item"},
		{"DELETE", "/items/{id}", "delete_item"},
	}

	for i, exp := range expected {
		if routes[i].method != exp.method {
			t.Errorf("route %d: expected method %s, got %s", i, exp.method, routes[i].method)
		}
		if routes[i].path != exp.path {
			t.Errorf("route %d: expected path %s, got %s", i, exp.path, routes[i].path)
		}
		if routes[i].handler != exp.handler {
			t.Errorf("route %d: expected handler %s, got %s", i, exp.handler, routes[i].handler)
		}
	}
}
