//ff:func feature=scan type=parse control=iteration dimension=1 topic=fastapi
//ff:what 모듈 경로에서 선행 점 개수를 센다
package fastapi

// countLeadingDots counts leading dots in a module path.
func countLeadingDots(module string) int {
	dots := 0
	for _, ch := range module {
		if ch == '.' {
			dots++
		} else {
			break
		}
	}
	return dots
}
