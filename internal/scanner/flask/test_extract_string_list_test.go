//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractStringList 테스트
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

	if len(got) != 2 || got[0] != "GET" || got[1] != "POST" {
		t.Fatalf("got %v", got)
	}
}
