//ff:func feature=scan type=test control=sequence
//ff:what TestBuildRequestBody_FilesBranch 테스트
package scanner

import "testing"

func TestBuildRequestBody_FilesBranch(t *testing.T) {
	req := &Request{Files: []Param{{Name: "avatar", Type: "file"}}}
	result := buildRequestBody(req, map[string]any{})
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
