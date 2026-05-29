//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what enum <Name> { ... } 블록명+본문 라인 추출 (model/generator/datasource 무시)
package prisma

import "strings"

// findEnumBlocks scans stripped source and returns enum name -> body lines.
// Non-enum top-level blocks (model/generator/datasource) are skipped.
func findEnumBlocks(src string) map[string][]string {
	blocks := make(map[string][]string)
	lines := strings.Split(src, "\n")
	i := 0
	for i < len(lines) {
		name, ok := enumHeader(lines[i])
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
