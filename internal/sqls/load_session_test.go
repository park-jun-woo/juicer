//ff:func feature=sql type=parse control=sequence
//ff:what TestLoadSession_NoFile 테스트
package sqls

import "testing"

func TestLoadSession_NoFile(t *testing.T) {
	_, err := LoadSession()
	if err == nil {
		t.Fatal("expected error when no session file")
	}
}
