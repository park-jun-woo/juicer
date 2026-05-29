//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what 모델명을 @@map 반영 테이블명(비인용 raw)으로 변환 (없으면 모델명)
package prisma

// rawTableName returns the unquoted table name for a model. Used where the raw
// name is needed before composing (e.g. index name generation) — distinct from
// tableNameOf which returns the quoted form for direct emission.
func rawTableName(modelName string, s schema) string {
	if t, ok := s.tableNames[modelName]; ok && t != "" {
		return t
	}
	return modelName
}
