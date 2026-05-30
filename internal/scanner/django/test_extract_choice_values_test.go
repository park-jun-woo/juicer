//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractChoiceValues 테스트
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
