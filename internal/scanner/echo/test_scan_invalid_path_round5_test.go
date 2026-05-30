//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestScan_InvalidPath_Round5 테스트
package echo

import "testing"

func TestScan_InvalidPath_Round5(t *testing.T) {

	if _, err := Scan(string([]byte{0})); err == nil {

	}
}
