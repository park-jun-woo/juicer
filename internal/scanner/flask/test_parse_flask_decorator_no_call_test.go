//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParseFlaskDecorator_NoCall 테스트
package flask

import "testing"

func TestParseFlaskDecorator_NoCall(t *testing.T) {

	dec, b := firstDecorator(t, "@staticmethod\ndef h():\n    pass\n")
	if routes := parseFlaskDecorator(dec, b, make(blueprintPrefix), "h", "app.py", 1); routes != nil {
		t.Fatalf("bare decorator should yield nil, got %v", routes)
	}
}
