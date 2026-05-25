//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildRequestBody_NoBody 테스트
package scanner

import "testing"

func TestBuildRequestBody_NoBody(t *testing.T) {
	req := &Request{}
	schemas := map[string]any{}
	result := buildRequestBody(req, schemas)
	if result != nil {
		t.Fatal("expected nil")
	}
}
