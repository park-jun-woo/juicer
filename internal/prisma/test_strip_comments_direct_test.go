//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what stripComments 라인/후행 주석 제거 직접 테스트
package prisma

import "testing"

func TestStripCommentsDirect(t *testing.T) {
	src := "model User { // header\nid Int @id\n// full line\n/// triple"
	got := stripComments(src)
	want := "model User { \nid Int @id\n\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	// no comment line preserved verbatim
	if stripComments("plain") != "plain" {
		t.Error("plain line changed")
	}
}
