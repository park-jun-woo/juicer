//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractStringList 테스트
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

	if got[0] != "GET" || got[1] != "POST" {
		t.Errorf("got %v, want [GET POST]", got)
	}
}
