//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 데코레이터 목록에서 첫 번째 HTTP 라우트 데코레이터를 찾는다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findRouteDecorator iterates decorators and returns the first HTTP route match.
func findRouteDecorator(decorators []*sitter.Node, src []byte) (string, string, string, int, string, string) {
	for _, dec := range decorators {
		m, p, rv, sc, rm, rc := parseRouteDecorator(dec, src)
		if m != "" {
			return m, p, rv, sc, rm, rc
		}
	}
	return "", "", "", 0, "", ""
}
