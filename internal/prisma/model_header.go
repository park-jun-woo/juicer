//ff:func feature=prisma type=parse control=sequence topic=prisma
//ff:what 한 라인이 model <Name> { 헤더인지 판정하고 모델명 반환
package prisma

import "strings"

// modelHeader reports whether the line opens a `model <Name> {` block.
func modelHeader(line string) (string, bool) {
	fields := strings.Fields(strings.TrimSpace(line))
	if len(fields) < 3 || fields[0] != "model" {
		return "", false
	}
	if fields[2] != "{" {
		return "", false
	}
	return fields[1], true
}
