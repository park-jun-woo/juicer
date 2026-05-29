//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 모델의 필드명 목록을 실제 컬럼명(@map 반영) 목록으로 해석
package prisma

// resolveColumns maps field names to their column names within a model,
// honoring @map overrides; unknown fields fall back to their own name.
func resolveColumns(modelName string, fields []string, s schema) []string {
	cols := make([]string, 0, len(fields))
	lookup := s.columns[modelName]
	for _, f := range fields {
		if c, ok := lookup[f]; ok {
			cols = append(cols, c)
			continue
		}
		cols = append(cols, f)
	}
	return cols
}
