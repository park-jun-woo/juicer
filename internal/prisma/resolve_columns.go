//ff:func feature=prisma type=convert control=iteration dimension=1 topic=prisma
//ff:what 모델의 필드명 목록을 실제 컬럼명(@map 반영)으로 해석 후 인용된 목록 반환
package prisma

// resolveColumns maps field names to their quoted column names within a model,
// honoring @map overrides; unknown fields fall back to their own name. The
// returned names are quoted for direct emission in key/FK lines.
func resolveColumns(modelName string, fields []string, s schema) []string {
	raw := resolveRawColumns(modelName, fields, s)
	cols := make([]string, 0, len(raw))
	for _, c := range raw {
		cols = append(cols, quoteIdent(c))
	}
	return cols
}
