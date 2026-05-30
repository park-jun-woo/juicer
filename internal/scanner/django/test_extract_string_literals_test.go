//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractStringLiterals 테스트
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
