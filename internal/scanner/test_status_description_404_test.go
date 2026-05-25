//ff:func feature=scan type=extract control=sequence
//ff:what TestStatusDescription_404 테스트
package scanner

import "testing"

func TestStatusDescription_404(t *testing.T) {
	if statusDescription("404") != "Not Found" {
		t.Fatal("expected Not Found")
	}
}
