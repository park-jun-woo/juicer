//ff:func feature=scan type=test control=sequence
//ff:what TestBuildOperationParams_PathAndQuery 테스트
package scanner

import "testing"

func TestBuildOperationParams_PathAndQuery(t *testing.T) {
	req := &Request{
		PathParams: []Param{{Name: "id", Type: "string"}},
		Query:      []Param{{Name: "limit", Type: "int", Default: "10"}},
	}
	params := buildOperationParams(req)
	if len(params) != 2 {
		t.Fatalf("expected 2, got %d", len(params))
	}
}
