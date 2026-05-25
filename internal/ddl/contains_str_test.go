//ff:func feature=ddl type=extract control=iteration dimension=1
//ff:what 테스트 헬퍼 — 문자열 포함 여부 반복 검색
package ddl

func containsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
