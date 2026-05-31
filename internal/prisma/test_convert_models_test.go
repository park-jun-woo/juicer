//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what convertModels 모델 슬라이스 → 테이블명→Table 맵 테스트
package prisma

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/ddl"
)

func TestConvertModels(t *testing.T) {
	models := []model{{
		name:      "User",
		tableName: "users",
		fields:    []field{{name: "id", baseType: "Int", attrs: []string{"@id"}}},
	}}
	tables := convertModels(models, []ddl.EnumType{})
	tbl, ok := tables[`"users"`]
	if !ok {
		t.Fatalf("table not found: keys=%v", tables)
	}
	if len(tbl.Columns) != 1 {
		t.Errorf("columns: %+v", tbl.Columns)
	}
}
