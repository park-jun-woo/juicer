//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what 모델명을 @@map 반영 테이블명으로 변환 (없으면 모델명)
package prisma

// tableNameOf returns the table name for a model, falling back to the model
// name when no @@map override was recorded.
func tableNameOf(modelName string, s schema) string {
	if t, ok := s.tableNames[modelName]; ok && t != "" {
		return t
	}
	return modelName
}
