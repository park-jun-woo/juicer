//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what INSERT 메서드의 sqlc 힌트를 RETURNING 여부로 결정한다
package sqls

import "strings"

// sqlcHintInsert determines :one or :exec for INSERT methods.
func sqlcHintInsert(sk *MethodSkeleton) string {
	for _, frag := range sk.SQLFragments {
		if strings.Contains(strings.ToUpper(frag), "RETURNING") {
			return ":one"
		}
	}
	return ":exec"
}
