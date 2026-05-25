//ff:func feature=scan type=extract control=sequence
//ff:what TestLcFirst_SMSResult 테스트
package scanner

import "testing"

func TestLcFirst_SMSResult(t *testing.T) {
	if lcFirst("SMSResult") != "smsResult" {
		t.Fatal("expected smsResult")
	}
}
