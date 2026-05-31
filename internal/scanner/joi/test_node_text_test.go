//ff:func feature=scan type=test control=sequence topic=joi
//ff:what nodeText 노드 소스 텍스트 추출 테스트
package joi

import "testing"

func TestNodeText(t *testing.T) {
	root, src := parseJoiTS(t, `const a = 5;`)
	id := firstOfType(root, "identifier")
	if id == nil {
		t.Fatal("no identifier")
	}
	if got := nodeText(id, src); got != "a" {
		t.Errorf("got %q, want a", got)
	}
}
