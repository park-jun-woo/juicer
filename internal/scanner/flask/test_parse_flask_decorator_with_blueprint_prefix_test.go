//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParseFlaskDecorator_WithBlueprintPrefix 테스트
package flask

import "testing"

func TestParseFlaskDecorator_WithBlueprintPrefix(t *testing.T) {
	dec, b := firstDecorator(t, "@api.route('/items')\ndef h():\n    pass\n")
	prefixes := blueprintPrefix{"api": "/v1"}
	routes := parseFlaskDecorator(dec, b, prefixes, "h", "app.py", 1)
	if len(routes) != 1 || routes[0].path != "/v1/items" {
		t.Fatalf("blueprint prefix not applied: %v", routes)
	}
}
