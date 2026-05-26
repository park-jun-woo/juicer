//ff:func feature=sql type=test control=sequence
//ff:what TestAppendUnique_DuplicateCov 테스트
package sqls

import "testing"

func TestAppendUnique_DuplicateCov(t *testing.T) {
	result := appendUnique([]string{"a"}, "a")
	if len(result) != 1 {
		t.Fatal("expected 1")
	}
}
