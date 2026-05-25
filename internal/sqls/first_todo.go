//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what 첫 번째 TODO 항목 인덱스 반환 (-1 if none)
package sqls

// firstTODO returns the index of the first TODO method, or -1.
//
func firstTODO(sess *Session) int {
	for i, m := range sess.Methods {
		if m.Status == "TODO" {
			return i
		}
	}
	return -1
}

