//ff:func feature=scan type=test control=sequence topic=flask
//ff:what extractStringList 테스트
package flask

import "testing"

func TestExtractStringList(t *testing.T) {
	b := []byte("x = ['get', 'post', 5]\n")
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	lists := findAllByType(root, "list")
	if len(lists) == 0 {
		t.Fatal("no list")
	}
	got := extractStringList(lists[0], b)
	// strings uppercased; non-string element (5) skipped
	if len(got) != 2 || got[0] != "GET" || got[1] != "POST" {
		t.Fatalf("got %v", got)
	}
}

func TestExtractStringList_Empty(t *testing.T) {
	b := []byte("x = []\n")
	root, _ := parsePython(b)
	lists := findAllByType(root, "list")
	if got := extractStringList(lists[0], b); len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}
}
