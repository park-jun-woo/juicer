//ff:func feature=scan type=test control=sequence
//ff:what TestBuildOperationParams_WithParamsCov 테스트
package scanner

import "testing"

func TestBuildOperationParams_WithParamsCov(t *testing.T) {
	req := &Request{
		PathParams: []Param{{Name: "id", Type: "int"}},
		Query:      []Param{{Name: "page", Type: "int", Default: "1"}},
	}
	result := buildOperationParams(req)
	if len(result) != 2 {
		t.Fatalf("expected 2 params, got %d", len(result))
	}
}
