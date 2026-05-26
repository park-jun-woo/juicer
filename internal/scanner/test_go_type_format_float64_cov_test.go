//ff:func feature=scan type=test control=sequence
//ff:what TestGoTypeFormat_Float64Cov 테스트
package scanner

import "testing"

func TestGoTypeFormat_Float64Cov(t *testing.T) {
	if got := goTypeFormat("float64", Field{}); got != "double" {
		t.Fatalf("expected double, got %s", got)
	}
}
