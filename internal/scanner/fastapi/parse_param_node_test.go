//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what parseParamNode 테스트
package fastapi

import "testing"

func TestParseParamNode(t *testing.T) {
	// typed_parameter
	src := []byte("def f(x: int): pass\n")
	root, _ := parsePython(src)
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")
	tp := findChildByType(params, "typed_parameter")
	if tp != nil {
		name, typeName, _, _ := parseParamNode(tp, src)
		if name != "x" || typeName != "int" {
			t.Fatalf("typed: name=%q type=%q", name, typeName)
		}
	}

	// default_parameter
	src2 := []byte("def f(x=5): pass\n")
	root2, _ := parsePython(src2)
	funcDef2 := findChildByType(root2, "function_definition")
	params2 := findChildByType(funcDef2, "parameters")
	dp := findChildByType(params2, "default_parameter")
	if dp != nil {
		name, _, defVal, _ := parseParamNode(dp, src2)
		if name != "x" || defVal != "5" {
			t.Fatalf("default: name=%q default=%q", name, defVal)
		}
	}
}
