//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 두 미들웨어 슬라이스를 중복 없이 병합한다
package laravel

// mergeMiddleware merges two middleware slices, avoiding duplicates.
func mergeMiddleware(a, b []string) []string {
	if len(b) == 0 {
		return a
	}
	seen := make(map[string]bool, len(a))
	for _, m := range a {
		seen[m] = true
	}
	result := make([]string, len(a))
	copy(result, a)
	for _, m := range b {
		if !seen[m] {
			result = append(result, m)
		}
	}
	return result
}
