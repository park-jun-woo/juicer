//ff:func feature=scan type=test control=sequence topic=express
//ff:what 경로 파라미터 없음 테스트
package express

import "testing"

func TestExtractPathParams_NoParams(t *testing.T) {
	params := extractPathParams("/users")
	if len(params) != 0 {
		t.Fatalf("expected 0 params, got %d", len(params))
	}
}
