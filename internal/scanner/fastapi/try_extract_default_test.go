//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what tryExtractDefault 테스트
package fastapi

import "testing"

func TestTryExtractDefault(t *testing.T) {
	src := []byte("def f(x: int = Query(default=5)): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")
	param := findChildByType(params, "typed_default_parameter")
	if param == nil {
		t.Skip("no typed_default_parameter found in AST")
	}

	foundCall := false
	for i := 0; i < int(param.ChildCount()); i++ {
		child := param.Child(i)
		val, call := tryExtractDefault(child, src)
		if call == "" {
			continue
		}
		foundCall = true
		if val == "" {
			t.Fatal("expected non-empty val for call")
		}
	}
	if !foundCall {
		tryExtractDefaultFallback(t)
	}
}
