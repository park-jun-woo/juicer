//ff:func feature=scan type=test control=sequence
//ff:what TestPickBestResponse_Single 테스트
package scanner

import "testing"

func TestPickBestResponse_Single(t *testing.T) {
	resps := []Response{{Status: "200"}}
	best := pickBestResponse(resps)
	if best.Status != "200" {
		t.Fatal("expected 200")
	}
}

