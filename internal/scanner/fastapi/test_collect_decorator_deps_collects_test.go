//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestCollectDecoratorDeps_Collects 테스트
package fastapi

import "testing"

func TestCollectDecoratorDeps_Collects(t *testing.T) {
	src := []byte(`
@router.get("/x", dependencies=[Depends(auth), Depends(log)])
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
	if len(deps) != 2 || deps[0] != "auth" || deps[1] != "log" {
		t.Fatalf("got %v", deps)
	}
}
