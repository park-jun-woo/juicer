//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what Prisma 소스의 // /// 라인/말미 주석 제거
package prisma

import "strings"

// stripComments removes // and /// line and trailing comments from Prisma source.
func stripComments(src string) string {
	lines := strings.Split(src, "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		if idx := strings.Index(line, "//"); idx >= 0 {
			line = line[:idx]
		}
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}
