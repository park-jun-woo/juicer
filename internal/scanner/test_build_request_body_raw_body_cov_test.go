//ff:func feature=scan type=test control=sequence
//ff:what TestBuildRequestBody_RawBodyCov 테스트
package scanner

import "testing"

func TestBuildRequestBody_RawBodyCov(t *testing.T) {
	req := &Request{RawBody: true}
	result := buildRequestBody(req, map[string]any{})
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
