//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParamNameAndType_Untyped 테스트
package nestjs

import "testing"

func TestParamNameAndType_Untyped(t *testing.T) {
	src := []byte(`function f(x) {}`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	params := findAllByType(root, "required_parameter")
	if len(params) == 0 {
		t.Fatal("no params")
	}
	name, typ := paramNameAndType(params[0], src)
	if name != "x" {
		t.Fatalf("expected x, got %q", name)
	}
	if typ != "string" {
		t.Fatalf("expected default string, got %q", typ)
	}
}
