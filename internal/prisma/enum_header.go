//ff:func feature=prisma type=parse control=sequence topic=prisma
//ff:what 한 라인이 enum <Name> { 헤더인지 판정하고 enum명 반환
package prisma

import "strings"

// enumHeader reports whether the line opens an `enum <Name> {` block.
func enumHeader(line string) (string, bool) {
	fields := strings.Fields(strings.TrimSpace(line))
	if len(fields) < 3 || fields[0] != "enum" {
		return "", false
	}
	if fields[2] != "{" {
		return "", false
	}
	return fields[1], true
}
