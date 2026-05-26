//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what classifyParam isPydanticModelType 분기에서 bodyVarName 설정 테스트
package fastapi

import "testing"

func TestClassifyParam_PydanticVarName(t *testing.T) {
	src := []byte("def f(user: UserCreate): pass\n")
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

	if ri.bodyType != "UserCreate" {
		t.Fatalf("bodyType: got %q, want UserCreate", ri.bodyType)
	}
	if ri.bodyVarName != "user" {
		t.Fatalf("bodyVarName: got %q, want user", ri.bodyVarName)
	}
}
