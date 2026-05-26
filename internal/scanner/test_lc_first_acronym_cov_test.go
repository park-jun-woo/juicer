//ff:func feature=scan type=test control=sequence
//ff:what TestLcFirst_AcronymCov 테스트
package scanner

import "testing"

func TestLcFirst_AcronymCov(t *testing.T) {
	if got := lcFirst("SMSResult"); got != "smsResult" {
		t.Fatalf("expected smsResult, got %s", got)
	}
}
