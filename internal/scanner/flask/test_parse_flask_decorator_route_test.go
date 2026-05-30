//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParseFlaskDecorator_Route 테스트
package flask

import "testing"

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
