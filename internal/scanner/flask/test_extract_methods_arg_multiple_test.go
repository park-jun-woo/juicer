//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractMethodsArg_Multiple 테스트
package flask

import "testing"

func TestExtractMethodsArg_Multiple(t *testing.T) {
	src := []byte(`from flask import Flask

app = Flask(__name__)

@app.route('/users', methods=['GET', 'POST'])
def users():
    pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}

	bpPrefixes := make(blueprintPrefix)
	routes := extractRoutes(root, src, bpPrefixes, "app.py")

	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}
	if routes[0].method != "GET" {
		t.Errorf("route 0: expected GET, got %s", routes[0].method)
	}
	if routes[1].method != "POST" {
		t.Errorf("route 1: expected POST, got %s", routes[1].method)
	}
	if routes[0].path != "/users" || routes[1].path != "/users" {
		t.Errorf("expected /users for both routes, got %s and %s", routes[0].path, routes[1].path)
	}
}
