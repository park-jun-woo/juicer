//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 데코레이터 이름에서 인증 레벨을 매칭한다
package nestjs

import "strings"

// matchAuthLevel returns the auth level for a decorator name.
// Returns "" if no pattern matches.
func matchAuthLevel(decoratorName string) string {
	for _, p := range authLevelPatterns {
		if strings.Contains(decoratorName, p.contains) {
			return p.level
		}
	}
	return ""
}
