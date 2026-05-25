//ff:func feature=scan type=extract control=sequence
//ff:what TestRender_OpenAPI 테스트
package scanner

import "testing"

func TestRender_OpenAPI(t *testing.T) {
	result := &ScanResult{}
	data, err := Render(result, FormatOpenAPI)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}
