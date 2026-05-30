//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractStringList_NoStrings 테스트
package django

import "testing"

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
