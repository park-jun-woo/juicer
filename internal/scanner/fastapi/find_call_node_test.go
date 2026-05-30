//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what findCallNode call 자식 노드 탐색 테스트
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

func TestFindCallNode_NotFound(t *testing.T) {
	src := []byte("x = 5\n")
	root, _ := parsePython(src)
	assigns := findAllByType(root, "assignment")
	if len(assigns) == 0 {
		t.Fatal("no assignment")
	}
	if findCallNode(assigns[0]) != nil {
		t.Fatal("expected nil for no call child")
	}
}
