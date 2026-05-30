//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractBodyKwargs_NoBodyCall 테스트
package fastapi

import "testing"

func TestExtractBodyKwargs_NoBodyCall(t *testing.T) {
	src := []byte("def f(x: int): pass\n")
	root, _ := parsePython(src)
	fn := findChildByType(root, "function_definition")
	params := findChildByType(fn, "parameters")
	tp := findChildByType(params, "typed_parameter")
	if tp == nil {
		t.Skip("no typed_parameter")
	}
	alias, embed := extractBodyKwargs(tp, src)
	if alias != "" || embed {
		t.Fatalf("alias=%q embed=%v", alias, embed)
	}
}
