//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what 모델명을 @@map 반영 테이블명으로 변환 후 인용 (없으면 모델명)
package prisma

// tableNameOf returns the quoted table name for a model, falling back to the
// model name when no @@map override was recorded. Quoting here makes every
// emit site (Table.Name, FK references) consistently quoted.
func tableNameOf(modelName string, s schema) string {
	if t, ok := s.tableNames[modelName]; ok && t != "" {
		return quoteIdent(t)
	}
	return quoteIdent(modelName)
}
