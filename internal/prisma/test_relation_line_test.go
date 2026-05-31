//ff:func feature=prisma type=test control=sequence topic=prisma
//ff:what relationLine FOREIGN KEY 라인 생성 및 비소유측 빈 반환 테스트
package prisma

import "testing"

func TestRelationLine(t *testing.T) {
	s := schema{
		tableNames: map[string]string{"Org": "orgs", "Child": "children"},
		columns: map[string]map[string]string{
			"Child": {"orgId": "org_id"},
			"Org":   {"id": "id"},
		},
	}
	m := model{name: "Child"}
	f := field{name: "org", baseType: "Org"}
	rel := "fields: [orgId], references: [id], onDelete: Cascade"
	got := relationLine(m, f, rel, s)
	want := `FOREIGN KEY ("org_id") REFERENCES "orgs" ("id") ON DELETE CASCADE`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	// non-owning side: no fields
	if got := relationLine(m, f, "references: [id]", s); got != "" {
		t.Errorf("non-owning: got %q", got)
	}
}
