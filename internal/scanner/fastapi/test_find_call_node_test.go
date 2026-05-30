//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestFindCallNode 테스트
package fastapi

import "testing"

func TestFindCallNode(t *testing.T) {
	src := []byte(`def f(x: int = Body(...)): pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")

	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if !isParamNode(child) {
			continue
		}
		callNode := findCallNode(child)
		if callNode == nil {
			t.Fatal("expected call node")
		}
		if callNode.Type() != "call" {
			t.Fatalf("expected call type, got %s", callNode.Type())
		}
		return
	}
	t.Fatal("no param node found")
}
