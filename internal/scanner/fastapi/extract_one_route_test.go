//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractOneRoute 테스트
package fastapi

import "testing"

func TestExtractOneRoute(t *testing.T) {
	src := []byte("@router.get('/users/{user_id}')\nasync def get_user(user_id: int) -> UserResponse:\n    pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	defs := findAllByType(root, "decorated_definition")
	if len(defs) == 0 {
		t.Fatal("no decorated_definition")
	}
	prefixes := map[string]string{"router": "/api"}
	ri := extractOneRoute(defs[0], src, prefixes, "main.py")
	if ri == nil {
		t.Fatal("expected route info")
	}
	if ri.method != "GET" {
		t.Fatalf("method: got %q", ri.method)
	}
	if ri.handler != "get_user" {
		t.Fatalf("handler: got %q", ri.handler)
	}

	// non-route decorator
	src2 := []byte("@some_decorator\ndef f(): pass\n")
	root2, _ := parsePython(src2)
	defs2 := findAllByType(root2, "decorated_definition")
	if len(defs2) > 0 {
		ri2 := extractOneRoute(defs2[0], src2, nil, "x.py")
		if ri2 != nil {
			t.Fatal("expected nil for non-route decorator")
		}
	}
}
