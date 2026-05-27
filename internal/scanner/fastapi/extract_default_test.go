//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDefault default-found 분기 테스트
package fastapi

import "testing"

func TestExtractDefault(t *testing.T) {
	src := []byte("def f(x: int = 5): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")
	param := findChildByType(params, "typed_default_parameter")
	if param == nil {
		param = findChildByType(params, "default_parameter")
	}
	if param == nil {
		t.Fatal("no default parameter")
	}
	val, call, _ := extractDefault(param, src)
	if val != "5" {
		t.Fatalf("expected '5', got %q", val)
	}
	if call != "" {
		t.Fatalf("expected empty call, got %q", call)
	}
}
