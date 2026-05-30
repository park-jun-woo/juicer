//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractIntValue 테스트
package django

import "testing"

func TestExtractIntValue(t *testing.T) {
	src := []byte("x = CharField(max_length=100)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if len(kw) == 0 {
		t.Fatal("no keyword_argument")
	}
	v := extractIntValue(kw[0], src)
	if v == nil || *v != 100 {
		t.Fatalf("expected 100, got %v", v)
	}
}
