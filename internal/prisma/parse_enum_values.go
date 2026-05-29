//ff:func feature=prisma type=parse control=iteration dimension=1 topic=prisma
//ff:what enum 본문 라인들에서 값 이름(맨 앞 토큰) 추출 (@map/빈 줄 무시, v1은 값 이름만)
package prisma

import "strings"

// parseEnumValues extracts enum value names (leading token) from body lines.
// Empty lines and attribute-only lines are skipped; per-value @map is ignored.
func parseEnumValues(body []string) []string {
	values := make([]string, 0, len(body))
	for _, line := range body {
		fields := strings.Fields(strings.TrimSpace(line))
		if len(fields) == 0 {
			continue
		}
		name := fields[0]
		if strings.HasPrefix(name, "@") {
			continue
		}
		values = append(values, name)
	}
	return values
}
