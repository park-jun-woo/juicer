//ff:func feature=scan type=test control=sequence topic=flask
//ff:what 기본 @app.route 데코레이터에서 GET 라우트를 추출한다
package flask

import (
	"testing"
)

func TestExtractOneRoute_BasicGet(t *testing.T) {
	src := []byte(`from flask import Flask

app = Flask(__name__)

@app.route('/users')
def list_users():
    return []
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}

	bpPrefixes := make(blueprintPrefix)
	routes := extractRoutes(root, src, bpPrefixes, "app.py")

	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	r := routes[0]
	if r.method != "GET" {
		t.Errorf("expected GET, got %s", r.method)
	}
	if r.path != "/users" {
		t.Errorf("expected /users, got %s", r.path)
	}
	if r.handler != "list_users" {
		t.Errorf("expected list_users, got %s", r.handler)
	}
}
