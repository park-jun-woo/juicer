//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractAutoloadCall_NotRegister 테스트
package fastify

import "testing"

func TestExtractAutoloadCall_NotRegister(t *testing.T) {

	fi, calls := acFirstCall(t, "app.listen(3000);\n")
	names := map[string]bool{"autoload": true}
	for _, c := range calls {
		if _, _, ok := extractAutoloadCall(c, fi.Src, names); ok {
			t.Fatal("listen should not be autoload")
		}
	}
}
