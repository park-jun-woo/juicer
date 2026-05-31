//ff:func feature=ddl type=test control=sequence
//ff:what pickReady 의존성 충족된 알파벳 최선두 노드 선택 테스트
package ddl

import "testing"

func TestPickReady(t *testing.T) {
	names := []string{"a", "b", "c"}
	deps := map[string]map[string]bool{
		"a": {"b": true},
		"b": {},
		"c": {},
	}
	emitted := map[string]bool{}
	// b is first alphabetically with no unmet deps (a depends on b)
	if got := pickReady(names, deps, emitted); got != "b" {
		t.Errorf("got %q, want b", got)
	}
	emitted["b"] = true
	if got := pickReady(names, deps, emitted); got != "a" {
		t.Errorf("after b: got %q, want a", got)
	}
	// none ready
	emitted = map[string]bool{"a": true, "b": true, "c": true}
	if got := pickReady(names, deps, emitted); got != "" {
		t.Errorf("all emitted: got %q", got)
	}
}
