//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractChoiceValues — choices 리스트 추출 분기를 검증
package django

import "testing"

func TestExtractChoiceValues(t *testing.T) {
	src := []byte("x = ChoiceField(choices=['a', 'b'])\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if len(kw) == 0 {
		t.Fatal("no keyword_argument")
	}
	choices := extractChoiceValues(kw[0], src)
	if len(choices) != 2 || choices[0] != "a" || choices[1] != "b" {
		t.Fatalf("expected [a b], got %v", choices)
	}
}

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

func TestExtractChoiceValues_EmptyList(t *testing.T) {
	src := []byte("x = ChoiceField(choices=[])\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if c := extractChoiceValues(kw[0], src); c != nil {
		t.Fatalf("expected nil for empty list (no string values), got %v", c)
	}
}
