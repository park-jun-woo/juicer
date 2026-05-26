//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what collectEndpoints 테스트
package fastapi

import "testing"

func TestCollectEndpoints(t *testing.T) {
	src := []byte("@router.get('/users')\nasync def list_users(): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	files := []fileInfo{
		{absPath: "/main.py", relPath: "main.py", src: src, root: root, prefixes: map[string]string{}},
	}
	eps, reqs := collectEndpoints(files)
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	_ = reqs

	// empty
	eps2, _ := collectEndpoints(nil)
	if len(eps2) != 0 {
		t.Fatalf("expected 0, got %d", len(eps2))
	}
}
