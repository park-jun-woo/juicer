//ff:func feature=sql type=parse control=selection
//ff:what CRUD + 반환 타입으로 sqlc 힌트 (:one/:many/:exec) 결정
package sqls

// sqlcHint returns the sqlc annotation hint (:one, :many, :exec) based on CRUD and return types.
func sqlcHint(sk *MethodSkeleton) string {
	switch sk.CRUD {
	case "SELECT":
		return sqlcHintSelect(sk)
	case "INSERT":
		return sqlcHintInsert(sk)
	default:
		return ":exec"
	}
}
