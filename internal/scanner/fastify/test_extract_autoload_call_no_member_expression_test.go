//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractAutoloadCall_NoMemberExpression 테스트
package fastify

import "testing"

func TestExtractAutoloadCall_NoMemberExpression(t *testing.T) {

	fi, calls := acFirstCall(t, "foo(autoload, {});\n")
	names := map[string]bool{"autoload": true}
	for _, c := range calls {
		if _, _, ok := extractAutoloadCall(c, fi.Src, names); ok {
			t.Fatal("plain call should not be autoload")
		}
	}
}
