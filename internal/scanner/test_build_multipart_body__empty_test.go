//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildMultipartBody_Empty 테스트
package scanner

import "testing"

func TestBuildMultipartBody_Empty(t *testing.T) {
	req := &Request{}
	result := buildMultipartBody(req)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
