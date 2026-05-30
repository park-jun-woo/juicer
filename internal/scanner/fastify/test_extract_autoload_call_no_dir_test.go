//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractAutoloadCall_NoDir 테스트
package fastify

import "testing"

func TestExtractAutoloadCall_NoDir(t *testing.T) {

	fi, calls := acFirstCall(t, `app.register(autoload, { options: { prefix: "/api" } });`+"\n")
	names := map[string]bool{"autoload": true}
	for _, c := range calls {
		if _, _, ok := extractAutoloadCall(c, fi.Src, names); ok {
			t.Fatal("missing dir should yield false")
		}
	}
}
