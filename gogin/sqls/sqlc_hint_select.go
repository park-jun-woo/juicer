//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what SELECT 메서드의 sqlc 힌트를 반환 타입으로 결정한다
package sqls

import "strings"

// sqlcHintSelect determines :one or :many for SELECT methods.
func sqlcHintSelect(sk *MethodSkeleton) string {
	for _, r := range sk.Returns {
		if strings.HasPrefix(r, "[]") {
			return ":many"
		}
	}
	return ":one"
}
