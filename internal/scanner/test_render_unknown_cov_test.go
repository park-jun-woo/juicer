//ff:func feature=scan type=test control=sequence
//ff:what TestRender_UnknownCov 테스트
package scanner

import "testing"

func TestRender_UnknownCov(t *testing.T) {
	_, err := Render(&ScanResult{}, Format(99))
	if err == nil {
		t.Fatal("expected error for unknown format")
	}
}
