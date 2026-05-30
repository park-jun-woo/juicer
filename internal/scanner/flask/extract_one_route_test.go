//ff:func feature=scan type=test control=sequence topic=flask
//ff:what extractOneRoute 테스트
package flask

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstDecoratedDef(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	defs := findAllByType(root, "decorated_definition")
	if len(defs) == 0 {
		t.Fatal("no decorated_definition")
	}
	return defs[0], b
}

func TestExtractOneRoute_Route(t *testing.T) {
	src := `from flask import Flask
app = Flask(__name__)

@app.route('/users/<int:user_id>')
def get_user(user_id):
    return {}
`
	def, b := firstDecoratedDef(t, src)
	routes := extractOneRoute(def, b, make(blueprintPrefix), "app.py")
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if routes[0].handler != "get_user" || routes[0].path != "/users/{user_id}" {
		t.Fatalf("route = %+v", routes[0])
	}
}

func TestExtractOneRoute_NoFunctionDef(t *testing.T) {
	// a decorated class (no function_definition child) -> nil
	src := `@register
class MyView:
    pass
`
	def, b := firstDecoratedDef(t, src)
	if routes := extractOneRoute(def, b, make(blueprintPrefix), "app.py"); routes != nil {
		t.Fatalf("decorated class should yield nil, got %v", routes)
	}
}

func TestExtractOneRoute_NoDecorators(t *testing.T) {
	// pass a plain function_definition node (no decorators) -> nil
	b := []byte("def plain():\n    pass\n")
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	fns := findAllByType(root, "function_definition")
	if len(fns) == 0 {
		t.Fatal("no function_definition")
	}
	if routes := extractOneRoute(fns[0], b, make(blueprintPrefix), "app.py"); routes != nil {
		t.Fatalf("non-decorated def should yield nil, got %v", routes)
	}
}

func TestExtractOneRoute_NonRouteDecorator(t *testing.T) {
	src := `@staticmethod
def helper():
    pass
`
	def, b := firstDecoratedDef(t, src)
	if routes := extractOneRoute(def, b, make(blueprintPrefix), "app.py"); routes != nil {
		t.Fatalf("non-route decorator should yield nil, got %v", routes)
	}
}
