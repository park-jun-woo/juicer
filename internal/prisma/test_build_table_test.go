//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what buildTable 모델 → ddl.Table 변환(관계 필드 스킵) 테스트
package prisma

import "testing"

func TestBuildTable(t *testing.T) {
	s := schema{
		models:     map[string]bool{"User": true, "Org": true},
		tableNames: map[string]string{"User": "users"},
		columns:    map[string]map[string]string{"User": {"id": "id"}},
	}
	m := model{name: "User", fields: []field{
		{name: "id", baseType: "Int", attrs: []string{"@id"}},
		{name: "org", baseType: "Org"}, // relation field, skipped
	}}
	tbl := buildTable(m, s)
	if tbl.Name != `"users"` {
		t.Errorf("name: %q", tbl.Name)
	}
	if len(tbl.Columns) != 1 || tbl.Columns[0].Name != "id" {
		t.Errorf("columns (relation must be skipped): %+v", tbl.Columns)
	}
	if len(tbl.Constraints) != 1 {
		t.Errorf("constraints: %v", tbl.Constraints)
	}
}
