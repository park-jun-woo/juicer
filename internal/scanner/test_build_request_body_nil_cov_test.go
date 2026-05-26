//ff:func feature=scan type=test control=sequence
//ff:what TestBuildRequestBody_NilCov 테스트
package scanner

import "testing"

func TestBuildRequestBody_NilCov(t *testing.T) {
	req := &Request{}
	result := buildRequestBody(req, map[string]any{})
	if result != nil {
		t.Fatal("expected nil")
	}
}
