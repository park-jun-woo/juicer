//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractReturnType 테스트
package fastapi

import "testing"

func TestExtractReturnType(t *testing.T) {
	src := []byte("async def get_user(user_id: int) -> UserResponse:\n    pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	funcDef := findChildByType(root, "function_definition")
	if funcDef == nil {
		t.Fatal("no function_definition")
	}
	ri := &routeInfo{}
	extractReturnType(funcDef, src, ri)
	if ri.returnType != "UserResponse" {
		t.Fatalf("expected 'UserResponse', got %q", ri.returnType)
	}

	// no return type
	src2 := []byte("def f(): pass\n")
	root2, _ := parsePython(src2)
	funcDef2 := findChildByType(root2, "function_definition")
	ri2 := &routeInfo{}
	extractReturnType(funcDef2, src2, ri2)
	if ri2.returnType != "" {
		t.Fatalf("expected empty, got %q", ri2.returnType)
	}
}
