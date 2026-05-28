//ff:func feature=ddl type=parse control=sequence
//ff:what /* 로 시작하는 줄에서 블록 주석 열기 — 인라인 */ 처리 포함
package ddl

import "strings"

// openBlockComment handles a line starting with /*. If */ appears on the same
// line, it removes the inline block comment and returns the remainder.
// Otherwise it enters block-comment mode.
func openBlockComment(lines []string, trimmed string) ([]string, bool) {
	idx := strings.Index(trimmed[2:], "*/")
	if idx < 0 {
		return lines[1:], true
	}
	rest := strings.TrimSpace(trimmed[2+idx+2:])
	if rest == "" {
		return lines[1:], false
	}
	lines[0] = rest
	return lines, false
}
