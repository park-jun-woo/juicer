//ff:func feature=scan type=test control=sequence topic=express
//ff:what 복수 경로 파라미터 추출 테스트
package express

import "testing"

func TestExtractPathParams_Multiple(t *testing.T) {
	params := extractPathParams("/users/:userId/posts/:postId")
	if len(params) != 2 {
		t.Fatalf("expected 2 params, got %d", len(params))
	}
	if params[0] != "userId" {
		t.Errorf("param[0]: want userId, got %s", params[0])
	}
	if params[1] != "postId" {
		t.Errorf("param[1]: want postId, got %s", params[1])
	}
}
