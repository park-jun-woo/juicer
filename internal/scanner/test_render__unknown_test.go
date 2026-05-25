//ff:func feature=scan type=extract control=sequence
//ff:what TestRender_Unknown 테스트
package scanner

import "testing"

func TestRender_Unknown(t *testing.T) {
	result := &ScanResult{}
	_, err := Render(result, Format(99))
	if err == nil {
		t.Fatal("expected error")
	}
}
