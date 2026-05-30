//ff:func feature=scan type=test control=sequence topic=django
//ff:what nodeText — 노드 소스 텍스트 반환을 검증
package django

import "testing"

func TestNodeText(t *testing.T) {
	src := []byte("foo_bar\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	ids := findAllByType(root, "identifier")
	if len(ids) == 0 {
		t.Fatal("no identifier")
	}
	if got := nodeText(ids[0], src); got != "foo_bar" {
		t.Fatalf("got %q, want foo_bar", got)
	}
}
