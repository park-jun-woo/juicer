//ff:func feature=hurl type=parse control=iteration dimension=1
//ff:what 첫 번째 TODO 항목 인덱스 반환 (-1 if none)
package hurls

// firstTODO returns the index of the first TODO endpoint, or -1.
func firstTODO(sess *Session) int {
	for i, e := range sess.Endpoints {
		if e.Status == "TODO" {
			return i
		}
	}
	return -1
}
