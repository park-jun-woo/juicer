//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what buildSchema 모델/enum 교차 룩업 구성 테스트
package prisma

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/ddl"
)

func TestBuildSchema(t *testing.T) {
	models := []model{{
		name:      "User",
		tableName: "users",
		fields:    []field{{name: "id"}, {name: "orgId", attrs: []string{`@map("org_id")`}}},
	}}
	enums := []ddl.EnumType{{Name: "Role"}}
	s := buildSchema(models, enums)
	if !s.models["User"] || s.tableNames["User"] != "users" {
		t.Errorf("models/tableNames: %+v", s)
	}
	if s.columns["User"]["orgId"] != "org_id" {
		t.Errorf("columns: %v", s.columns["User"])
	}
	if !s.enums["Role"] {
		t.Errorf("enums: %v", s.enums)
	}
}
