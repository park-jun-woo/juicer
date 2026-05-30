//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractDecoratorDeps: 정상 / nil call / args없음 분기
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

func TestExtractDecoratorDeps_Nil(t *testing.T) {
	if deps := extractDecoratorDeps(nil, nil); deps != nil {
		t.Fatalf("expected nil, got %v", deps)
	}
}

func TestExtractDecoratorDeps_NoArgs(t *testing.T) {
	// a decorator without a call (just @router) -> callNode without argument_list
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
