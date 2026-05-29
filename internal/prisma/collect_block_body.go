//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what 시작 라인부터 중괄호 균형으로 블록 본문 라인들과 다음 인덱스 반환
package prisma

import "strings"

// collectBlockBody returns non-empty body lines from start until the matching
// closing brace, plus the index of the line after the block.
func collectBlockBody(lines []string, start int) ([]string, int) {
	body := make([]string, 0, 16)
	depth := 1
	i := start
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		depth += strings.Count(line, "{") - strings.Count(line, "}")
		if depth <= 0 {
			return body, i + 1
		}
		if line != "" {
			body = append(body, line)
		}
		i++
	}
	return body, i
}
