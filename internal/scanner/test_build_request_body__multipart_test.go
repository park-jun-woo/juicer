//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildRequestBody_Multipart 테스트
package scanner

import "testing"

func TestBuildRequestBody_Multipart(t *testing.T) {
	req := &Request{Files: []Param{{Name: "file"}}}
	schemas := map[string]any{}
	result := buildRequestBody(req, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
