//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what 경로 파라미터 없는 경로 테스트
package fastify

import "testing"

func TestExtractPathParams_None(t *testing.T) {
	params := extractPathParams("/users")
	if len(params) != 0 {
		t.Errorf("expected 0 params, got %d", len(params))
	}
}
