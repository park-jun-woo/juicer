//ff:func feature=scan type=test topic=joi control=sequence
//ff:what collectStringArgs 인자 노드에서 구분자 제외 문자열 수집 테스트
package joi

import (
	"reflect"
	"testing"
)

func TestCollectStringArgs(t *testing.T) {
	root, src := parseJoiTS(t, `f('A', 'B');`)
	call := firstOfType(root, "call_expression")
	args := findChildByType(call, "arguments")
	if args == nil {
		t.Fatal("no arguments")
	}
	got := collectStringArgs(args, src)
	if !reflect.DeepEqual(got, []string{"A", "B"}) {
		t.Errorf("got %v", got)
	}
}
