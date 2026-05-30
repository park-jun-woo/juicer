//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractIntValue_NoInteger 테스트
package django

import "testing"

func TestExtractIntValue_NoInteger(t *testing.T) {
	src := []byte("x = CharField(required=False)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if v := extractIntValue(kw[0], src); v != nil {
		t.Fatalf("expected nil for non-integer value, got %v", *v)
	}
}
