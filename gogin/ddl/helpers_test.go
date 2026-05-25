//ff:func feature=ddl type=extract control=iteration dimension=1
//ff:what 테스트 헬퍼 — 문자열 위치 검색
package ddl

func indexOfStr(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
