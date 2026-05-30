//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestParseParamNode_TypedDefault 테스트
package fastapi

import "testing"

func TestParseParamNode_TypedDefault(t *testing.T) {
	src := []byte("def f(x: int = 5): pass\n")
	root, _ := parsePython(src)
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")
	tdp := findChildByType(params, "typed_default_parameter")
	if tdp == nil {
		t.Fatal("no typed_default_parameter")
	}
	name, typeName, defVal, _, _ := parseParamNode(tdp, src)
	if name != "x" || typeName != "int" || defVal != "5" {
		t.Fatalf("got name=%q type=%q default=%q", name, typeName, defVal)
	}
}
