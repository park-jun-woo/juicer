//ff:func feature=ddl type=parse control=sequence
//ff:what 블록 주석 내부 줄 소비 — */ 발견 시 나머지 반환
package ddl

import "strings"

// consumeBlockCommentLine processes a line inside a block comment.
// Returns the remaining lines and whether we are still inside the block.
func consumeBlockCommentLine(lines []string, trimmed string) ([]string, bool) {
	idx := strings.Index(trimmed, "*/")
	if idx < 0 {
		return lines[1:], true
	}
	rest := strings.TrimSpace(trimmed[idx+2:])
	if rest == "" {
		return lines[1:], false
	}
	lines[0] = rest
	return lines, false
}
