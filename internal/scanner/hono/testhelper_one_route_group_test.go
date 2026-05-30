//ff:func feature=scan type=test control=sequence topic=hono
//ff:what oneRouteGroup 테스트 헬퍼
package hono

import "testing"

func oneRouteGroup(t *testing.T, src string, vars map[string]bool) *routeGroup {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))
	call := findAllByType(fi.Root, "call_expression")[0]
	return extractRouteGroup(call, fi.Src, vars)
}
