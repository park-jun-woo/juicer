//ff:func feature=scan type=test control=sequence topic=hono
//ff:what midArgOf 테스트 헬퍼
package hono

import "testing"

func midArgOf(t *testing.T, src string) (string, string) {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)

	return extractMiddlewareName(nodes[1], fi.Src), ""
}
