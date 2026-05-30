//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDecoratorDeps_NoArgs 테스트
package fastapi

import "testing"

func TestExtractDecoratorDeps_NoArgs(t *testing.T) {

	src := []byte(`
@router
def h():
    pass
`)
	root, _ := parsePython(src)
	defs := findAllByType(root, "decorated_definition")
	decorators := childrenOfType(defs[0], "decorator")
	callNode, _ := findDecoratorNodes(decorators[0])
	deps := extractDecoratorDeps(callNode, src)
	if deps != nil {
		t.Fatalf("expected nil, got %v", deps)
	}
}
