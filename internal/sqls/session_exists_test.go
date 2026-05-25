//ff:func feature=sql type=parse control=sequence
//ff:what TestSessionExists_NoFile 테스트
package sqls

import "testing"

func TestSessionExists_NoFile(t *testing.T) {
	if SessionExists() {
		t.Fatal("expected false in test environment")
	}
}
