//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what model <Name> { ... } 블록명+본문 라인 추출 (generator/datasource/enum 무시)
package prisma

import "strings"

// findModelBlocks scans stripped source and returns model name -> body lines.
// Non-model top-level blocks (generator/datasource/enum) are skipped.
func findModelBlocks(src string) map[string][]string {
	blocks := make(map[string][]string)
	lines := strings.Split(src, "\n")
	i := 0
	for i < len(lines) {
		name, ok := modelHeader(lines[i])
		if !ok {
			i++
			continue
		}
		body, next := collectBlockBody(lines, i+1)
		blocks[name] = body
		i = next
	}
	return blocks
}
