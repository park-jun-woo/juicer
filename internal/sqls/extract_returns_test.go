//ff:func feature=sql type=test control=sequence
//ff:what TestExtractReturns_Nil 테스트
package sqls

import "testing"

func TestExtractReturns_Nil(t *testing.T) {
	if extractReturns(nil) != nil {
		t.Fatal("expected nil")
	}
}

