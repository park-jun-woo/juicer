//ff:func feature=prisma type=test control=sequence topic=prisma
//ff:what indexName 결정적 인덱스명 생성 테스트
package prisma

import "testing"

func TestIndexName(t *testing.T) {
	if got := indexName("users", []string{"org_id", "email"}); got != "users_org_id_email_idx" {
		t.Errorf("got %q", got)
	}
	if got := indexName("t", nil); got != "t_idx" {
		t.Errorf("no cols: got %q", got)
	}
}
