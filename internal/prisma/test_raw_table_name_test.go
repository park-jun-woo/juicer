//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what rawTableName 비인용 테이블명/폴백 테스트
package prisma

import "testing"

func TestRawTableName(t *testing.T) {
	s := schema{tableNames: map[string]string{"User": "users", "Empty": ""}}
	if got := rawTableName("User", s); got != "users" {
		t.Errorf("mapped: got %q, want users", got)
	}
	if got := rawTableName("Post", s); got != "Post" {
		t.Errorf("fallback: got %q, want Post", got)
	}
	if got := rawTableName("Empty", s); got != "Empty" {
		t.Errorf("empty mapping falls back: got %q, want Empty", got)
	}
}
