//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findAssignedClassName: var=Class() 매칭 / 이름불일치 / call없음 / 없음
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

func TestFindAssignedClassName_NameMismatch(t *testing.T) {
	src := []byte("app = FastAPI()\n")
	root, _ := parsePython(src)
	if got := findAssignedClassName(root, "other", src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestFindAssignedClassName_NoCall(t *testing.T) {
	src := []byte("x = 5\n")
	root, _ := parsePython(src)
	if got := findAssignedClassName(root, "x", src); got != "" {
		t.Fatalf("got %q", got)
	}
}
