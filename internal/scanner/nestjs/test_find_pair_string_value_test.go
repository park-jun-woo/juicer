//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindPairStringValue 테스트
package nestjs

import "testing"

func TestFindPairStringValue(t *testing.T) {
	src := []byte(`const o = { name: 'hello' };`)
	root, _ := parseTypeScript(src)
	pairs := findAllByType(root, "pair")
	if got := findPairStringValue(pairs[0], src); got != "hello" {
		t.Fatalf("got %q", got)
	}
	src2 := []byte(`const o = { count: 5 };`)
	root2, _ := parseTypeScript(src2)
	pairs2 := findAllByType(root2, "pair")
	if got := findPairStringValue(pairs2[0], src2); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
