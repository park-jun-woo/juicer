//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestExtractBodyKwargs_EmbedFalse 테스트
package fastapi

import "testing"

func TestExtractBodyKwargs_EmbedFalse(t *testing.T) {
	src := []byte(`def f(item: Item = Body(embed=False)): pass` + "\n")
	root, _ := parsePython(src)
	fn := findChildByType(root, "function_definition")
	params := findChildByType(fn, "parameters")
	for i := 0; i < int(params.ChildCount()); i++ {
		c := params.Child(i)
		if !isParamNode(c) {
			continue
		}
		if _, embed := extractBodyKwargs(c, src); embed {
			t.Fatalf("expected embed=false")
		}
		return
	}
}
