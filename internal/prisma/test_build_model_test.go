//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what buildModel 본문 라인 → 필드/블록속성 조립 테스트
package prisma

import "testing"

func TestBuildModel(t *testing.T) {
	m := buildModel("User", []string{
		"id Int @id",
		"email String",
		`@@map("users")`,
	})
	if m.name != "User" || len(m.fields) != 2 {
		t.Fatalf("got %+v", m)
	}
	if m.tableName != "users" {
		t.Errorf("@@map tableName: %q", m.tableName)
	}
	if len(m.blockAttrs) != 1 {
		t.Errorf("blockAttrs: %v", m.blockAttrs)
	}
}
