//ff:func feature=ddl type=extract control=sequence
//ff:what 테스트 헬퍼 — 문자열 포함 여부 확인
package ddl

func contains(s, substr string) bool {
	return len(s) >= len(substr) && containsStr(s, substr)
}
