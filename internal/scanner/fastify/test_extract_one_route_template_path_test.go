//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractOneRoute_TemplatePath 테스트
package fastify

import "testing"

func TestExtractOneRoute_TemplatePath(t *testing.T) {

	fi, calls := routeCalls(t, "app.post(`/users`, h);\n")
	inst := map[string]bool{"app": true}
	var got *routeInfo
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			got = ri
		}
	}
	if got == nil || got.Method != "POST" {
		t.Fatalf("template path route = %+v", got)
	}
}
