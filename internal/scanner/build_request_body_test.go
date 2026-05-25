//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildRequestBody_JSON 테스트
package scanner

import "testing"

func TestBuildRequestBody_JSON(t *testing.T) {
	req := &Request{Body: &Body{TypeName: "User", Fields: []Field{{Name: "name", JSON: "name"}}}}
	schemas := map[string]any{}
	result := buildRequestBody(req, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
