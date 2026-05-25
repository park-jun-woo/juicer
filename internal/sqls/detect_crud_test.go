//ff:func feature=sql type=parse control=sequence
//ff:what TestDetectCRUD_Nil 테스트
package sqls

import "testing"

func TestDetectCRUD_Nil(t *testing.T) {
	if detectCRUD(nil) != "" {
		t.Fatal("expected empty")
	}
}
