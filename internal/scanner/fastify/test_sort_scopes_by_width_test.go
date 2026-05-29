//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what sortScopesByWidth: 넓은 스코프가 앞으로 정렬됨을 검증
package fastify

import "testing"

func TestSortScopesByWidth(t *testing.T) {
	scopes := []wrapperScope{
		{Start: 10, End: 50, Prefix: "/narrow"},
		{Start: 0, End: 100, Prefix: "/wide"},
	}
	sortScopesByWidth(scopes)
	if scopes[0].Prefix != "/wide" || scopes[1].Prefix != "/narrow" {
		t.Fatalf("want wide first, got %v", scopes)
	}
}
