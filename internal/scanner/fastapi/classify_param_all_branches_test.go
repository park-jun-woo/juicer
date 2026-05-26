//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestClassifyParam_AllBranches self/UploadFile/Query/Pydantic/typed 전 분기 테스트
package fastapi

import "testing"

func TestClassifyParam_AllBranches(t *testing.T) {
	// self param (should be skipped), UploadFile, Query() default, Pydantic body, typed query
	src := []byte(`def f(self, photo: UploadFile, limit: int = Query(10), data: UserCreate = None, tag: str): pass
`)
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

	if len(ri.files) == 0 {
		t.Fatal("expected file upload param for UploadFile")
	}
	if ri.bodyType == "" {
		t.Fatal("expected body type for UserCreate")
	}
}
