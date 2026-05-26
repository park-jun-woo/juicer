//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestClassifyParam_InlineAnnotatedDependsWithFunc 인라인 Annotated Depends(func) 분류 테스트
package fastapi

import "testing"

func TestClassifyParam_InlineAnnotatedDependsWithFunc(t *testing.T) {
	src := []byte("def f(db: Annotated[Session, Depends(get_db)]): pass\n")
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
	if ri.middleware[0] != "get_db" {
		t.Errorf("expected get_db, got %s", ri.middleware[0])
	}
}
