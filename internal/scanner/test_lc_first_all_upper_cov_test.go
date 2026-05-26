//ff:func feature=scan type=test control=sequence
//ff:what TestLcFirst_AllUpperCov 테스트
package scanner

import "testing"

func TestLcFirst_AllUpperCov(t *testing.T) {
	if lcFirst("ID") != "id" {
		t.Fatalf("expected id, got %s", lcFirst("ID"))
	}
}
