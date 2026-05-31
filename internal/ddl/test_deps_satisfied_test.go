//ff:func feature=ddl type=test control=sequence
//ff:what depsSatisfied 모든 의존성 emit 여부 테스트
package ddl

import "testing"

func TestDepsSatisfied(t *testing.T) {
	emitted := map[string]bool{"a": true, "b": true}
	if !depsSatisfied(map[string]bool{"a": true}, emitted) {
		t.Error("a emitted should be satisfied")
	}
	if depsSatisfied(map[string]bool{"a": true, "c": true}, emitted) {
		t.Error("c not emitted should be unsatisfied")
	}
	if !depsSatisfied(map[string]bool{}, emitted) {
		t.Error("no deps should be satisfied")
	}
}
