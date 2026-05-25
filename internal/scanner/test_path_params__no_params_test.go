//ff:func feature=scan type=extract control=sequence
//ff:what TestPathParams_NoParams 테스트
package scanner

import "testing"

func TestPathParams_NoParams(t *testing.T) {
	params := pathParams("/api/users")
	if len(params) != 0 {
		t.Fatalf("expected 0, got %d", len(params))
	}
}
