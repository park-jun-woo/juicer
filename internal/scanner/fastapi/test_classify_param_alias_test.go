//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestClassifyParam_TypeAlias 타입 별칭 미들웨어 분류 테스트
package fastapi

import "testing"

func TestClassifyParam_TypeAlias(t *testing.T) {
	src := []byte("def f(session: SessionDep, user: CurrentUser): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")
	pathNames := map[string]bool{}
	aliasMap := map[string]string{
		"SessionDep":  "get_db",
		"CurrentUser": "get_current_user",
	}

	ri := &routeInfo{}
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if isParamNode(child) {
			classifyParam(child, src, ri, pathNames, aliasMap)
		}
	}

	if len(ri.middleware) != 2 {
		t.Fatalf("expected 2 middleware, got %d", len(ri.middleware))
	}
	if ri.middleware[0] != "get_db" {
		t.Errorf("expected get_db, got %s", ri.middleware[0])
	}
	if ri.middleware[1] != "get_current_user" {
		t.Errorf("expected get_current_user, got %s", ri.middleware[1])
	}
	if ri.bodyType != "" {
		t.Errorf("body type should be empty, got %s", ri.bodyType)
	}
}
