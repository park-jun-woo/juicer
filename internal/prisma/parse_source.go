//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what Prisma 소스 문자열을 주석제거 후 model 목록으로 파싱
package prisma

// parseSource strips comments and parses all model blocks in one source string.
func parseSource(src string) []model {
	blocks := findModelBlocks(stripComments(src))
	models := make([]model, 0, len(blocks))
	for name, body := range blocks {
		models = append(models, buildModel(name, body))
	}
	return models
}
