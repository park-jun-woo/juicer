//ff:func feature=scan type=test control=sequence
//ff:what TestGoTypeFormat_TimeCov 테스트
package scanner

import "testing"

func TestGoTypeFormat_TimeCov(t *testing.T) {
	if got := goTypeFormat("time.Time", Field{}); got != "date-time" {
		t.Fatalf("expected date-time, got %s", got)
	}
}
