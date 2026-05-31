//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what relationAttr @relation(...) 내부 인자 추출 테스트
package prisma

import "testing"

func TestRelationAttr(t *testing.T) {
	attrs := []string{"@id", "@relation(fields: [aId], references: [id])"}
	v, ok := relationAttr(attrs)
	if !ok || v != "fields: [aId], references: [id]" {
		t.Errorf("got (%q,%v)", v, ok)
	}
	if _, ok := relationAttr([]string{"@id"}); ok {
		t.Error("no relation must be false")
	}
	// no closing paren: keep remainder
	v, ok = relationAttr([]string{"@relation(fields: [x]"})
	if !ok || v != "fields: [x]" {
		t.Errorf("no close paren: got (%q,%v)", v, ok)
	}
}
