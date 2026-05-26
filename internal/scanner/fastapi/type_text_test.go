//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what typeText 테스트
package fastapi

import "testing"

func TestTypeText(t *testing.T) {
	src := []byte("def f(x: int): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	if funcDef == nil {
		t.Fatal("no function_definition")
	}
	params := findChildByType(funcDef, "parameters")
	if params == nil {
		t.Fatal("no parameters")
	}
	param := findChildByType(params, "typed_parameter")
	if param == nil {
		t.Fatal("no typed_parameter")
	}
	got := typeText(param, src)
	if got != "int" {
		t.Fatalf("expected 'int', got %q", got)
	}

	// no type child
	got2 := typeText(root, src)
	if got2 != "" {
		t.Fatalf("expected empty, got %q", got2)
	}
}
