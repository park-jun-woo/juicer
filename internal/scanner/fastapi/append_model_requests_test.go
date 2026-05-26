//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what appendModelRequests 테스트 — body/response 모델 요청 추가
package fastapi

import "testing"

func TestAppendModelRequests(t *testing.T) {
	// Case 1: bodyType set, no response model
	ri := routeInfo{bodyType: "CreateUser"}
	fi := fileInfo{absPath: "/app/main.py"}
	reqs := appendModelRequests(nil, ri, fi, 0)
	if len(reqs) != 1 {
		t.Fatalf("expected 1 request, got %d", len(reqs))
	}
	if !reqs[0].isBody {
		t.Fatal("expected isBody=true")
	}

	// Case 2: responseModel is a Pydantic type
	ri2 := routeInfo{responseModel: "UserOut"}
	reqs2 := appendModelRequests(nil, ri2, fi, 1)
	if len(reqs2) != 1 {
		t.Fatalf("expected 1 request, got %d", len(reqs2))
	}
	if reqs2[0].isBody {
		t.Fatal("expected isBody=false")
	}

	// Case 3: returnType used when responseModel is empty
	ri3 := routeInfo{returnType: "OrderResponse"}
	reqs3 := appendModelRequests(nil, ri3, fi, 2)
	if len(reqs3) != 1 {
		t.Fatalf("expected 1 request, got %d", len(reqs3))
	}

	// Case 4: both body and response
	ri4 := routeInfo{bodyType: "CreateOrder", responseModel: "OrderOut"}
	reqs4 := appendModelRequests(nil, ri4, fi, 3)
	if len(reqs4) != 2 {
		t.Fatalf("expected 2 requests, got %d", len(reqs4))
	}

	// Case 5: no body, builtin return type (not pydantic)
	ri5 := routeInfo{returnType: "str"}
	reqs5 := appendModelRequests(nil, ri5, fi, 4)
	if len(reqs5) != 0 {
		t.Fatalf("expected 0 requests, got %d", len(reqs5))
	}

	// Case 6: empty everything
	ri6 := routeInfo{}
	reqs6 := appendModelRequests(nil, ri6, fi, 5)
	if len(reqs6) != 0 {
		t.Fatalf("expected 0 requests, got %d", len(reqs6))
	}
}
