//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDecoratorDeps_Multiple 복수 dependencies 추출 테스트
package fastapi

import "testing"

func TestExtractDecoratorDeps_Multiple(t *testing.T) {
	src := []byte(`
@router.get("/secure", dependencies=[Depends(verify_token), Depends(log_request)])
async def secure_endpoint():
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
	if len(deps) != 2 {
		t.Fatalf("expected 2 deps, got %d", len(deps))
	}
	if deps[0] != "verify_token" {
		t.Errorf("expected verify_token, got %s", deps[0])
	}
	if deps[1] != "log_request" {
		t.Errorf("expected log_request, got %s", deps[1])
	}
}
