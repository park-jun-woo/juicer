//ff:func feature=scan type=test control=sequence topic=actix
//ff:what structIndexFor 테스트 헬퍼
package actix

import "testing"

func structIndexFor(t *testing.T, src string) structIndex {
	t.Helper()
	root, b := aParse(t, src)
	return buildStructIndex([]*fileInfo{{src: b, root: root}})
}
