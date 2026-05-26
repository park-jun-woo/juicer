//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestClassifyParam_InlineAnnotatedDepends 인라인 Annotated Depends 분류 테스트
package fastapi

import "testing"

func TestClassifyParam_InlineAnnotatedDepends(t *testing.T) {
	src := []byte("def f(form_data: Annotated[OAuth2PasswordRequestForm, Depends()]): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")
	pathNames := map[string]bool{}

	ri := &routeInfo{}
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if isParamNode(child) {
			classifyParam(child, src, ri, pathNames, nil)
		}
	}

	if len(ri.middleware) != 1 {
		t.Fatalf("expected 1 middleware, got %d", len(ri.middleware))
	}
	if ri.bodyType != "" {
		t.Errorf("body type should be empty, got %s", ri.bodyType)
	}
}
