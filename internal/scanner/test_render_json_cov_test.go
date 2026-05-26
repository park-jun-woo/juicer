//ff:func feature=scan type=test control=sequence
//ff:what TestRender_JSONCov 테스트
package scanner

import "testing"

func TestRender_JSONCov(t *testing.T) {
	result := &ScanResult{}
	data, err := Render(result, FormatJSON)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}
