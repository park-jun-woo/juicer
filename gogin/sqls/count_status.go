//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what 지정된 상태의 메서드 수 카운트
package sqls

// countStatus counts methods with the given status.
//
func countStatus(sess *Session, status string) int {
	n := 0
	for _, m := range sess.Methods {
		if m.Status == status {
			n++
		}
	}
	return n
}

