//ff:func feature=sql type=parse control=sequence
//ff:what TestQueryExists_MissingDir 테스트
package sqls

import "testing"

func TestQueryExists_MissingDir(t *testing.T) {
	if queryExists("/nonexistent", "x") {
		t.Fatal("expected false")
	}
}
