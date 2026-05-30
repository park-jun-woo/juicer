//ff:func feature=scan type=test control=sequence topic=flask
//ff:what parseFlaskDecorator 테스트
package flask

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstDecorator(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	decs := findAllByType(root, "decorator")
	if len(decs) == 0 {
		t.Fatal("no decorator")
	}
	return decs[0], b
}

func TestParseFlaskDecorator_Route(t *testing.T) {
	dec, b := firstDecorator(t, "@app.route('/users/<int:id>', methods=['GET','POST'])\ndef h():\n    pass\n")
	routes := parseFlaskDecorator(dec, b, make(blueprintPrefix), "h", "app.py", 1)
	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}
	if routes[0].path != "/users/{id}" {
		t.Errorf("path = %q", routes[0].path)
	}
	if len(routes[0].params) != 1 {
		t.Errorf("params = %v", routes[0].params)
	}
}

func TestParseFlaskDecorator_NoCall(t *testing.T) {
	// a bare decorator without a call (e.g. @staticmethod) -> nil
	dec, b := firstDecorator(t, "@staticmethod\ndef h():\n    pass\n")
	if routes := parseFlaskDecorator(dec, b, make(blueprintPrefix), "h", "app.py", 1); routes != nil {
		t.Fatalf("bare decorator should yield nil, got %v", routes)
	}
}

func TestParseFlaskDecorator_NoAttribute(t *testing.T) {
	// decorator with a call but no attribute (plain function call) -> nil
	dec, b := firstDecorator(t, "@decorator('x')\ndef h():\n    pass\n")
	if routes := parseFlaskDecorator(dec, b, make(blueprintPrefix), "h", "app.py", 1); routes != nil {
		t.Fatalf("non-attribute call should yield nil, got %v", routes)
	}
}

func TestParseFlaskDecorator_ShortcutGet(t *testing.T) {
	// @app.get('/x') shortcut -> single GET method route
	dec, b := firstDecorator(t, "@app.get('/x')\ndef h():\n    pass\n")
	routes := parseFlaskDecorator(dec, b, make(blueprintPrefix), "h", "app.py", 1)
	if len(routes) != 1 || routes[0].method != "GET" {
		t.Fatalf("shortcut get: %v", routes)
	}
}

func TestParseFlaskDecorator_NonRouteMethod(t *testing.T) {
	// @app.before_request(...) is an attribute call but not a route method ->
	// resolveHTTPMethods returns empty -> nil.
	dec, b := firstDecorator(t, "@app.before_request('x')\ndef h():\n    pass\n")
	if routes := parseFlaskDecorator(dec, b, make(blueprintPrefix), "h", "app.py", 1); routes != nil {
		t.Fatalf("non-route method should yield nil, got %v", routes)
	}
}

func TestParseFlaskDecorator_WithBlueprintPrefix(t *testing.T) {
	dec, b := firstDecorator(t, "@api.route('/items')\ndef h():\n    pass\n")
	prefixes := blueprintPrefix{"api": "/v1"}
	routes := parseFlaskDecorator(dec, b, prefixes, "h", "app.py", 1)
	if len(routes) != 1 || routes[0].path != "/v1/items" {
		t.Fatalf("blueprint prefix not applied: %v", routes)
	}
}
