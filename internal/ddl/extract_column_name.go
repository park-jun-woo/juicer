//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what 컬럼 정의에서 컬럼명 추출
package ddl

import "strings"

// extractColumnName returns the first token from a column definition line.
func extractColumnName(line string) string {
	line = strings.TrimSpace(line)
	// Remove leading line comment if present
	for strings.HasPrefix(line, "--") {
		nl := strings.IndexByte(line, '\n')
		if nl < 0 {
			return ""
		}
		line = strings.TrimSpace(line[nl+1:])
	}
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return ""
	}
	return strings.ToLower(fields[0])
}
