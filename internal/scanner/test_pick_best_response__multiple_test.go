//ff:func feature=scan type=extract control=sequence
//ff:what TestPickBestResponse_Multiple 테스트
package scanner

import "testing"

func TestPickBestResponse_Multiple(t *testing.T) {
	resps := []Response{
		{Status: "200"},
		{Status: "200", TypeName: "User", Fields: []Field{{Name: "ID"}}},
	}
	best := pickBestResponse(resps)
	if best.TypeName != "User" {
		t.Fatal("expected User (richer response)")
	}
}
