//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractPluginRef_Arrow 테스트
package fastify

import "testing"

func TestExtractPluginRef_Arrow(t *testing.T) {
	n, src := firstNodeOfType(t, "const x = () => 1;\n", "arrow_function")
	if got := extractPluginRef(n, src); got != inlineRef {
		t.Fatalf("arrow: got %q, want %q", got, inlineRef)
	}
}
