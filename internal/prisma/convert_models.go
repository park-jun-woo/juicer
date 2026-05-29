//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 모델 목록을 테이블명->*ddl.Table 맵으로 변환 (2-pass: schema 구성 후 변환)
package prisma

import "github.com/park-jun-woo/codistill/internal/ddl"

// convertModels converts parsed models into a table-name -> *ddl.Table map,
// using enum names for enum-typed default-value quoting.
func convertModels(models []model, enums []ddl.EnumType) map[string]*ddl.Table {
	s := buildSchema(models, enums)
	tables := make(map[string]*ddl.Table, len(models))
	for _, m := range models {
		t := buildTable(m, s)
		tables[t.Name] = t
	}
	return tables
}
