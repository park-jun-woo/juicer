//ff:func feature=prisma type=test control=sequence topic=prisma
//ff:what buildConstraints PK/UNIQUE/FK 라인 조립 직접호출 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestBuildConstraintsDirect(t *testing.T) {
	s := schema{
		tableNames: map[string]string{"Org": "Org", "Child": "Child"},
		columns: map[string]map[string]string{
			"Child": {"id": "id", "orgId": "orgId", "email": "email"},
			"Org":   {"id": "id"},
		},
	}
	m := model{name: "Child", fields: []field{
		{name: "id", attrs: []string{"@id"}},
		{name: "email", attrs: []string{"@unique"}},
		{name: "org", baseType: "Org", attrs: []string{"@relation(fields: [orgId], references: [id])"}},
	}}
	got := buildConstraints(m, s)
	want := []string{
		`PRIMARY KEY ("id")`,
		`UNIQUE ("email")`,
		`FOREIGN KEY ("orgId") REFERENCES "Org" ("id")`,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
