//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what classifyByDefault 테스트
package fastapi

import "testing"

func TestClassifyByDefault(t *testing.T) {
	// Query
	ri := &routeInfo{}
	classifyByDefault("Query", "limit", "int", "", nil, nil, ri)
	if len(ri.query) != 1 || ri.query[0].Name != "limit" {
		t.Fatalf("Query: got %+v", ri)
	}

	// Body (nil param — no kwargs extraction)
	ri2 := &routeInfo{}
	classifyByDefault("Body", "data", "UserCreate", "", nil, nil, ri2)
	if ri2.bodyType != "UserCreate" {
		t.Fatalf("Body: got %q", ri2.bodyType)
	}
	if ri2.bodyVarName != "data" {
		t.Fatalf("Body var_name: got %q", ri2.bodyVarName)
	}

	// File
	ri3 := &routeInfo{}
	classifyByDefault("File", "upload", "UploadFile", "", nil, nil, ri3)
	if len(ri3.files) != 1 {
		t.Fatalf("File: got %d files", len(ri3.files))
	}

	// Depends
	ri4 := &routeInfo{}
	classifyByDefault("Depends", "user", "", "Depends(get_current_user)", nil, nil, ri4)
	if len(ri4.middleware) != 1 {
		t.Fatalf("Depends: got %d middleware", len(ri4.middleware))
	}
}
