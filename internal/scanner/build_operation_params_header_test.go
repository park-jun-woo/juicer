//ff:func feature=scan type=test control=sequence
//ff:what TestBuildOperationParams_Header — header 파라미터 생성 테스트
package scanner

import "testing"

func TestBuildOperationParams_Header(t *testing.T) {
	req := &Request{
		Headers: []Param{{Name: "Authorization", Type: "string"}},
	}
	params := buildOperationParams(req)
	if len(params) != 1 {
		t.Fatalf("expected 1 param, got %d", len(params))
	}
	if params[0]["in"] != "header" {
		t.Errorf("param in: want header, got %v", params[0]["in"])
	}
	if params[0]["name"] != "Authorization" {
		t.Errorf("param name: want Authorization, got %v", params[0]["name"])
	}
	schema := params[0]["schema"].(map[string]any)
	if schema["type"] != "string" {
		t.Errorf("schema type: want string, got %v", schema["type"])
	}
}
