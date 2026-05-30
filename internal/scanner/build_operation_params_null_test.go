//ff:func feature=scan type=test control=sequence
//ff:what buildOperationParams — DefaultIsNull 쿼리 파라미터 분기를 검증
package scanner

import "testing"

func TestBuildOperationParams_QueryDefaultNull(t *testing.T) {
	req := &Request{
		Query: []Param{{Name: "filter", Type: "string", DefaultIsNull: true}},
	}
	params := buildOperationParams(req)
	if len(params) != 1 {
		t.Fatalf("expected 1 param, got %d", len(params))
	}
	schema := params[0]["schema"].(map[string]any)
	if schema["nullable"] != true {
		t.Errorf("expected nullable true, got %v", schema["nullable"])
	}
	if v, ok := schema["default"]; !ok || v != nil {
		t.Errorf("expected default nil present, got %v ok=%v", v, ok)
	}
}
