//ff:func feature=sql type=parse control=sequence
//ff:what TestDetectDynamic_Nil 테스트
package sqls

import "testing"

func TestDetectDynamic_Nil(t *testing.T) {
	if detectDynamic(nil) {
		t.Fatal("expected false")
	}
}
