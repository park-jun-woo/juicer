//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractAutoloadCall_NotAutoloadFirstArg 테스트
package fastify

import "testing"

func TestExtractAutoloadCall_NotAutoloadFirstArg(t *testing.T) {

	fi, calls := acFirstCall(t, `app.register(somePlugin, { dir: join(d, "r") });`+"\n")
	names := map[string]bool{"autoload": true}
	for _, c := range calls {
		if _, _, ok := extractAutoloadCall(c, fi.Src, names); ok {
			t.Fatal("non-autoload plugin should not match")
		}
	}
}
