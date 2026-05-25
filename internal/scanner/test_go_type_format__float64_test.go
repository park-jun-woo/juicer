//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeFormat_Float64 테스트
package scanner

import "testing"

func TestGoTypeFormat_Float64(t *testing.T) {
	got := goTypeFormat("float64", Field{})
	if got != "double" {
		t.Fatalf("expected double, got %s", got)
	}
}
