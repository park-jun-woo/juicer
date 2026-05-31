//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what foreignKeyConstraints @relation 소유측 FK 라인 수집 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestForeignKeyConstraints(t *testing.T) {
	s := schema{
		tableNames: map[string]string{"Org": "Org", "Child": "Child"},
		columns: map[string]map[string]string{
			"Child": {"orgId": "orgId"},
			"Org":   {"id": "id"},
		},
	}
	m := model{name: "Child", fields: []field{
		{name: "id"},
		{name: "org", baseType: "Org", attrs: []string{"@relation(fields: [orgId], references: [id])"}},
		{name: "noop", baseType: "Org", attrs: []string{"@relation(references: [id])"}},
	}}
	got := foreignKeyConstraints(m, s)
	want := []string{`FOREIGN KEY ("orgId") REFERENCES "Org" ("id")`}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
