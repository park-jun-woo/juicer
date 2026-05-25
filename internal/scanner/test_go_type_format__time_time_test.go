//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeFormat_TimeTime 테스트
package scanner

import "testing"

func TestGoTypeFormat_TimeTime(t *testing.T) {
	got := goTypeFormat("time.Time", Field{})
	if got != "date-time" {
		t.Fatalf("expected date-time, got %s", got)
	}
}
