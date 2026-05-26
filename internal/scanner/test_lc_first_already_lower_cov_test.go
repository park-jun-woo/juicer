//ff:func feature=scan type=test control=sequence
//ff:what TestLcFirst_AlreadyLowerCov 테스트
package scanner

import "testing"

func TestLcFirst_AlreadyLowerCov(t *testing.T) {
	if lcFirst("building") != "building" {
		t.Fatal("expected building")
	}
}
