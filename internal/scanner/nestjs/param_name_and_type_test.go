//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParamNameAndType_Typed 테스트
package nestjs

import "testing"

func TestParamNameAndType_Typed(t *testing.T) {
	src := []byte(`function f(id: string) {}`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	params := findAllByType(root, "required_parameter")
	if len(params) == 0 {
		t.Fatal("no params")
	}
	name, typ := paramNameAndType(params[0], src)
	if name != "id" {
		t.Fatalf("expected id, got %q", name)
	}
	if typ != "string" {
		t.Fatalf("expected string, got %q", typ)
	}
}
