//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDefault_NoDefault default 없는 분기 테스트
package fastapi

import "testing"

func TestExtractDefault_NoDefault(t *testing.T) {
	// param without default value (typed_parameter)
	src := []byte("def f(x: int): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")
	// typed_parameter has no default
	param := findChildByType(params, "typed_parameter")
	if param == nil {
		t.Skip("no typed_parameter found")
	}
	val, call := extractDefault(param, src)
	if val != "" {
		t.Fatalf("expected empty val, got %q", val)
	}
	if call != "" {
		t.Fatalf("expected empty call, got %q", call)
	}
}
