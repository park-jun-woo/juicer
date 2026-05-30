//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParseFlaskDecorator_ShortcutGet 테스트
package flask

import "testing"

func TestParseFlaskDecorator_ShortcutGet(t *testing.T) {

	dec, b := firstDecorator(t, "@app.get('/x')\ndef h():\n    pass\n")
	routes := parseFlaskDecorator(dec, b, make(blueprintPrefix), "h", "app.py", 1)
	if len(routes) != 1 || routes[0].method != "GET" {
		t.Fatalf("shortcut get: %v", routes)
	}
}
