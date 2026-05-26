//ff:func feature=sql type=test control=sequence
//ff:what TestQueryExists_BadDirCov 테스트
package sqls

import "testing"

func TestQueryExists_BadDirCov(t *testing.T) {
	if queryExists("/nonexistent", "x") {
		t.Fatal("expected false")
	}
}
