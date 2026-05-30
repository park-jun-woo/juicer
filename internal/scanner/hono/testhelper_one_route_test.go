//ff:func feature=scan type=test control=sequence topic=hono
//ff:what oneRoute 테스트 헬퍼
package hono

import "testing"

func oneRoute(t *testing.T, src string, vars map[string]bool) *routeInfo {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))
	call := findAllByType(fi.Root, "call_expression")[0]
	return extractOneRoute(call, fi.Src, vars)
}
