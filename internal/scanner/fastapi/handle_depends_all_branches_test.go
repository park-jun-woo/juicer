//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestHandleDepends_AllBranches valid/empty depends 전 분기 테스트
package fastapi

import "testing"

func TestHandleDepends_AllBranches(t *testing.T) {
	// valid depends
	ri := &routeInfo{}
	handleDepends("user", "Depends(get_current_user)", ri)
	if len(ri.middleware) != 1 || ri.middleware[0] != "get_current_user" {
		t.Fatal("expected middleware entry")
	}

	// empty depends
	ri2 := &routeInfo{}
	handleDepends("x", "Depends()", ri2)
	if len(ri2.middleware) != 0 {
		t.Fatal("expected no middleware for empty Depends()")
	}
}
