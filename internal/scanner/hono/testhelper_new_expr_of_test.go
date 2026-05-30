//ff:func feature=scan type=test control=sequence topic=hono
//ff:what newExprOf 테스트 헬퍼
package hono

import "testing"

func newExprOf(t *testing.T, src string) (*fileInfo, bool) {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))
	news := findAllByType(fi.Root, "new_expression")
	if len(news) == 0 {
		return fi, false
	}
	return fi, isNewHonoCall(news[0], fi.Src)
}
