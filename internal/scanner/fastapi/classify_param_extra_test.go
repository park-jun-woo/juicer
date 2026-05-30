//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what classifyParam: path/defaultVal/None/Depends/self-skip 추가 분기
package fastapi

import "testing"

func classifyAllParams(src []byte, ri *routeInfo, pathNames map[string]bool, aliasMap map[string]string) {
	root, err := parsePython(src)
	if err != nil {
		return
	}
	funcDef := findChildByType(root, "function_definition")
	params := findChildByType(funcDef, "parameters")
	for i := 0; i < int(params.ChildCount()); i++ {
		child := params.Child(i)
		if isParamNode(child) {
			classifyParam(child, src, ri, pathNames, aliasMap)
		}
	}
}

func TestClassifyParam_PathParam(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(item_id: int): pass\n"), ri, map[string]bool{"item_id": true}, nil)
	if len(ri.params) != 1 || ri.params[0].Name != "item_id" {
		t.Fatalf("expected path param, got %+v", ri.params)
	}
}

func TestClassifyParam_DefaultValueQuery(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(limit: int = 10): pass\n"), ri, map[string]bool{}, nil)
	if len(ri.query) != 1 || ri.query[0].Default != "10" {
		t.Fatalf("expected query with default, got %+v", ri.query)
	}
}

func TestClassifyParam_NoneDefault(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(q: str = None): pass\n"), ri, map[string]bool{}, nil)
	if len(ri.query) != 1 || !ri.query[0].DefaultIsNull {
		t.Fatalf("expected nullable query, got %+v", ri.query)
	}
}

func TestClassifyParam_TypedQuery(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(tag: str): pass\n"), ri, map[string]bool{}, nil)
	if len(ri.query) != 1 || ri.query[0].Name != "tag" {
		t.Fatalf("expected typed query, got %+v", ri.query)
	}
}

func TestClassifyParam_SelfSkipped(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(self, cls): pass\n"), ri, map[string]bool{}, nil)
	if len(ri.query) != 0 || len(ri.params) != 0 {
		t.Fatalf("self/cls should be skipped, got %+v", ri)
	}
}

func TestClassifyParam_AnnotatedDepends(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(user: Annotated[User, Depends(get_user)]): pass\n"), ri, map[string]bool{}, nil)
	if len(ri.middleware) != 1 || ri.middleware[0] != "get_user" {
		t.Fatalf("expected middleware, got %+v", ri.middleware)
	}
}
