//ff:func feature=scan type=test control=sequence topic=flask
//ff:what methods=['GET', 'POST'] 인자로 복수 메서드 라우트를 추출한다
package flask

import (
	"testing"
)

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

func TestExtractMethodsArg_Direct(t *testing.T) {
	args, src := argListOf(t, `route('/x', methods=['GET', 'POST', 'PUT'])`+"\n")
	got := extractMethodsArg(args, src)
	if len(got) != 3 || got[0] != "GET" || got[2] != "PUT" {
		t.Fatalf("methods = %v", got)
	}
}

func TestExtractMethodsArg_NoMethods(t *testing.T) {
	args, src := argListOf(t, `route('/x')`+"\n")
	if got := extractMethodsArg(args, src); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestExtractMethodsArg_OtherKwargBeforeMethods(t *testing.T) {
	// a non-methods kwarg precedes methods -> exercises the key-mismatch continue
	args, src := argListOf(t, `route('/x', strict_slashes=False, methods=['GET'])`+"\n")
	got := extractMethodsArg(args, src)
	if len(got) != 1 || got[0] != "GET" {
		t.Fatalf("methods = %v", got)
	}
}

func TestExtractMethodsArg_NotList(t *testing.T) {
	// methods present but value is not a list literal -> nil
	args, src := argListOf(t, `route('/x', methods=ALLOWED)`+"\n")
	if got := extractMethodsArg(args, src); got != nil {
		t.Fatalf("expected nil for non-list methods, got %v", got)
	}
}
