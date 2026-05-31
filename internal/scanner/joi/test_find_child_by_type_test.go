//ff:func feature=scan type=test topic=joi control=sequence
//ff:what findChildByType 첫 일치 직계 자식 반환/없음 테스트
package joi

import "testing"

func TestFindChildByType(t *testing.T) {
	root, _ := parseJoiTS(t, `f(1, 2);`)
	call := firstOfType(root, "call_expression")
	if call == nil {
		t.Fatal("no call")
	}
	if args := findChildByType(call, "arguments"); args == nil {
		t.Error("arguments child should be found")
	}
	if got := findChildByType(call, "nonexistent"); got != nil {
		t.Errorf("nonexistent: %v", got)
	}
}
