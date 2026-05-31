//ff:func feature=prisma type=test control=sequence topic=prisma
//ff:what applyBlockAttr 블록 속성 기록 및 @@map 테이블명 갱신 테스트
package prisma

import "testing"

func TestApplyBlockAttr(t *testing.T) {
	m := &model{name: "User"}
	applyBlockAttr(m, "@@index([a])")
	if len(m.blockAttrs) != 1 || m.tableName != "" {
		t.Errorf("non-map: blockAttrs=%v tableName=%q", m.blockAttrs, m.tableName)
	}
	applyBlockAttr(m, `@@map("users")`)
	if len(m.blockAttrs) != 2 || m.tableName != "users" {
		t.Errorf("@@map: blockAttrs=%v tableName=%q", m.blockAttrs, m.tableName)
	}
}
