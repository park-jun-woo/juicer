//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestCollectDecoratorDeps_None 테스트
package fastapi

import "testing"

func TestCollectDecoratorDeps_None(t *testing.T) {
	src := []byte(`
@router.get("/x")
def h():
    pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	defs := findAllByType(root, "decorated_definition")
	decorators := childrenOfType(defs[0], "decorator")
	deps := collectDecoratorDeps(decorators, src)
	if len(deps) != 0 {
		t.Fatalf("expected none, got %v", deps)
	}
}
