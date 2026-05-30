//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractChoiceValues_NoList 테스트
package django

import "testing"

func TestExtractChoiceValues_NoList(t *testing.T) {
	src := []byte("x = ChoiceField(choices=MY_CHOICES)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if c := extractChoiceValues(kw[0], src); c != nil {
		t.Fatalf("expected nil for non-list choices, got %v", c)
	}
}
