//ff:func feature=scan type=test control=sequence
//ff:what TestToOpenAPI_EmptyResult 테스트
package scanner

import "testing"

func TestToOpenAPI_EmptyResult(t *testing.T) {
	result := &ScanResult{}
	data, err := ToOpenAPI(result, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}

