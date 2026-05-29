//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what routePrefixes: 파일 prefix 없음/있음 + wrapper 합성을 검증
package fastify

import "testing"

func TestRoutePrefixes(t *testing.T) {
	r := routeInfo{StartByte: 20}
	scopes := []wrapperScope{{Start: 0, End: 100, Prefix: "/w"}}

	if got := routePrefixes(routeInfo{StartByte: 500}, nil, nil); len(got) != 1 || got[0] != "" {
		t.Fatalf("no prefix, no wrapper: want [\"\"], got %v", got)
	}
	got := routePrefixes(r, []string{"/api"}, scopes)
	if len(got) != 1 || got[0] != "/api/w" {
		t.Fatalf("file+wrapper: want [/api/w], got %v", got)
	}
}
