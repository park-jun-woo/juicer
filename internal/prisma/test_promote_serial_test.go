//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what promoteSerial integer/bigint autoincrement → serial 승격 테스트
package prisma

import "testing"

func TestPromoteSerial(t *testing.T) {
	if got := promoteSerial("integer", "", true); got != "serial" {
		t.Errorf("integer autoincrement: got %q, want serial", got)
	}
	if got := promoteSerial("bigint", "", true); got != "bigserial" {
		t.Errorf("bigint autoincrement: got %q, want bigserial", got)
	}
	if got := promoteSerial("text", "", true); got != "text" {
		t.Errorf("non-int unchanged: got %q", got)
	}
	if got := promoteSerial("integer", "", false); got != "integer" {
		t.Errorf("no default unchanged: got %q", got)
	}
	if got := promoteSerial("integer", "5", true); got != "integer" {
		t.Errorf("non-empty default unchanged: got %q", got)
	}
}
