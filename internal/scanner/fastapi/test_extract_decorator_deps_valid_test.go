//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDecoratorDeps_Valid 테스트
package fastapi

import "testing"

func TestExtractDecoratorDeps_Valid(t *testing.T) {
	src := []byte(`
@router.get("/x", dependencies=[Depends(auth)])
def h():
    pass
`)
	root, _ := parsePython(src)
	defs := findAllByType(root, "decorated_definition")
	decorators := childrenOfType(defs[0], "decorator")
	callNode, _ := findDecoratorNodes(decorators[0])
	deps := extractDecoratorDeps(callNode, src)
	if len(deps) != 1 || deps[0] != "auth" {
		t.Fatalf("got %v", deps)
	}
}
