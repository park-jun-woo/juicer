//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractStringList_Empty 테스트
package flask

import "testing"

func TestExtractStringList_Empty(t *testing.T) {
	b := []byte("x = []\n")
	root, _ := parsePython(b)
	lists := findAllByType(root, "list")
	if got := extractStringList(lists[0], b); len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}
}
