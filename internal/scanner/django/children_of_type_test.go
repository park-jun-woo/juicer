//ff:func feature=scan type=test control=sequence topic=django
//ff:what childrenOfType — 직접 자식 중 지정 타입 수집을 검증
package django

import "testing"

func TestChildrenOfType(t *testing.T) {
	src := []byte("x = [1, 2, 3]\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	lst := findAllByType(root, "list")
	if len(lst) == 0 {
		t.Fatal("no list node")
	}
	ints := childrenOfType(lst[0], "integer")
	if len(ints) != 3 {
		t.Fatalf("expected 3 integers, got %d", len(ints))
	}
	if got := childrenOfType(lst[0], "string"); len(got) != 0 {
		t.Fatalf("expected 0 strings, got %d", len(got))
	}
}
