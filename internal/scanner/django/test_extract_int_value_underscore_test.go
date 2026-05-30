//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractIntValue_Underscore 테스트
package django

import "testing"

func TestExtractIntValue_Underscore(t *testing.T) {

	src := []byte("x = CharField(max_length=1_000)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	v := extractIntValue(kw[0], src)
	if v == nil || *v != 1000 {
		t.Fatalf("expected 1000, got %v", v)
	}
}
