//ff:func feature=ddl type=parse control=sequence
//ff:what 테이블명에서 스키마 접두사와 따옴표 제거
package ddl

import "strings"

// cleanTableName normalises a captured table name:
//  1. Strips schema prefix: public.profiles -> profiles
//  2. Strips double quotes: "UserEvents" -> UserEvents
//  3. Handles both: public."UserEvents" -> UserEvents
func cleanTableName(raw string) string {
	// Take only the part after the last dot (remove schema)
	if idx := strings.LastIndex(raw, "."); idx >= 0 {
		raw = raw[idx+1:]
	}
	// Strip surrounding double quotes
	raw = strings.Trim(raw, `"`)
	return raw
}
