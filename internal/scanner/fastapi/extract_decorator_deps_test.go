//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDecoratorDeps_Single 단일 dependencies 추출 테스트
package fastapi

import "testing"

func TestExtractDecoratorDeps_Single(t *testing.T) {
	src := []byte(`
@router.get("/admin", dependencies=[Depends(verify_admin)])
async def admin_panel():
    pass
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	defs := findAllByType(root, "decorated_definition")
	if len(defs) == 0 {
		t.Fatal("no decorated_definition found")
	}
	decorators := childrenOfType(defs[0], "decorator")
	if len(decorators) == 0 {
		t.Fatal("no decorator found")
	}
	callNode, _ := findDecoratorNodes(decorators[0])
	deps := extractDecoratorDeps(callNode, src)
	if len(deps) != 1 {
		t.Fatalf("expected 1 dep, got %d", len(deps))
	}
	if deps[0] != "verify_admin" {
		t.Errorf("expected verify_admin, got %s", deps[0])
	}
}
