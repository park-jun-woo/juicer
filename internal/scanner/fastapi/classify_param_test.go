//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestClassifyParam path/query/body/file/default/self/empty 분기 테스트
package fastapi

import "testing"

func TestClassifyParam(t *testing.T) {
	src := []byte("def f(user_id: int, q: str = 'default', body: UserCreate = None): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")
	pathNames := map[string]bool{"user_id": true}

	ri := &routeInfo{}
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if isParamNode(child) {
			classifyParam(child, src, ri, pathNames)
		}
	}

	if len(ri.params) == 0 {
		t.Fatal("expected at least 1 path param")
	}
}
