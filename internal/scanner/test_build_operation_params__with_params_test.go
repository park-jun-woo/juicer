//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildOperationParams_WithParams 테스트
package scanner

import "testing"

func TestBuildOperationParams_WithParams(t *testing.T) {
	req := &Request{
		PathParams: []Param{{Name: "id"}},
		Query:      []Param{{Name: "page"}},
	}
	result := buildOperationParams(req)
	if len(result) != 2 {
		t.Fatalf("expected 2, got %d", len(result))
	}
}
