//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what isParamNode 테스트
package fastapi

import "testing"

func TestIsParamNode(t *testing.T) {
	src := []byte("def f(x: int, y=5, z: str = 'a'): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")

	paramCount := 0
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if isParamNode(child) {
			paramCount++
		}
	}
	if paramCount < 2 {
		t.Fatalf("expected at least 2 param nodes, got %d", paramCount)
	}
}
