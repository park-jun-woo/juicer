//ff:func feature=scan type=test control=sequence
//ff:what TestRender_OpenAPICov 테스트
package scanner

import "testing"

func TestRender_OpenAPICov(t *testing.T) {
	result := &ScanResult{}
	_, err := Render(result, FormatOpenAPI)
	if err != nil {
		t.Fatal(err)
	}
}
