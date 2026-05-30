//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractOneRoute_NotInstance 테스트
package fastify

import "testing"

func TestExtractOneRoute_NotInstance(t *testing.T) {
	fi, calls := routeCalls(t, `other.get("/x", h);`+"\n")
	inst := map[string]bool{"app": true}
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			t.Fatalf("non-instance should yield nil, got %+v", ri)
		}
	}
}
