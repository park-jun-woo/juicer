//ff:func feature=scan type=test control=sequence topic=hono
//ff:what firstCall 테스트 헬퍼
package hono

import "testing"

func firstCall(t *testing.T, src string) *fileInfo {
	t.Helper()
	return mustParse(t, []byte(src+"\n"))
}
