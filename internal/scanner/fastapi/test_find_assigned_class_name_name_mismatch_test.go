//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindAssignedClassName_NameMismatch 테스트
package fastapi

import "testing"

func TestFindAssignedClassName_NameMismatch(t *testing.T) {
	src := []byte("app = FastAPI()\n")
	root, _ := parsePython(src)
	if got := findAssignedClassName(root, "other", src); got != "" {
		t.Fatalf("got %q", got)
	}
}
