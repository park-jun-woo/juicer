//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what hasAttr 정확 매칭 및 괄호 접두 매칭 테스트
package prisma

import "testing"

func TestHasAttr(t *testing.T) {
	attrs := []string{"@id", "@default(now())", "@map2"}
	if !hasAttr(attrs, "@id") {
		t.Error("exact @id should match")
	}
	if !hasAttr(attrs, "@default") {
		t.Error("@default(...) prefix should match")
	}
	if hasAttr(attrs, "@unique") {
		t.Error("absent @unique should not match")
	}
	if hasAttr(attrs, "@map") {
		t.Error("@map2 must not match @map (no paren)")
	}
}
