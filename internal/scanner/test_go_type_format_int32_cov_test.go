//ff:func feature=scan type=test control=sequence
//ff:what TestGoTypeFormat_Int32Cov 테스트
package scanner

import "testing"

func TestGoTypeFormat_Int32Cov(t *testing.T) {
	if got := goTypeFormat("int32", Field{}); got != "int32" {
		t.Fatalf("expected int32, got %s", got)
	}
}
