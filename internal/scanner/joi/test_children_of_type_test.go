//ff:func feature=scan type=test topic=joi control=sequence
//ff:what childrenOfType 직계 자식 타입 필터 테스트
package joi

import "testing"

func TestChildrenOfType(t *testing.T) {
	root, _ := parseJoiTS(t, `const o = { a: 1, b: 2 };`)
	obj := firstOfType(root, "object")
	if obj == nil {
		t.Fatal("no object")
	}
	pairs := childrenOfType(obj, "pair")
	if len(pairs) != 2 {
		t.Errorf("want 2 pairs, got %d", len(pairs))
	}
	if got := childrenOfType(obj, "nonexistent"); got != nil {
		t.Errorf("nonexistent type: %v", got)
	}
}
