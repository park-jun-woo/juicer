//ff:func feature=scan type=test control=sequence
//ff:what TestGoTypeFormat_Float32Cov 테스트
package scanner

import "testing"

func TestGoTypeFormat_Float32Cov(t *testing.T) {
	if got := goTypeFormat("float32", Field{}); got != "float" {
		t.Fatalf("expected float, got %s", got)
	}
}
