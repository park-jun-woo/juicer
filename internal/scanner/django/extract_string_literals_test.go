//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractStringLiterals — 리스트 노드의 따옴표 제거 문자열 추출을 검증
package django

import "testing"

func TestExtractStringLiterals(t *testing.T) {
	src := []byte("x = ['id', 'name', 7]\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	lists := findAllByType(root, "list")
	if len(lists) == 0 {
		t.Fatal("no list node")
	}
	got := extractStringLiterals(lists[0], src)
	if len(got) != 2 || got[0] != "id" || got[1] != "name" {
		t.Fatalf("got %v, want [id name]", got)
	}
}

func TestExtractStringLiterals_None(t *testing.T) {
	src := []byte("x = [1]\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	lists := findAllByType(root, "list")
	if got := extractStringLiterals(lists[0], src); len(got) != 0 {
		t.Fatalf("expected none, got %v", got)
	}
}
