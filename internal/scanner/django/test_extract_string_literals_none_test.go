//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractStringLiterals_None 테스트
package django

import "testing"

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
