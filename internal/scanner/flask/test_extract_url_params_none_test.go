//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractURLParams_None 테스트
package flask

import "testing"

func TestExtractURLParams_None(t *testing.T) {
	if got := extractURLParams("/static/path"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}
