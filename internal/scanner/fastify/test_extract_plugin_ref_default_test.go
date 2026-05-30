//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractPluginRef_Default 테스트
package fastify

import "testing"

func TestExtractPluginRef_Default(t *testing.T) {
	n, src := firstNodeOfType(t, `const x = "lit";`+"\n", "string")
	if got := extractPluginRef(n, src); got != "" {
		t.Fatalf("default: got %q, want empty", got)
	}
}
