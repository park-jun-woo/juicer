//ff:func feature=sql type=parse control=sequence
//ff:what TestAppendUnique_Duplicate 테스트
package sqls

import "testing"

func TestAppendUnique_Duplicate(t *testing.T) {
	result := appendUnique([]string{"a"}, "a")
	if len(result) != 1 {
		t.Fatal("expected 1")
	}
}
