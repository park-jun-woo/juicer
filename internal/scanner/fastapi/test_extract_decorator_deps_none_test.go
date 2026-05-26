//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDecoratorDeps_NoDependencies dependencies 미존재 테스트
package fastapi

import "testing"

func TestExtractDecoratorDeps_NoDependencies(t *testing.T) {
	src := []byte(`
@router.get("/users")
async def list_users():
    pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	defs := findAllByType(root, "decorated_definition")
	decorators := childrenOfType(defs[0], "decorator")
	callNode, _ := findDecoratorNodes(decorators[0])
	deps := extractDecoratorDeps(callNode, src)
	if len(deps) != 0 {
		t.Fatalf("expected 0 deps, got %d", len(deps))
	}
}
