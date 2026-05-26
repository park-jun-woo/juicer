//ff:func feature=ddl type=parse control=sequence
//ff:what Column.Raw에서 타입 부분(두 번째 토큰)을 새 타입으로 교체한다
package ddl

import "strings"

// changeColumnType replaces the type portion (second token) in the Raw string.
func changeColumnType(raw, newType string) string {
	fields := strings.Fields(raw)
	if len(fields) < 2 {
		return raw
	}
	fields[1] = newType
	return strings.Join(fields, " ")
}
