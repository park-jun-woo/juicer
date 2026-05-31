//ff:func feature=ddl type=test control=sequence
//ff:what appendCycleFallback 미emit(순환) 노드 알파벳 추가 테스트
package ddl

import (
	"reflect"
	"testing"
)

func TestAppendCycleFallback(t *testing.T) {
	names := []string{"a", "b", "c"}
	emitted := map[string]bool{"a": true}
	got := appendCycleFallback(names, emitted, []string{"a"})
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("cyclic appended: %v", got)
	}
	// no cycle: unchanged
	all := map[string]bool{"a": true, "b": true, "c": true}
	got = appendCycleFallback(names, all, []string{"a", "b", "c"})
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("no cycle: %v", got)
	}
}
