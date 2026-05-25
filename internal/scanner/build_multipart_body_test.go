//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildMultipartBody 테스트
package scanner

import "testing"

func TestBuildMultipartBody(t *testing.T) {
	req := &Request{
		FormFields: []Param{{Name: "title"}},
		Files:      []Param{{Name: "file"}},
	}
	result := buildMultipartBody(req)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
