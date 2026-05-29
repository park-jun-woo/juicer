//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 파싱된 모델들로부터 교차 참조용 schema(모델집합/enum집합/테이블명/컬럼맵) 구성
package prisma

import "github.com/park-jun-woo/codistill/internal/ddl"

// buildSchema builds cross-model lookups from parsed models for FK/key
// resolution (model set, enum set, table names, field->column maps).
func buildSchema(models []model, enums []ddl.EnumType) schema {
	s := schema{
		models:     make(map[string]bool, len(models)),
		enums:      make(map[string]bool, len(enums)),
		tableNames: make(map[string]string, len(models)),
		columns:    make(map[string]map[string]string, len(models)),
	}
	for _, m := range models {
		s.models[m.name] = true
		s.tableNames[m.name] = m.tableName
		s.columns[m.name] = fieldColumnMap(m)
	}
	for _, e := range enums {
		s.enums[e.Name] = true
	}
	return s
}
