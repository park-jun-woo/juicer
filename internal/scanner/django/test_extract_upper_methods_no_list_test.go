//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractUpperMethods_NoList 테스트
package django

import "testing"

func TestExtractUpperMethods_NoList(t *testing.T) {
	src := []byte("x = action(methods=DEFAULT)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if m := extractUpperMethods(kw[0], src); m != nil {
		t.Fatalf("expected nil for non-list methods, got %v", m)
	}
}
