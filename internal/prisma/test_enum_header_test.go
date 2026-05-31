//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what enumHeader `enum <Name> {` 인식 테스트
package prisma

import "testing"

func TestEnumHeader(t *testing.T) {
	if n, ok := enumHeader("enum Role {"); !ok || n != "Role" {
		t.Errorf("got (%q,%v)", n, ok)
	}
	if _, ok := enumHeader("model User {"); ok {
		t.Error("model is not enum header")
	}
	if _, ok := enumHeader("enum Role"); ok {
		t.Error("missing brace must be false")
	}
	if _, ok := enumHeader("enum Role ["); ok {
		t.Error("non-brace third token must be false")
	}
}
