//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParseFlaskDecorator_NonRouteMethod 테스트
package flask

import "testing"

func TestParseFlaskDecorator_NonRouteMethod(t *testing.T) {

	dec, b := firstDecorator(t, "@app.before_request('x')\ndef h():\n    pass\n")
	if routes := parseFlaskDecorator(dec, b, make(blueprintPrefix), "h", "app.py", 1); routes != nil {
		t.Fatalf("non-route method should yield nil, got %v", routes)
	}
}
