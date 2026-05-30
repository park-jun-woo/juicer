//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractOneRoute_TooFewArgs 테스트
package fastify

import "testing"

func TestExtractOneRoute_TooFewArgs(t *testing.T) {
	fi, calls := routeCalls(t, `app.get("/x");`+"\n")
	inst := map[string]bool{"app": true}
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			t.Fatalf("single arg should yield nil, got %+v", ri)
		}
	}
}
