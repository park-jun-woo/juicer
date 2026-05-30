//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractURLParams_IntParam 테스트
package flask

import "testing"

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
