//ff:func feature=scan type=test control=sequence
//ff:what TestGoTypeFormat_Int64 테스트
package scanner

import "testing"

func TestGoTypeFormat_Int64(t *testing.T) {
	got := goTypeFormat("int64", Field{})
	if got != "int64" {
		t.Fatalf("expected int64, got %s", got)
	}
}
