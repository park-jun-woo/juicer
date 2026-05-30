//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParseFlaskDecorator_NoAttribute 테스트
package flask

import "testing"

func TestParseFlaskDecorator_NoAttribute(t *testing.T) {

	dec, b := firstDecorator(t, "@decorator('x')\ndef h():\n    pass\n")
	if routes := parseFlaskDecorator(dec, b, make(blueprintPrefix), "h", "app.py", 1); routes != nil {
		t.Fatalf("non-attribute call should yield nil, got %v", routes)
	}
}
