//ff:func feature=sql type=parse control=sequence
//ff:what TestExtractParams_Nil 테스트
package sqls

import "testing"

func TestExtractParams_Nil(t *testing.T) {
	if extractParams(nil) != nil {
		t.Fatal("expected nil")
	}
}
