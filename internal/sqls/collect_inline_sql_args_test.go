//ff:func feature=sql type=test control=sequence
//ff:what TestCollectInlineSQLArgs_Nil 테스트
package sqls

import "testing"

func TestCollectInlineSQLArgs_Nil(t *testing.T) {
	result := collectInlineSQLArgs(nil)
	if result != nil {
		t.Fatal("expected nil")
	}
}
