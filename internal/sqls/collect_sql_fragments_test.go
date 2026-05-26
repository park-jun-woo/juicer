//ff:func feature=sql type=test control=sequence
//ff:what TestCollectSQLFragments_Nil 테스트
package sqls

import "testing"

func TestCollectSQLFragments_Nil(t *testing.T) {
	if collectSQLFragments(nil) != nil {
		t.Fatal("expected nil")
	}
}

