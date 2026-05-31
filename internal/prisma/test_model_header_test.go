//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what modelHeader `model <Name> {` 인식 테스트
package prisma

import "testing"

func TestModelHeader(t *testing.T) {
	if n, ok := modelHeader("  model User {  "); !ok || n != "User" {
		t.Errorf("got (%q,%v)", n, ok)
	}
	if _, ok := modelHeader("enum Role {"); ok {
		t.Error("enum is not model header")
	}
	if _, ok := modelHeader("model User"); ok {
		t.Error("missing brace must be false")
	}
	if _, ok := modelHeader("model User ("); ok {
		t.Error("non-brace third token must be false")
	}
}
