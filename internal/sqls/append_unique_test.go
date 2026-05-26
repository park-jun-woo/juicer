//ff:func feature=sql type=test control=sequence
//ff:what TestAppendUnique_New 테스트
package sqls

import "testing"

func TestAppendUnique_New(t *testing.T) {
	result := appendUnique([]string{"a"}, "b")
	if len(result) != 2 {
		t.Fatal("expected 2")
	}
}
