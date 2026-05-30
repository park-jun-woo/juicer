//ff:func feature=scan type=test control=sequence topic=flask
//ff:what URL 변수 <int:user_id>에서 path parameter를 추출한다
package flask

import (
	"testing"
)

func TestExtractURLParams_IntParam(t *testing.T) {
	src := []byte(`from flask import Flask

app = Flask(__name__)

@app.route('/users/<int:user_id>')
def get_user(user_id):
    pass
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
	if r.path != "/users/{user_id}" {
		t.Errorf("expected /users/{user_id}, got %s", r.path)
	}
	if len(r.params) != 1 {
		t.Fatalf("expected 1 param, got %d", len(r.params))
	}
	if r.params[0].name != "user_id" {
		t.Errorf("expected user_id, got %s", r.params[0].name)
	}
	if r.params[0].converter != "int" {
		t.Errorf("expected int converter, got %s", r.params[0].converter)
	}
}

func TestExtractURLParams_Direct(t *testing.T) {
	params := extractURLParams("/users/<int:user_id>/posts/<post_id>/<path:rest>")
	if len(params) != 3 {
		t.Fatalf("expected 3 params, got %d: %v", len(params), params)
	}
	if params[0].name != "user_id" || params[0].converter != "int" {
		t.Errorf("param 0 = %+v", params[0])
	}
	if params[1].name != "post_id" || params[1].converter != "" {
		t.Errorf("param 1 = %+v", params[1])
	}
	if params[2].name != "rest" || params[2].converter != "path" {
		t.Errorf("param 2 = %+v", params[2])
	}
}

func TestExtractURLParams_None(t *testing.T) {
	if got := extractURLParams("/static/path"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}
