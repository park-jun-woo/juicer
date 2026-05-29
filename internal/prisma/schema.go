//ff:type feature=prisma type=model topic=prisma
//ff:what 제약 해석에 필요한 모델 교차참조 룩업(모델집합/테이블명/컬럼맵)
package prisma

// schema holds cross-model lookups needed to resolve constraints.
type schema struct {
	models     map[string]bool              // declared model names
	enums      map[string]bool              // declared enum names
	tableNames map[string]string            // model name -> table name
	columns    map[string]map[string]string // model name -> field name -> column name
}
