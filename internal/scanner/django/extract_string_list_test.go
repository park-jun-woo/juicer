//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractStringList — 리스트 노드의 문자열 값(대문자화) 추출을 검증
package django

import "testing"

func TestExtractStringList(t *testing.T) {
	src := []byte("x = ['get', 'post', 123]\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	lists := findAllByType(root, "list")
	if len(lists) == 0 {
		t.Fatal("no list node")
	}
	got := extractStringList(lists[0], src)
	if len(got) != 2 {
		t.Fatalf("expected 2 strings, got %v", got)
	}
	// values are upper-cased; the integer 123 is skipped.
	if got[0] != "GET" || got[1] != "POST" {
		t.Errorf("got %v, want [GET POST]", got)
	}
}

func TestExtractStringList_NoStrings(t *testing.T) {
	src := []byte("x = [1, 2]\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	lists := findAllByType(root, "list")
	if got := extractStringList(lists[0], src); len(got) != 0 {
		t.Fatalf("expected no strings, got %v", got)
	}
}
