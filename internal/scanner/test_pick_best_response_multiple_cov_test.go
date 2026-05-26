//ff:func feature=scan type=test control=sequence
//ff:what TestPickBestResponse_MultipleCov 테스트
package scanner

import "testing"

func TestPickBestResponse_MultipleCov(t *testing.T) {
	resps := []Response{
		{Status: "200", Kind: "json"},
		{Status: "200", Kind: "json", TypeName: "User", Fields: []Field{{Name: "id"}}},
	}
	best := pickBestResponse(resps)
	if best.TypeName != "User" {
		t.Fatal("expected richer response")
	}
}
