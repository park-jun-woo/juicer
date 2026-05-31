//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what tableNameOf @@map 반영 인용 테이블명/폴백 테스트
package prisma

import "testing"

func TestTableNameOf(t *testing.T) {
	s := schema{tableNames: map[string]string{"User": "users", "Empty": ""}}
	if got := tableNameOf("User", s); got != `"users"` {
		t.Errorf("mapped: got %q, want \"users\"", got)
	}
	if got := tableNameOf("Post", s); got != `"Post"` {
		t.Errorf("fallback: got %q, want \"Post\"", got)
	}
	if got := tableNameOf("Empty", s); got != `"Empty"` {
		t.Errorf("empty mapping falls back: got %q, want \"Empty\"", got)
	}
}
