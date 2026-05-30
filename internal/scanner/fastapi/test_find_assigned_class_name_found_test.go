//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindAssignedClassName_Found 테스트
package fastapi

import "testing"

func TestFindAssignedClassName_Found(t *testing.T) {
	src := []byte("app = FastAPI()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := findAssignedClassName(root, "app", src); got != "FastAPI" {
		t.Fatalf("got %q", got)
	}
}
