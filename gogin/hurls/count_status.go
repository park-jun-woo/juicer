//ff:func feature=hurl type=parse control=iteration dimension=1
//ff:what 지정된 상태의 엔드포인트 수 카운트
package hurls

// countStatus counts endpoints with the given status.
func countStatus(sess *Session, status string) int {
	n := 0
	for _, e := range sess.Endpoints {
		if e.Status == status {
			n++
		}
	}
	return n
}
