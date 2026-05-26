//ff:func feature=scan type=test control=sequence
//ff:what TestBuildOperationParams_QueryNoDefaultCov 테스트
package scanner

import "testing"

func TestBuildOperationParams_QueryNoDefaultCov(t *testing.T) {
	req := &Request{
		Query: []Param{{Name: "q", Type: "string"}},
	}
	result := buildOperationParams(req)
	if len(result) != 1 {
		t.Fatalf("expected 1 param, got %d", len(result))
	}
}
