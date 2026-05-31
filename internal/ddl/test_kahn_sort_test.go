//ff:func feature=ddl type=test control=sequence
//ff:what kahnSort 결정적 위상정렬 및 순환 폴백 테스트
package ddl

import (
	"reflect"
	"testing"
)

func TestKahnSort(t *testing.T) {
	names := []string{"a", "b", "c"}
	// a -> b, c -> b : b first, then a, c alphabetically among ready
	deps := map[string]map[string]bool{
		"a": {"b": true},
		"b": {},
		"c": {"b": true},
	}
	got := kahnSort(names, deps)
	if !reflect.DeepEqual(got, []string{"b", "a", "c"}) {
		t.Errorf("acyclic: got %v", got)
	}
	// cycle a<->b: appended alphabetically
	cyc := map[string]map[string]bool{
		"a": {"b": true},
		"b": {"a": true},
	}
	got = kahnSort([]string{"a", "b"}, cyc)
	if !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("cycle: got %v", got)
	}
}
