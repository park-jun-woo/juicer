//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what FK REFERENCES 대상명을 테이블 맵 키로 해석 (정확 일치 → 인용제거+대소문자무시 보정)
package ddl

import "strings"

// resolveRefKey extracts the REFERENCES target from a constraint and maps it to
// the matching table-map key: exact match first, then a quote-stripped /
// case-insensitive fallback. Returns "" when the target is not in the map.
func resolveRefKey(tables map[string]*Table, constraint string) string {
	m := reFKReference.FindStringSubmatch(constraint)
	if m == nil {
		return ""
	}
	raw := m[1]
	if _, ok := tables[raw]; ok {
		return raw
	}
	want := strings.ToLower(cleanTableName(raw))
	for key := range tables {
		if strings.ToLower(cleanTableName(key)) == want {
			return key
		}
	}
	return ""
}
