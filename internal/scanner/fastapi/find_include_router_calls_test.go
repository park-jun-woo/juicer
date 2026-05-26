//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findIncludeRouterCalls 테스트
package fastapi

import "testing"

func TestFindIncludeRouterCalls(t *testing.T) {
	src := []byte("app.include_router(users_router)\napp.include_router(orders_router, prefix='/orders')\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	calls := findIncludeRouterCalls(root, src)
	if len(calls) != 2 {
		t.Fatalf("expected 2 calls, got %d", len(calls))
	}
}
