//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what 경로 파라미터 추출 테스트: /users/:id -> ["id"]
package fastify

import "testing"

func TestExtractPathParams(t *testing.T) {
	params := extractPathParams("/users/:id")
	if len(params) != 1 {
		t.Fatalf("expected 1 param, got %d", len(params))
	}
	if params[0] != "id" {
		t.Errorf("param: want id, got %s", params[0])
	}
}
