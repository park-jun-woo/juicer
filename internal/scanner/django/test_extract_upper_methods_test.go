//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractUpperMethods 테스트
package django

import "testing"

func TestExtractUpperMethods(t *testing.T) {
	src := []byte("x = action(methods=['get', 'post'])\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if len(kw) == 0 {
		t.Fatal("no keyword_argument")
	}
	methods := extractUpperMethods(kw[0], src)
	if len(methods) != 2 || methods[0] != "GET" || methods[1] != "POST" {
		t.Fatalf("got %v, want [GET POST]", methods)
	}
}
