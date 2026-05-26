//ff:func feature=scan type=test control=sequence
//ff:what TestGoTypeFormat_UUID uuid.UUID format 테스트
package scanner

import "testing"

func TestGoTypeFormat_UUID(t *testing.T) {
	if got := goTypeFormat("uuid.UUID", Field{}); got != "uuid" {
		t.Fatalf("goTypeFormat(uuid.UUID)=%q, want uuid", got)
	}
}
